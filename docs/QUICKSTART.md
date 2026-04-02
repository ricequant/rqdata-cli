# RQData CLI 快速上手

本指南基于当前仓库中的 CLI 实现编写，目标是在几分钟内完成安装、认证并发起第一批查询。

## 1. 安装

### npm 安装

```bash
npm install -g @ricequant2026/rqdata-cli
rqdata --version
rqdata --help
```

说明：

- npm 包会自动解析当前平台并调用对应的二进制包
- 使用 npm 需要 Node.js 18+

### 从源码构建

前置要求：Go 1.21+

```bash
git clone https://github.com/ricequant/rqdata-cli.git
cd rqdata-cli
VERSION=1.0.0 ./build.sh
./rqdata --help
```

如果需要交叉编译，请参考 [BUILD_GO.md](../BUILD_GO.md)。

## 2. 配置认证

CLI 会按以下优先级查找凭证：

1. `RQDATAC_CONF` / `RQDATAC2_CONF`
2. `RQDATA_USERNAME` / `RQDATA_PASSWORD`
3. 系统 Keyring
4. `~/.rqdata/credentials`
5. 终端交互输入

推荐方式是直接配置环境变量：

```bash
export RQDATA_USERNAME="your_phone_or_email"
export RQDATA_PASSWORD="your_password"
```

也可以使用交互式登录：

```bash
rqdata auth login
```

验证当前状态：

```bash
rqdata auth status
```

输出为 JSON，示例：

```json
{
  "credentials": "RQDATA_USERNAME",
  "username": "your_phone_or_email",
  "token_cached": true,
  "token_valid": true,
  "expires_at": "2026-04-02T12:00:00+08:00"
}
```

清除已保存凭证：

```bash
rqdata auth logout
```

## 3. 理解命令结构

数据命令统一采用如下形式：

```bash
rqdata <group> <subgroup?> <command> --payload '{...}'
```

例如：

- `rqdata stock cn price`
- `rqdata index constituents`
- `rqdata fund indicators`
- `rqdata futures dominant-price`
- `rqdata macro interest-rate`

所有数据命令都支持以下通用参数：

- `--payload`：请求参数，必填
- `--format`：输出格式，默认 `ndjson`，可选 `json`、`csv`
- `--fields`：限制返回字段
- `--schema`：查看当前命令的 schema

查看命令树：

```bash
rqdata --help
rqdata schema list
```

查看单个命令的参数说明：

```bash
rqdata stock cn price --schema
rqdata stock cn financial --schema
rqdata calendar trading-dates --schema
```

## 4. 第一批查询

### A 股行情

```bash
rqdata stock cn price --payload '{
  "order_book_ids": ["000001.XSHE", "600000.XSHG"],
  "start_date": "2024-01-01",
  "end_date": "2024-01-31",
  "fields": ["open", "high", "low", "close", "volume"]
}'
```

默认输出为 NDJSON，每行一个 JSON 对象。

### 指数成分股

```bash
rqdata index constituents --payload '{
  "order_book_id": "000300.XSHG",
  "date": "2024-01-31"
}' --format json
```

### A 股财务数据

```bash
rqdata stock cn financial --payload '{
  "order_book_ids": ["000001.XSHE"],
  "fields": ["revenue", "net_profit", "total_assets", "total_liabilities"],
  "start_quarter": "2023Q1",
  "end_quarter": "2024Q3"
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

### 主力连续期货行情

```bash
rqdata futures dominant-price --payload '{
  "underlying_symbols": "IF",
  "start_date": "2024-01-01",
  "end_date": "2024-12-31",
  "fields": ["open", "close", "volume", "open_interest"]
}' --format json
```

### 宏观利率

```bash
rqdata macro interest-rate --payload '{
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}'
```

### 交易日列表

```bash
rqdata calendar trading-dates --payload '{
  "market": "cn",
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}' --format json
```

## 5. 输出格式

### NDJSON

适合管道、流式处理和 Agent：

```bash
rqdata stock cn price --payload '{...}'
```

### JSON

适合程序一次性读取完整结果：

```bash
rqdata stock cn price --payload '{...}' --format json
```

### CSV

适合导出或接 Excel：

```bash
rqdata stock cn price --payload '{...}' --format csv
```

## 6. 常用技巧

### 用 `--schema` 确认参数

在不确定 payload 字段名时，不要猜，直接看 schema：

```bash
rqdata options greeks --schema
rqdata stock hk financial --schema
```

### 用 `--fields` 缩小返回结果

```bash
rqdata stock cn shares --payload '{
  "order_book_ids": ["000001.XSHE"],
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}' --fields total_a,total_circulation_a --format json
```

### 配合 `jq` 做简单处理

```bash
rqdata stock cn price --payload '{...}' | jq '.close'
rqdata stock cn price --payload '{...}' | jq -s 'map(.close | tonumber) | add / length'
```

### 输出到文件

```bash
rqdata stock cn price --payload '{...}' > data.ndjson
rqdata stock cn price --payload '{...}' --format json > data.json
rqdata stock cn price --payload '{...}' --format csv > data.csv
```

## 7. 常见格式

### 代码格式

- A 股上交所：`600000.XSHG`
- A 股深交所：`000001.XSHE`
- 港股：`00700.XHKG`
- 指数：`000300.XSHG`

### 日期格式

- 日级日期：`YYYY-MM-DD`
- 季度：`YYYYQ1`、`YYYYQ2`、`YYYYQ3`、`YYYYQ4`

## 8. 下一步

- 全量命令与参数说明见 [rqdata_cli_commands.md](./rqdata_cli_commands.md)
- 构建与发布见 [BUILD_GO.md](../BUILD_GO.md)
- 测试说明见 [tests/README.md](../tests/README.md)
- 示例脚本见 `examples/`
