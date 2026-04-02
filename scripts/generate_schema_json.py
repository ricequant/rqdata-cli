#!/usr/bin/env python3
"""
从 rqdatac 自动生成 schema.json。

用法：
    python3 scripts/generate_schema_json.py [--output schema.json]

依赖：
    pip install rqdatac
"""

import inspect
import json
import argparse
import re
import sys
import typing


# ─── 类型推断规则 ───────────────────────────────────────────────────────────────

# 按参数名精确匹配
_NAME_EXACT: typing.Dict[str, str] = {
    "market":       "string",
    "expect_df":    "boolean",
    "skip_suspended": "boolean",
    "is_total":     "boolean",
    "return_create_tm": "boolean",
    "adjusted":     "boolean",
    "frequency":    "string",
    "adjust_type":  "string",
    "statements":   "string",
    "source":       "string",
    "exchange":     "string",
    "interval":     "string",
    "option_type":  "string",
    "maturity":     "string",
    "underlying":         "string",
    "underlying_symbol":  "string",
    "underlying_symbols": "string | array<string>",
    "factor":       "string",
    "universe":     "string | array<string>",
    "margin_type":  "string",
    "trading_market": "string",
    "adjust_method":  "string",
    "price_type":     "string",
    "reserve_type":   "string",
    "factors":        "string | array<string>",
    "time_slice":     "array<string>",
    "level":        "integer",
    "rule":         "integer",
    "rank":         "integer",
    "start_rank":   "integer",
    "end_rank":     "integer",
    "n":            "integer",
    "strike":       "number",
    "type":         "string",
}

# 按参数名后缀/前缀模式匹配（顺序敏感：先具体后宽泛）
_NAME_PATTERNS: typing.List[typing.Tuple[re.Pattern, str]] = [
    # IDs
    (re.compile(r"order_book_ids$"),        "string | array<string>"),
    (re.compile(r"order_book_id$"),         "string"),
    (re.compile(r"_ids$"),                  "string | array<string>"),
    (re.compile(r"_id$"),                   "string"),
    # Dates / quarters
    (re.compile(r"(^|_)date$"),             "date"),
    (re.compile(r"(start|end)_date$"),      "date"),
    (re.compile(r"(start|end)_quarter$"),   "string"),
    (re.compile(r"quarter$"),               "string"),
    # Boolean prefixes
    (re.compile(r"^(is_|has_|skip_|return_)"), "boolean"),
    # fields
    (re.compile(r"^fields?$"),              "string | array<string>"),
    # symbols
    (re.compile(r"symbol$"),                "string"),
]


def _infer_type_from_doc(desc: str) -> typing.Optional[str]:
    """从文档描述推断类型（返回 None 表示无法推断）。"""
    desc_lower = desc.lower()
    # Boolean hints
    if re.search(r"(是否|默认为\s*(true|false)|bool)", desc_lower):
        return "boolean"
    # Integer hints (but not tuple patterns like "(int, int)")
    if re.search(r"\bint\b", desc_lower) and not re.search(r"\(int,\s*int\)", desc_lower):
        return "integer"
    # Float hints
    if re.search(r"(float|浮点)", desc_lower):
        return "number"
    # Date hints: patterns like '2015-01-07' or '20200101'
    if re.search(r"'\d{4}-\d{2}-\d{2}'", desc) or re.search(r"'\d{8}'", desc):
        return "date"
    # List/array hints
    if re.search(r"(str或list|string or list|列表|list类型|array)", desc_lower):
        return "string | array<string>"
    # String hints
    if re.search(r"(如\s*'[^']+'\s*[,，]|str类型|字符串|string类型)", desc_lower):
        return "string"
    return None


def infer_type(name: str, desc: str) -> str:  # noqa: E501
    """综合参数名和文档描述推断参数类型。"""
    # 1. 精确名称匹配
    if name in _NAME_EXACT:
        return _NAME_EXACT[name]
    # 2. 名称模式匹配
    for pattern, typ in _NAME_PATTERNS:
        if pattern.search(name):
            return typ
    # 3. 从文档推断
    inferred = _infer_type_from_doc(desc)
    if inferred:
        return inferred
    # 4. 兜底
    return "any"


# ─── 文档解析 ──────────────────────────────────────────────────────────────────

def _parse_py_type(py_type: str) -> typing.Optional[str]:
    """把 Python 类型标注字符串转换为 schema 类型。
    例如：
      'str | list[str]'  -> 'string | array<string>'
      'list[str]'        -> 'array<string>'
      'str'              -> 'string'
      'int'              -> 'integer'
      'bool'             -> 'boolean'
      'float'            -> 'number'
    """
    t = py_type.strip().lower()
    # 去掉 optional 包装
    t = re.sub(r'optional\[(.+)\]', r'\1', t).strip()

    has_list = bool(re.search(r'list(\[.*?\])?', t))
    has_str  = bool(re.search(r'\bstr\b', t))
    has_int  = bool(re.search(r'\bint\b', t))
    has_bool = bool(re.search(r'\bbool\b', t))
    has_float = bool(re.search(r'\bfloat\b', t))

    if has_bool:
        return "boolean"
    if has_int and not has_str and not has_list:
        return "integer"
    if has_float and not has_str and not has_list:
        return "number"
    if has_list and has_str:
        return "string | array<string>"
    if has_list:
        return "array<string>"
    if has_str:
        return "string"
    # date 类型
    if re.search(r'date', t):
        return "date"
    return None


