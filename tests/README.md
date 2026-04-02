# RQData CLI 测试框架

## 目录结构

```
tests/
├── integration/          # 集成测试
│   ├── runner.py        # 测试执行器
│   ├── cases.json       # 测试用例配置
│   └── report.py        # 报告生成（待实现）
├── utils/               # 工具模块
│   ├── cli.py          # CLI 封装
│   └── compare.py      # 结果对比
└── README.md           # 本文档
```

## 代码格式规范

**不带交易所后缀：**
- 期货：IF2406, RB2405
- 期权：10004697, 510300
- 基金：510300

**带交易所后缀：**
- A股：000001.XSHE, 600519.XSHG
- 指数：000300.XSHG, 000905.XSHG, 000852.XSHG
- 港股：00700.XHKG

## 快速开始

### 1. 设置环境变量
```bash
export RQDATA_USERNAME="your_username"
export RQDATA_PASSWORD="your_password"
```

### 2. 运行测试

```bash
# 运行所有测试
/home/lhz/.miniconda3/envs/zz1000/bin/python tests/integration/runner.py

# 运行指定套件
/home/lhz/.miniconda3/envs/zz1000/bin/python tests/integration/runner.py --suite index

# 运行 P0 优先级测试
/home/lhz/.miniconda3/envs/zz1000/bin/python tests/integration/runner.py --priority P0
```

## 测试用例配置

测试用例定义在 `tests/integration/cases.json`：

```json
{
  "suites": {
    "index": [
      {
        "id": "idx-001",
        "name": "单指数日线行情",
        "priority": "P0",
        "cmd": "index price",
        "payload": {...},
        "python": "rqdatac.get_price(...)",
        "validate": {"min_rows": 1}
      }
    ]
  }
}
```

## 添加新测试用例

### 1. 获取 API Schema

当遇到参数错误时，使用 `--schema` 查看正确的参数定义：

```bash
# 查看命令的 schema
./build/rqdata fund nav --schema

# 查看期权合约的 schema
./build/rqdata options contracts --schema
```

### 2. 编辑测试用例

根据 schema 中的参数定义，在 `tests/integration/cases.json` 中添加用例：

```json
{
  "id": "test-001",
  "name": "测试描述",
  "priority": "P0",
  "cmd": "fund nav",
  "payload": {
    "order_book_ids": ["510300"],
    "start_date": "2024-01-02"
  },
  "python": "rqdatac.fund.get_nav('510300', '2024-01-02')",
  "validate": {"min_rows": 1}
}
```

### 3. 运行测试验证

```bash
/home/lhz/.miniconda3/envs/zz1000/bin/python tests/integration/runner.py --suite fund
```

## 添加新测试用例（旧版）

## CI/CD 集成

GitHub Actions 工作流：
- 快速测试：每次 PR（P0 用例）
- 完整测试：每日定时 + Release
