#!/home/lhz/.miniconda3/envs/zz1000/bin/python
"""集成测试执行器"""
import os
import sys
import json
import subprocess
import rqdatac
import pandas as pd
from datetime import datetime
from typing import Dict, List, Any, Tuple, Optional

CLI_PATH = "./rqdata"

def run_command(cmd: str, payload: Dict, format: str = "json", timeout: int = 30) -> Optional[Dict]:
    """执行 CLI 命令并返回结果"""
    try:
        args = [CLI_PATH] + cmd.split() + ['--payload', json.dumps(payload), '--format', format]
        result = subprocess.run(args, capture_output=True, text=True, timeout=timeout)

        if result.returncode != 0:
            return {'error': result.stderr or result.stdout}

        output = result.stdout.strip()
        if not output:
            return {'data': []}

        parsed = json.loads(output)
        if isinstance(parsed, dict) and 'data' in parsed:
            return {'data': parsed['data']}
        elif isinstance(parsed, list):
            return {'data': parsed}
        else:
            return {'data': [parsed]}
    except subprocess.TimeoutExpired:
        return {'error': 'Timeout'}
    except json.JSONDecodeError as e:
        return {'error': f'JSON decode error: {e}'}
    except Exception as e:
        return {'error': str(e)}

def compare_results(cli_data: Any, python_data: Any) -> Tuple[bool, str]:
    """对比 CLI 和 Python API 结果"""
    try:
        # 转换 Python 数据
        if isinstance(python_data, pd.DataFrame):
            python_data = python_data.to_dict('records')
        elif isinstance(python_data, pd.Series):
            python_data = python_data.tolist()

        # 对比行数
        cli_rows = len(cli_data) if isinstance(cli_data, list) else 1
        py_rows = len(python_data) if isinstance(python_data, list) else 1

        if cli_rows != py_rows:
            return False, f"行数不匹配: CLI={cli_rows}, Python={py_rows}"

        return True, "通过"
    except Exception as e:
        return False, f"对比失败: {e}"

class TestRunner:
    def __init__(self, cases_file: str):
        with open(cases_file, 'r', encoding='utf-8') as f:
            self.cases = json.load(f)
        self.results = []
        self._init_rqdata()

    def _init_rqdata(self):
        username = os.getenv('RQDATA_USERNAME')
        password = os.getenv('RQDATA_PASSWORD')

        # 如果环境变量未设置，尝试从凭证文件读取
        if not username or not password:
            cred_file = os.path.expanduser('~/.rqdata/credentials')
            if os.path.exists(cred_file):
                try:
                    with open(cred_file, 'r') as f:
                        creds = json.load(f)
                        username = creds.get('username')
                        password = creds.get('password')
                except:
                    pass

        if not username or not password:
            print("❌ 错误：请设置 RQDATA_USERNAME 和 RQDATA_PASSWORD 或运行 rqdata auth login")
            sys.exit(1)

        rqdatac.init(username, password)
        print("✅ RQData 初始化成功")

    def run_suite(self, suite_name: str = None, priority: str = None):
        """运行测试套件"""
        total = passed = failed = 0

        for suite, cases in self.cases['suites'].items():
            if suite_name and suite != suite_name:
                continue

            print(f"\n{'='*60}")
            print(f"测试套件: {suite.upper()}")
            print('='*60)

            for case in cases:
                if priority and case.get('priority') != priority:
                    continue

                total += 1
                result = self._run_case(case)
                self.results.append(result)

                if result['passed']:
                    passed += 1
                    print(f"✅ {case['id']}: {case['name']}")
                else:
                    failed += 1
                    print(f"❌ {case['id']}: {case['name']}")
                    print(f"   {result['message']}")

        print(f"\n{'='*60}")
        print(f"总计: {total} | 通过: {passed} | 失败: {failed}")
        print('='*60)

        return passed, failed

    def _run_case(self, case: Dict) -> Dict:
        """运行单个测试用例"""
        try:
            # 执行 CLI
            cli_result = run_command(case['cmd'], case['payload'])
            if 'error' in cli_result:
                return {'id': case['id'], 'passed': False, 'message': f"CLI错误: {cli_result['error']}"}

            # 执行 Python API
            env = {'rqdatac': rqdatac}
            exec(f"result = {case['python']}", env)
            python_result = env['result']

            # 对比结果
            passed, message = compare_results(cli_result['data'], python_result)

            return {'id': case['id'], 'passed': passed, 'message': message}
        except Exception as e:
            return {'id': case['id'], 'passed': False, 'message': f"异常: {e}"}

if __name__ == '__main__':
    import argparse
    parser = argparse.ArgumentParser(description='RQData CLI 集成测试')
    parser.add_argument('--suite', help='指定测试套件')
    parser.add_argument('--priority', choices=['P0', 'P1', 'P2'], help='按优先级过滤')
    args = parser.parse_args()

    runner = TestRunner('tests/integration/cases.json')
    passed, failed = runner.run_suite(args.suite, args.priority)

    sys.exit(0 if failed == 0 else 1)