def parse_docstring_params(doc: str) -> typing.Dict[str, typing.Dict[str, str]]:
    """从 NumPy 或 Sphinx 风格的 docstring 中提取 {param_name: {desc, py_type}}。"""
    if not doc:
        return {}
    params: dict = {}

    # 尝试 Sphinx 风格：:param name: description
    sphinx_pattern = re.compile(r':param\s+(\w+):\s*([^\n:]+(?:\n(?!\s*:)[^\n]+)*)', re.MULTILINE)
    for m in sphinx_pattern.finditer(doc):
        pname = m.group(1)
        desc = re.sub(r'\s+', ' ', m.group(2).strip())
        params[pname] = {"desc": desc, "py_type": None}

    if params:
        return params

    # 尝试 NumPy 风格：Parameters 部分
    params_match = re.search(
        r'Parameters\s*\n\s*-+\s*\n(.*?)(?=\n\s*(?:Returns|Raises|Notes|Examples|See Also|\Z))',
        doc, re.DOTALL | re.IGNORECASE
    )
    if not params_match:
        return {}

    params_section = params_match.group(1)

    # 匹配 "name : type\n    description" 格式，捕获 type 部分
    pattern = re.compile(r'^(\w+)\s*:\s*([^\n]*)\n((?:\s{4,}.*\n?)*)', re.MULTILINE)
    for m in pattern.finditer(params_section):
        pname   = m.group(1)
        py_type = m.group(2).strip()
        desc    = re.sub(r'\s+', ' ', m.group(3).strip())
        params[pname] = {"desc": desc, "py_type": py_type or None}

    return params


# ─── 主逻辑 ───────────────────────────────────────────────────────────────────

def extract_param_schema(func) -> dict:
    """从函数签名 + 文档推断参数 schema。"""
    try:
        sig = inspect.signature(func)
    except (ValueError, TypeError):
        return {}

    doc_params = parse_docstring_params(inspect.getdoc(func) or "")

    params = {}
    for name, param in sig.parameters.items():
        if name in ("self", "cls", "kwargs", "args"):
            continue
        required = param.default is inspect.Parameter.empty
        info     = doc_params.get(name, {})
        desc     = info.get("desc", "")
        py_type  = info.get("py_type")

        # 优先用 docstring 里的 Python 类型标注，再回落到 infer_type
        type_str = (_parse_py_type(py_type) if py_type else None) or infer_type(name, desc)

        entry: dict = {
            "type": type_str,
            "required": required,
            "description": desc,
        }
        if not required and param.default is not None and param.default is not inspect.Parameter.empty:
            entry["default"] = str(param.default)
        params[name] = entry
    return params


def get_method_func(rqdatac, method_name: str):
    """按 method_name 从 rqdatac 中找到对应函数（支持 a.b 格式）。"""
    parts = method_name.split(".")
    obj = rqdatac
    for part in parts:
        obj = getattr(obj, part, None)
        if obj is None:
            return None
    return obj if callable(obj) else None


def main():
    parser = argparse.ArgumentParser(description="Generate schema.json from rqdatac")
    parser.add_argument("--output", default="schema.json", help="Output file path")
    parser.add_argument("--commands", default="commands.json", help="commands.json path")
    args = parser.parse_args()

    try:
        import rqdatac
    except ImportError:
        print("Error: rqdatac is not installed. Run: pip install rqdatac", file=sys.stderr)
        sys.exit(1)

    rqdatac.init()

    with open(args.commands) as f:
        commands_data = json.load(f)

    schema = {}
    seen_methods: typing.Set[str] = set()

    def process_commands(commands):
        for cmd in commands:
            api_method = cmd["api_method"]
            if api_method in seen_methods:
                continue
            seen_methods.add(api_method)

            func = get_method_func(rqdatac, api_method)
            if func is None:
                print(f"  [warn] method not found in rqdatac: {api_method}", file=sys.stderr)
                schema[api_method] = {"parameters": {}, "returns": {}}
                continue

            params = extract_param_schema(func)
            schema[api_method] = {
                "parameters": params,
                "returns": {},
            }
            any_count = sum(1 for p in params.values() if p["type"] == "any")
            print(f"  [ok] {api_method} ({len(params)} params, {any_count} unresolved)")

    def process_group(group):
        if "commands" in group:
            process_commands(group["commands"])
        if "groups" in group:
            for subgroup in group["groups"]:
                process_group(subgroup)

    for group in commands_data["groups"]:
        process_group(group)

    with open(args.output, "w", encoding="utf-8") as f:
        json.dump(schema, f, ensure_ascii=False, indent=2)

    print(f"\nGenerated {args.output} with {len(schema)} API methods.")


if __name__ == "__main__":
    main()
