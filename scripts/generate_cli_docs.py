#!/usr/bin/env python3
"""
自动生成 CLI 命令参考文档
从 commands.json 和 schema.json 生成 rqdata_cli_commands.md
"""

import json
import sys
from pathlib import Path


def flatten_commands(groups, path_prefix=""):
    """从树形结构中提取所有命令，附带完整路径"""
    result = []
    for group in groups:
        group_path = (path_prefix + " " + group["name"]).strip()
        if "groups" in group:
            result.extend(flatten_commands(group["groups"], group_path))
        if "commands" in group:
            for cmd in group["commands"]:
                result.append({
                    "path": group_path.split() + [cmd["name"]],
                    "group_name": group["name"],
                    "group_desc": group.get("description", ""),
                    **cmd
                })
    return result


def load_data(project_root: Path):
    commands_file = project_root / 'internal' / 'configs' / 'commands.json'
    schema_file   = project_root / 'internal' / 'configs' / 'schema.json'

    if not commands_file.exists():
        print(f"Error: {commands_file} not found", file=sys.stderr)
        sys.exit(1)

    with open(commands_file, encoding='utf-8') as f:
        commands_data = json.load(f)

    schema_data = {}
    if schema_file.exists():
        with open(schema_file, encoding='utf-8') as f:
            schema_data = json.load(f)

    commands = flatten_commands(commands_data['groups'])
    return commands, schema_data, commands_data['groups']


def generate_markdown(commands, schema_data, groups, output_path: Path):
    # 按顶层分组组织
    categories = {}
    for cmd in commands:
        cat = cmd['path'][0]
        categories.setdefault(cat, {"desc": "", "commands": []})
        categories[cat]["commands"].append(cmd)

    # 从 groups 获取顶层分组描述
    for group in groups:
        if group["name"] in categories:
            categories[group["name"]]["desc"] = group.get("description", "")

    total = len(commands)
    md = f"""# RQData CLI 命令参考手册

> 自动生成于 commands.json

## 命令统计

- **总命令数**: {total} 个
- **命令类别**: {len(categories)} 个

## 目录

"""
    order = ['index', 'stock', 'calendar', 'fund', 'futures', 'options', 'macro']
    for cat in order:
        if cat in categories:
            desc = categories[cat]["desc"]
            md += f"- [{cat.upper()} - {desc}](#{cat})\n"

    md += "\n---\n\n"

    for cat in order:
        if cat not in categories:
            continue
        desc = categories[cat]["desc"]
        md += f"\n## {cat.upper()} - {desc}\n\n"
        for cmd in categories[cat]["commands"]:
            path_str = " ".join(cmd['path'])
            api_method = cmd['api_method']
            md += f"### `{path_str}`\n\n"
            md += f"**描述**: {cmd.get('description', '')}\n\n"
            md += f"**API Method**: `{api_method}`\n\n"

            schema = schema_data.get(api_method, {})
            params = schema.get('parameters', {})
            if params:
                md += "**Payload 参数**:\n\n"
                md += "| 参数名 | 类型 | 必填 | 说明 |\n"
                md += "|--------|------|------|------|\n"
                for pname, pinfo in params.items():
                    req = "✅" if pinfo.get('required') else "❌"
                    md += f"| `{pname}` | {pinfo.get('type','')} | {req} | {pinfo.get('description','')} |\n"
                md += "\n"

            returns = schema.get('returns', {})
            if returns:
                md += "**返回字段**:\n\n"
                md += "| 字段名 | 说明 |\n"
                md += "|--------|------|\n"
                for fname, fdesc in returns.items():
                    md += f"| `{fname}` | {fdesc} |\n"
                md += "\n"

            md += "---\n\n"

    with open(output_path, 'w', encoding='utf-8') as f:
        f.write(md)
    return len(md.splitlines())


def main():
    project_root = Path(__file__).parent.parent
    output_md = project_root / 'docs' / 'rqdata_cli_commands.md'

    print("Generating CLI commands reference documentation...")
    commands, schema_data, groups = load_data(project_root)
    print(f"  Found {len(commands)} commands, {len(schema_data)} schema entries")

    lines = generate_markdown(commands, schema_data, groups, output_md)
    print(f"  Generated {lines} lines -> {output_md}")
    print("Done.")


if __name__ == '__main__':
    main()
