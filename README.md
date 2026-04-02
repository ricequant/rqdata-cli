# RQData CLI

> 面向量化研究与 AI Agent 的 RQData 命令行工具

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org)
[![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey.svg)]()

RQData CLI 使用 Go 实现，提供单文件可执行程序，并通过统一的 `--payload` 输入和 `ndjson/json/csv` 输出封装 RQData 常用接口。当前仓库覆盖股票、指数、基金、期货、期权、宏观和交易日历 7 大数据域，内置 66 个数据命令，适合终端查询、脚本编排和 Agent 调用。

## 文档导航

- [快速上手](docs/QUICKSTART.md)
- [命令参考](docs/rqdata_cli_commands.md)
- [Go 构建说明](BUILD_GO.md)
- [迁移说明](MIGRATION.md)
- [更新日志](CHANGELOG.md)
- [测试说明](tests/README.md)

## 特性

- 单文件二进制，支持直接分发和脚本调用
- 统一命令接口：`rqdata <group> <command> --payload '{...}'`
- 默认输出 NDJSON，适合流式处理和 AI Agent 消费
- 支持 `--format json|csv` 切换输出格式
- 支持 `--schema` 查看单命令 schema，`rqdata schema list` 查看命令清单
- 支持 `--fields` 限制输出字段
- 自动处理认证、token 缓存和 token 失效重试
- 同时提供 Go 源码构建和 npm 平台包分发

## 安装

### 通过 npm 安装

```bash
npm install -g @ricequant2026/rqdata-cli
rqdata --version
rqdata --help
```

说明：

- npm 包会通过 `bin/rqdata.js` 自动选择当前平台对应的二进制包
- 使用 npm 需要 Node.js 18+

### 从源码构建

```bash
git clone https://github.com/ricequant/rqdata-cli.git
cd rqdata-cli
VERSION=1.0.0 ./build.sh
./rqdata --help
```

`build.sh` 会在当前平台生成 `./rqdata`，并使用本地 `.cache/go-build` 作为 Go 构建缓存目录。

### 直接交叉编译

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -trimpath -ldflags="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=1.0.0" \
  -o rqdata-linux main.go
```

更多构建方式见 [BUILD_GO.md](BUILD_GO.md)。

## 认证

CLI 会按以下优先级加载凭证：

1. 环境变量: `RQDATAC_CONF` / `RQDATAC2_CONF`
2. 环境变量: `RQDATA_USERNAME` / `RQDATA_PASSWORD`
3. 系统 Keyring
4. `~/.rqdata/credentials`
5. 终端交互输入

如果您已经是 rqdata 的付费用户，您的系统环境中应该已经配置好了 `RQDATAC_CONF`, 无需更多设置。

使用交互式登录：

```bash
rqdata auth login
rqdata auth status
rqdata auth logout
```

认证成功后，token 会缓存到 `~/.rqdata/token.cache`。

## 使用方式

所有数据命令都遵循同一套调用约定：

```bash
rqdata <group> <subgroup?> <command> --payload '{...}'
```

常用参数：

- `--payload`：JSON 字符串，请求参数必填
- `--format`：输出格式，默认 `ndjson`，可选 `json`、`csv`
- `--fields`：限制返回字段
- `--schema`：打印当前命令的 schema

查看命令树：

```bash
rqdata --help
rqdata schema list
```

查看单个命令的参数定义：

```bash
rqdata stock cn price --schema
rqdata futures dominant-price --schema
```

## 命令覆盖

当前 `internal/configs/commands.json` 中定义了 66 个数据命令：

- `stock`：30 个命令，覆盖 A 股与港股基础信息、行情、财务、公告、股东、行业、北向资金等
- `index`：5 个命令，覆盖指数列表、行情、成分股与权重
- `fund`：7 个命令，覆盖基金净值、持仓、指标、经理与分红
- `futures`：5 个命令，覆盖合约列表、行情、主力合约与主力行情
- `options`：8 个命令，覆盖合约、行情、Greeks、指标与主力月份
- `macro`：8 个命令，覆盖准备金率、货币供应、CPI、PPI、GDP、PMI、利率和通用宏观查询
- `calendar`：3 个命令，覆盖交易日列表、前一交易日、后一交易日

完整参数说明见 [docs/rqdata_cli_commands.md](docs/rqdata_cli_commands.md)。

## 示例

### A 股行情

```bash
rqdata stock cn price --payload '{
  "order_book_ids": ["000001.XSHE", "600000.XSHG"],
  "start_date": "2024-01-01",
  "end_date": "2024-01-31",
  "fields": ["open", "high", "low", "close", "volume"]
}'
```

### 指数成分股

```bash
rqdata index constituents --payload '{
  "order_book_id": "000300.XSHG",
  "date": "2024-01-31"
}'
```

### 主力连续期货行情

```bash
rqdata futures dominant-price --payload '{
  "underlying_symbols": "IF",
  "start_date": "2024-01-01",
  "end_date": "2024-12-31",
  "fields": ["open", "close", "volume", "open_interest"]
}' --format json
```

### 基金指标

```bash
rqdata fund indicators --payload '{
  "order_book_ids": ["000001"],
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}' --format json
```

### 宏观利率

```bash
rqdata macro interest-rate --payload '{
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}'
```

### 交易日历

```bash
rqdata calendar trading-dates --payload '{
  "market": "cn",
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}' --format json
```

## 输出格式

### NDJSON

默认输出格式，每行一个 JSON 对象，适合管道和流式处理：

```bash
rqdata stock cn price --payload '{...}'
```

### JSON

返回结构化 JSON：

```bash
rqdata stock cn price --payload '{...}' --format json
```

### CSV

返回原始 CSV 文本：

```bash
rqdata stock cn price --payload '{...}' --format csv
```

## 仓库内附内容

- `examples/`：研究与筛选示例脚本
- `docs/`：贡献说明、测试计划、命令参考
- `scripts/generate_cli_docs.py`：根据 `commands.json` / `schema.json` 生成命令文档
- `scripts/generate_schema_json.py`：从 rqdatac 环境生成 schema

## 开发

```bash
./build.sh
./rqdata --help
python3 scripts/generate_cli_docs.py
```

集成测试相关说明见 [tests/README.md](tests/README.md)。

## License

MIT
