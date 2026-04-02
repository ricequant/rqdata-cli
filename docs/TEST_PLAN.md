# RQData CLI 全面测试方案

> 更新时间: 2026-03-12 17:56
> 测试目标：验证所有 CLI 命令与 Python API 的一致性
> 最新测试结果: 74/74 通过 (100%)

## 📋 测试概览

### 测试范围

| 类别 | 测试数 | 通过 | 通过率 |
|------|--------|------|--------|
| INDEX (指数) | 8 | 8 | 100% |
| STOCK CN (A股) | 23 | 23 | 100% |
| STOCK HK (港股) | 8 | 8 | 100% |
| CALENDAR (日历) | 5 | 5 | 100% |
| FUND (基金) | 9 | 9 | 100% |
| FUTURES (期货) | 8 | 8 | 100% |
| OPTIONS (期权) | 7 | 7 | 100% |
| MACRO (宏观) | 6 | 6 | 100% |
| **总计** | **74** | **74** | **100%** |

### 测试优先级统计

| 优先级 | 总数 | 通过 | 通过率 |
|--------|------|------|--------|
| P0 | 15 | 15 | 100% |
| P1 | 26 | 26 | 100% |
| P2 | 33 | 33 | 100% |

---

## 📝 完整测试用例列表

| ID | 优先级 | 命令 | 场景 | 状态 |
|----|--------|------|------|------|
| 1 | P0 | `index price` | 单指数日线行情 | ✅ |
| 2 | P0 | `index price` | 多指数日线行情 | ✅ |
| 3 | P1 | `index price` | 指定返回字段 | ✅ |
| 4 | P0 | `index instruments` | 查询指定指数信息 | ✅ |
| 5 | P1 | `index instruments` | 查询所有指数 | ✅ |
| 6 | P0 | `index constituents` | 查询沪深300成分股 | ✅ |
| 7 | P1 | `index constituents` | 查询中证500成分股 | ✅ |
| 8 | P0 | `index weights` | 查询沪深300成分股权重 | ✅ |
| 9 | P0 | `stock cn price` | 单股票前复权日线 | ✅ |
| 10 | P0 | `stock cn price` | 多股票不复权 | ✅ |
| 11 | P1 | `stock cn price` | 分钟线数据 | ✅ |
| 12 | P0 | `stock cn instruments` | 查询指定股票信息 | ✅ |
| 13 | P1 | `stock cn instruments` | 查询所有A股 | ✅ |
| 14 | P0 | `stock cn dividend` | 查询分红数据 | ✅ |
| 15 | P1 | `stock cn dividend` | 多股票分红查询 | ✅ |
| 16 | P1 | `stock cn split` | 查询拆分数据 | ✅ |
| 17 | P1 | `stock cn shares` | 查询股本数据 | ✅ |
| 18 | P1 | `stock cn ex-factor` | 查询复权因子 | ✅ |
| 19 | P1 | `stock cn industry` | 查询申万行业分类 | ✅ |
| 20 | P1 | `stock cn turnover-rate` | 查询换手率 | ✅ |
| 21 | P2 | `stock cn margin` | 查询融资融券标的 | ✅ |
| 22 | P0 | `stock cn financial` | 查询财务数据 | ✅ |
| 23 | P1 | `stock cn financial` | 查询多季度财务数据 | ✅ |
| 24 | P2 | `stock cn suspended` | 查询停牌状态 | ✅ |
| 25 | P2 | `stock cn st-status` | 查询ST状态 | ✅ |
| 26 | P2 | `stock cn shares` | 查询股本数据(日期范围) | ✅ |
| 27 | P1 | `stock hk price` | 港股日线行情 | ✅ |
| 28 | P1 | `stock hk instruments` | 查询港股信息 | ✅ |
| 29 | P2 | `stock hk dividend` | 查询港股分红 | ✅ |
| 30 | P1 | `stock hk price` | 多只港股查询 | ✅ |
| 31 | P2 | `stock hk instruments` | 查询所有港股 | ✅ |
| 32 | P0 | `calendar trading-dates` | 查询交易日列表 | ✅ |
| 33 | P0 | `calendar prev` | 查询前一交易日 | ✅ |
| 34 | P0 | `calendar next` | 查询后一交易日 | ✅ |
| 35 | P1 | `calendar trading-dates` | 查询全年交易日 | ✅ |
| 36 | P1 | `calendar prev` | 节假日前一交易日 | ✅ |
| 37 | P0 | `fund instruments` | 查询基金信息 | ✅ |
| 38 | P0 | `fund nav` | 查询基金净值 | ✅ |
| 39 | P1 | `fund nav` | 多基金净值查询 | ✅ |
| 40 | P1 | `fund holdings` | 查询基金持仓 | ✅ |
| 41 | P2 | `fund instruments` | 查询所有基金 | ✅ |
| 42 | P1 | `fund price` | ETF行情查询 | ✅ |
| 43 | P2 | `fund indicators` | 查询基金指标 | ✅ |
| 44 | P2 | `fund dividend` | 查询基金分红 | ✅ |
| 45 | P1 | `futures instruments` | 查询期货合约信息 | ✅ |
| 46 | P2 | `futures instruments` | 查询所有期货合约 | ✅ |
| 47 | P1 | `futures price` | 期货日线行情 | ✅ |
| 48 | P1 | `futures price` | 多合约期货行情 | ✅ |
| 49 | P1 | `futures dominant` | 查询主力合约 | ✅ |
| 50 | P1 | `futures dominant-price` | 查询主力连续行情 | ✅ |
| 51 | P1 | `futures price` | 期货分钟线 | ✅ |
| 52 | P2 | `futures dominant-price` | 多品种主力连续行情 | ✅ |
| 53 | P2 | `options contracts` | 查询期权合约信息 | ✅ |
| 54 | P2 | `options price` | 期权日线行情 | ✅ |
| 55 | P2 | `options greeks` | 查询期权希腊字母 | ✅ |
| 56 | P2 | `options contracts` | 查询期权合约链 | ✅ |
| 57 | P2 | `options price` | 多期权合约行情 | ✅ |
| 58 | P2 | `options contract-property` | 查询期权合约属性 | ✅ |
| 59 | P2 | `options contracts` | 按行权价筛选 | ✅ |
| 60 | P2 | `macro factors` | 查询宏观因子数据 | ✅ |
| 61 | P2 | `macro reserve-ratio` | 查询存款准备金率 | ✅ |
| 62 | P2 | `macro money-supply` | 查询货币供应量 | ✅ |
| 63 | P2 | `stock cn northbound` | 查询北向资金 | ✅ |
| 64 | P2 | `stock cn financial-indicator` | 查询财务衍生指标 | ✅ |
| 65 | P2 | `stock cn shareholder-top10` | 查询十大股东 | ✅ |
| 66 | P2 | `stock cn financial-express` | 查询业绩快报 | ✅ |
| 67 | P2 | `stock cn forecast` | 查询业绩预告 | ✅ |
| 68 | P2 | `fund manager` | 查询基金经理 | ✅ |
| 69 | P2 | `stock hk turnover-rate` | 查询港股换手率 | ✅ |
| 70 | P2 | `stock hk shares` | 查询港股股本 | ✅ |
| 71 | P2 | `stock hk financial` | 查询港股财务数据 | ✅ |
| 72 | P2 | `macro gdp` | 查询GDP数据 | ✅ |
| 73 | P2 | `macro price-ppi` | 查询PPI数据 | ✅ |
| 74 | P2 | `macro price-cpi` | 查询CPI数据 | ✅ |

---

## INDEX (指数) 测试用例

### #1 单指数日线行情 ✅

- **优先级**: P0
- **命令**: `index price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (4 行, 7 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000300.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price("000300.XSHG", "2024-01-02", "2024-01-05", "1d")
```

---

### #2 多指数日线行情 ✅

- **优先级**: P0
- **命令**: `index price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (8 行, 7 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000300.XSHG",
    "000905.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price(["000300.XSHG", "000905.XSHG"], "2024-01-02", "2024-01-05", "1d")
```

---

### #3 指定返回字段 ✅

- **优先级**: P1
- **命令**: `index price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (4 行, 2 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000300.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05",
  "fields": [
    "close",
    "volume"
  ]
}
```

**Python 对照:**
```python
rqdatac.get_price("000300.XSHG", "2024-01-02", "2024-01-05", "1d", fields=["close", "volume"])
```

---

### #4 查询指定指数信息 ✅

- **优先级**: P0
- **命令**: `index instruments`
- **API Method**: `all_instruments`
- **结果**: ✅ 通过 (数据匹配 (1 行))

**Payload:**
```json
{
  "order_book_ids": [
    "000300.XSHG"
  ]
}
```

**Python 对照:**
```python
rqdatac.instruments("000300.XSHG")
```

---

### #5 查询所有指数 ✅

- **优先级**: P1
- **命令**: `index instruments`
- **API Method**: `all_instruments`
- **结果**: ✅ 通过 (数据匹配 (7749 行, 15 个共同字段))

**Payload:**
```json
{}
```

**Python 对照:**
```python
rqdatac.all_instruments(type="Index")
```

---

### #6 查询沪深300成分股 ✅

- **优先级**: P0
- **命令**: `index constituents`
- **API Method**: `index.get_constituents`
- **结果**: ✅ 通过 (数据匹配 (300 行))

**Payload:**
```json
{
  "order_book_id": "000300.XSHG",
  "date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.index_components("000300.XSHG", "2024-01-31")
```

---

### #7 查询中证500成分股 ✅

- **优先级**: P1
- **命令**: `index constituents`
- **API Method**: `index.get_constituents`
- **结果**: ✅ 通过 (数据匹配 (500 行))

**Payload:**
```json
{
  "order_book_id": "000905.XSHG",
  "date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.index_components("000905.XSHG", "2024-01-31")
```

---

### #8 查询沪深300成分股权重 ✅

- **优先级**: P0
- **命令**: `index weights`
- **API Method**: `index.get_weights`
- **结果**: ✅ 通过 (数据匹配 (300 行))

**Payload:**
```json
{
  "order_book_id": "000300.XSHG",
  "date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.index_weights("000300.XSHG", "2024-01-31")
```

---

## STOCK CN (A股) 测试用例

### #9 单股票前复权日线 ✅

- **优先级**: P0
- **命令**: `stock cn price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (4 行, 10 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05",
  "adjust_type": "pre"
}
```

**Python 对照:**
```python
rqdatac.get_price("600000.XSHG", "2024-01-02", "2024-01-05", "1d", adjust_type="pre")
```

---

### #10 多股票不复权 ✅

- **优先级**: P0
- **命令**: `stock cn price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (8 行, 10 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG",
    "000001.XSHE"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05",
  "adjust_type": "none"
}
```

**Python 对照:**
```python
rqdatac.get_price(["600000.XSHG", "000001.XSHE"], "2024-01-02", "2024-01-05", "1d", adjust_type="none")
```

---

### #11 分钟线数据 ✅

- **优先级**: P1
- **命令**: `stock cn price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (240 行, 7 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-02 09:30:00",
  "end_date": "2024-01-02 10:00:00",
  "frequency": "1m"
}
```

**Python 对照:**
```python
rqdatac.get_price("600000.XSHG", "2024-01-02 09:30:00", "2024-01-02 10:00:00", "1m")
```

---

### #12 查询指定股票信息 ✅

- **优先级**: P0
- **命令**: `stock cn instruments`
- **API Method**: `all_instruments`
- **结果**: ✅ 通过 (数据匹配 (1 行))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ]
}
```

**Python 对照:**
```python
rqdatac.instruments("600000.XSHG")
```

---

### #13 查询所有A股 ✅

- **优先级**: P1
- **命令**: `stock cn instruments`
- **API Method**: `all_instruments`
- **结果**: ✅ 通过 (数据匹配 (5531 行, 22 个共同字段))

**Payload:**
```json
{}
```

**Python 对照:**
```python
rqdatac.all_instruments(type="CS")
```

---

### #14 查询分红数据 ✅

- **优先级**: P0
- **命令**: `stock cn dividend`
- **API Method**: `get_dividend`
- **结果**: ✅ 通过 (数据匹配 (3 行, 7 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2023-01-01"
}
```

**Python 对照:**
```python
rqdatac.get_dividend("600000.XSHG", start_date="2023-01-01")
```

---

### #15 多股票分红查询 ✅

- **优先级**: P1
- **命令**: `stock cn dividend`
- **API Method**: `get_dividend`
- **结果**: ✅ 通过 (数据匹配 (8 行, 7 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG",
    "000001.XSHE"
  ],
  "start_date": "2023-01-01"
}
```

**Python 对照:**
```python
rqdatac.get_dividend(["600000.XSHG", "000001.XSHE"], start_date="2023-01-01")
```

---

### #16 查询拆分数据 ✅

- **优先级**: P1
- **命令**: `stock cn split`
- **API Method**: `get_split`
- **结果**: ✅ 通过 (数据为空（匹配）)

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2020-01-01"
}
```

**Python 对照:**
```python
rqdatac.get_split("600000.XSHG", start_date="2020-01-01")
```

---

### #17 查询股本数据 ✅

- **优先级**: P1
- **命令**: `stock cn shares`
- **API Method**: `get_shares`
- **结果**: ✅ 通过 (数据匹配 (22 行, 6 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.get_shares("600000.XSHG", start_date="2024-01-01", end_date="2024-01-31")
```

---

### #18 查询复权因子 ✅

- **优先级**: P1
- **命令**: `stock cn ex-factor`
- **API Method**: `get_ex_factor`
- **结果**: ✅ 通过 (数据为空（匹配）)

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.get_ex_factor("600000.XSHG", start_date="2024-01-01", end_date="2024-01-31")
```

---

### #19 查询申万行业分类 ✅

- **优先级**: P1
- **命令**: `stock cn industry`
- **API Method**: `shenwan_instrument_industry`
- **结果**: ✅ 通过 (数据匹配 (1 行, 2 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.get_instrument_industry("600000.XSHG", date="2024-01-31")
```

---

### #20 查询换手率 ✅

- **优先级**: P1
- **命令**: `stock cn turnover-rate`
- **API Method**: `get_turnover_rate`
- **结果**: ✅ 通过 (数据匹配 (4 行, 5 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_turnover_rate("600000.XSHG", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #21 查询融资融券标的 ✅

- **优先级**: P2
- **命令**: `stock cn margin`
- **API Method**: `get_margin_stocks`
- **结果**: ✅ 通过 (数据匹配 (3631 行))

**Payload:**
```json
{
  "date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.get_margin_stocks(date="2024-01-31")
```

---

### #22 查询财务数据 ✅

- **优先级**: P0
- **命令**: `stock cn financial`
- **API Method**: `get_pit_financials_ex`
- **结果**: ✅ 通过 (数据匹配 (1 行, 5 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "fields": [
    "revenue",
    "net_profit"
  ],
  "quarter": "2023Q4"
}
```

**Python 对照:**
```python
rqdatac.get_pit_financials_ex("600000.XSHG", fields=["revenue", "net_profit"], start_quarter="2023Q4", end_quarter="2023Q4")
```

---

### #23 查询多季度财务数据 ✅

- **优先级**: P1
- **命令**: `stock cn financial`
- **API Method**: `get_pit_financials_ex`
- **结果**: ✅ 通过 (数据匹配 (4 行, 5 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "fields": [
    "revenue",
    "net_profit"
  ],
  "start_quarter": "2023Q1",
  "end_quarter": "2023Q4"
}
```

**Python 对照:**
```python
rqdatac.get_pit_financials_ex("600000.XSHG", fields=["revenue", "net_profit"], start_quarter="2023Q1", end_quarter="2023Q4")
```

---

### #24 查询停牌状态 ✅

- **优先级**: P2
- **命令**: `stock cn suspended`
- **API Method**: `is_suspended`
- **结果**: ✅ 通过 (数据匹配 (4 行, 1 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.is_suspended("600000.XSHG", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #25 查询ST状态 ✅

- **优先级**: P2
- **命令**: `stock cn st-status`
- **API Method**: `is_st_stock`
- **结果**: ✅ 通过 (数据匹配 (4 行, 1 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.is_st_stock("600000.XSHG", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #26 查询股本数据(日期范围) ✅

- **优先级**: P2
- **命令**: `stock cn shares`
- **API Method**: `get_shares`
- **结果**: ✅ 通过 (数据匹配 (4 行, 6 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_shares("600000.XSHG", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #63 查询北向资金 ✅

- **优先级**: P2
- **命令**: `stock cn northbound`
- **API Method**: `get_stock_connect`
- **结果**: ✅ 通过 (数据匹配 (4 行, 3 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_stock_connect("600000.XSHG", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #64 查询财务衍生指标 ✅

- **优先级**: P2
- **命令**: `stock cn financial-indicator`
- **API Method**: `get_factor`
- **结果**: ✅ 通过 (数据匹配 (4 行, 2 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "factor": [
    "pe_ratio",
    "pb_ratio"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_factor("600000.XSHG", ["pe_ratio", "pb_ratio"], start_date="2024-01-02", end_date="2024-01-05")
```

---

### #65 查询十大股东 ✅

- **优先级**: P2
- **命令**: `stock cn shareholder-top10`
- **API Method**: `get_main_shareholder`
- **结果**: ✅ 通过 (数据匹配 (40 行, 9 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ],
  "start_date": "2023-01-01",
  "end_date": "2023-12-31"
}
```

**Python 对照:**
```python
rqdatac.get_main_shareholder("600000.XSHG", start_date="2023-01-01", end_date="2023-12-31")
```

---

### #66 查询业绩快报 ✅

- **优先级**: P2
- **命令**: `stock cn financial-express`
- **API Method**: `current_performance`
- **结果**: ✅ 通过 (数据匹配 (1 行, 40 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ]
}
```

**Python 对照:**
```python
rqdatac.current_performance("600000.XSHG")
```

---

### #67 查询业绩预告 ✅

- **优先级**: P2
- **命令**: `stock cn forecast`
- **API Method**: `performance_forecast`
- **结果**: ✅ 通过 (数据匹配 (2 行, 16 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "600000.XSHG"
  ]
}
```

**Python 对照:**
```python
rqdatac.performance_forecast("600000.XSHG")
```

---

## STOCK HK (港股) 测试用例

### #27 港股日线行情 ✅

- **优先级**: P1
- **命令**: `stock hk price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (4 行, 10 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "00700.XHKG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price("00700.XHKG", "2024-01-02", "2024-01-05", "1d", market="hk")
```

---

### #28 查询港股信息 ✅

- **优先级**: P1
- **命令**: `stock hk instruments`
- **API Method**: `all_instruments`
- **结果**: ✅ 通过 (数据匹配 (1 行))

**Payload:**
```json
{
  "order_book_ids": [
    "00700.XHKG"
  ]
}
```

**Python 对照:**
```python
rqdatac.instruments("00700.XHKG", market="hk")
```

---

### #29 查询港股分红 ✅

- **优先级**: P2
- **命令**: `stock hk dividend`
- **API Method**: `get_dividend`
- **结果**: ✅ 通过 (数据匹配 (3 行, 5 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "00700.XHKG"
  ],
  "start_date": "2023-01-01"
}
```

**Python 对照:**
```python
rqdatac.get_dividend("00700.XHKG", start_date="2023-01-01", market="hk")
```

---

### #30 多只港股查询 ✅

- **优先级**: P1
- **命令**: `stock hk price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (8 行, 10 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "00700.XHKG",
    "09988.XHKG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price(["00700.XHKG", "09988.XHKG"], "2024-01-02", "2024-01-05", "1d", market="hk")
```

---

### #31 查询所有港股 ✅

- **优先级**: P2
- **命令**: `stock hk instruments`
- **API Method**: `all_instruments`
- **结果**: ✅ 通过 (数据匹配 (3453 行, 15 个共同字段))

**Payload:**
```json
{}
```

**Python 对照:**
```python
rqdatac.all_instruments(type="CS", market="hk")
```

---

### #69 查询港股换手率 ✅

- **优先级**: P2
- **命令**: `stock hk turnover-rate`
- **API Method**: `get_turnover_rate`
- **结果**: ✅ 通过 (数据匹配 (4 行, 5 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "00700.XHKG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_turnover_rate("00700.XHKG", start_date="2024-01-02", end_date="2024-01-05", market="hk")
```

---

### #70 查询港股股本 ✅

- **优先级**: P2
- **命令**: `stock hk shares`
- **API Method**: `get_shares`
- **结果**: ✅ 通过 (数据匹配 (4 行, 7 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "00700.XHKG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_shares("00700.XHKG", start_date="2024-01-02", end_date="2024-01-05", market="hk")
```

---

### #71 查询港股财务数据 ✅

- **优先级**: P2
- **命令**: `stock hk financial`
- **API Method**: `get_pit_financials_ex`
- **结果**: ✅ 通过 (数据匹配 (1 行, 5 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "00700.XHKG"
  ],
  "fields": [
    "revenue"
  ],
  "start_quarter": "2023Q4",
  "end_quarter": "2023Q4"
}
```

**Python 对照:**
```python
rqdatac.get_pit_financials_ex(["00700.XHKG"], fields=["revenue"], start_quarter="2023Q4", end_quarter="2023Q4", market="hk")
```

---

## CALENDAR (日历) 测试用例

### #32 查询交易日列表 ✅

- **优先级**: P0
- **命令**: `calendar trading-dates`
- **API Method**: `get_trading_dates`
- **结果**: ✅ 通过 (数据匹配 (22 行))

**Payload:**
```json
{
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.get_trading_dates("2024-01-01", "2024-01-31")
```

---

### #33 查询前一交易日 ✅

- **优先级**: P0
- **命令**: `calendar prev`
- **API Method**: `get_previous_trading_date`
- **结果**: ✅ 通过 (数据匹配 (1 行))

**Payload:**
```json
{
  "date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.get_previous_trading_date("2024-01-31")
```

---

### #34 查询后一交易日 ✅

- **优先级**: P0
- **命令**: `calendar next`
- **API Method**: `get_next_trading_date`
- **结果**: ✅ 通过 (数据匹配 (1 行))

**Payload:**
```json
{
  "date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.get_next_trading_date("2024-01-31")
```

---

### #35 查询全年交易日 ✅

- **优先级**: P1
- **命令**: `calendar trading-dates`
- **API Method**: `get_trading_dates`
- **结果**: ✅ 通过 (数据匹配 (242 行))

**Payload:**
```json
{
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}
```

**Python 对照:**
```python
rqdatac.get_trading_dates("2024-01-01", "2024-12-31")
```

---

### #36 节假日前一交易日 ✅

- **优先级**: P1
- **命令**: `calendar prev`
- **API Method**: `get_previous_trading_date`
- **结果**: ✅ 通过 (数据匹配 (1 行))

**Payload:**
```json
{
  "date": "2024-02-10"
}
```

**Python 对照:**
```python
rqdatac.get_previous_trading_date("2024-02-10")
```

---

## FUND (基金) 测试用例

### #37 查询基金信息 ✅

- **优先级**: P0
- **命令**: `fund instruments`
- **API Method**: `fund.all_instruments`
- **结果**: ✅ 通过 (数据匹配 (1 行))

**Payload:**
```json
{
  "order_book_ids": [
    "000001"
  ]
}
```

**Python 对照:**
```python
rqdatac.fund.instruments("000001")
```

---

### #38 查询基金净值 ✅

- **优先级**: P0
- **命令**: `fund nav`
- **API Method**: `fund.get_nav`
- **结果**: ✅ 通过 (数据匹配 (4 行, 5 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000001"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.fund.get_nav("000001", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #39 多基金净值查询 ✅

- **优先级**: P1
- **命令**: `fund nav`
- **API Method**: `fund.get_nav`
- **结果**: ✅ 通过 (数据匹配 (8 行, 5 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000001",
    "110022"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.fund.get_nav(["000001", "110022"], start_date="2024-01-02", end_date="2024-01-05", expect_df=True)
```

---

### #40 查询基金持仓 ✅

- **优先级**: P1
- **命令**: `fund holdings`
- **API Method**: `fund.get_holdings`
- **结果**: ✅ 通过 (数据匹配 (202 行, 8 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000001"
  ],
  "date": "2023-12-31"
}
```

**Python 对照:**
```python
rqdatac.fund.get_holdings("000001", date="2023-12-31")
```

---

### #41 查询所有基金 ✅

- **优先级**: P2
- **命令**: `fund instruments`
- **API Method**: `fund.all_instruments`
- **结果**: ✅ 通过

**Payload:**
```json
{}
```

**Python 对照:**
```python
rqdatac.fund.all_instruments()
```

---

### #42 ETF行情查询 ✅

- **优先级**: P1
- **命令**: `fund price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (4 行, 11 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "510050.XSHG"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price("510050.XSHG", "2024-01-02", "2024-01-05", "1d")
```

---

### #43 查询基金指标 ✅

- **优先级**: P2
- **命令**: `fund indicators`
- **API Method**: `fund.get_indicators`
- **结果**: ✅ 通过 (数据匹配 (22 行, 228 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000001"
  ],
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.fund.get_indicators("000001", start_date="2024-01-01", end_date="2024-01-31")
```

---

### #44 查询基金分红 ✅

- **优先级**: P2
- **命令**: `fund dividend`
- **API Method**: `fund.get_dividend`
- **结果**: ✅ 通过 (数据匹配 (25 行, 3 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000001"
  ]
}
```

**Python 对照:**
```python
rqdatac.fund.get_dividend("000001")
```

---

### #68 查询基金经理 ✅

- **优先级**: P2
- **命令**: `fund manager`
- **API Method**: `fund.get_manager`
- **结果**: ✅ 通过 (数据匹配 (19 行, 6 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "000001"
  ]
}
```

**Python 对照:**
```python
rqdatac.fund.get_manager("000001")
```

---

## FUTURES (期货) 测试用例

### #45 查询期货合约信息 ✅

- **优先级**: P1
- **命令**: `futures instruments`
- **API Method**: `all_instruments`
- **结果**: ✅ 通过 (数据匹配 (1 行))

**Payload:**
```json
{
  "order_book_ids": [
    "CU2406"
  ]
}
```

**Python 对照:**
```python
rqdatac.instruments("CU2406")
```

---

### #46 查询所有期货合约 ✅

- **优先级**: P2
- **命令**: `futures instruments`
- **API Method**: `all_instruments`
- **结果**: ✅ 通过 (数据匹配 (10816 行, 19 个共同字段))

**Payload:**
```json
{}
```

**Python 对照:**
```python
rqdatac.all_instruments(type="Future")
```

---

### #47 期货日线行情 ✅

- **优先级**: P1
- **命令**: `futures price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (4 行, 13 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "CU2406"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price("CU2406", "2024-01-02", "2024-01-05", "1d")
```

---

### #48 多合约期货行情 ✅

- **优先级**: P1
- **命令**: `futures price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (8 行, 13 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "CU2406",
    "AL2406"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price(["CU2406", "AL2406"], "2024-01-02", "2024-01-05", "1d")
```

---

### #49 查询主力合约 ✅

- **优先级**: P1
- **命令**: `futures dominant`
- **API Method**: `futures.get_dominant`
- **结果**: ✅ 通过 (数据匹配 (22 行))

**Payload:**
```json
{
  "underlying_symbol": "CU",
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.futures.get_dominant("CU", start_date="2024-01-01", end_date="2024-01-31")
```

---

### #50 查询主力连续行情 ✅

- **优先级**: P1
- **命令**: `futures dominant-price`
- **API Method**: `futures.get_dominant_price`
- **结果**: ✅ 通过 (数据匹配 (4 行, 14 个共同字段))

**Payload:**
```json
{
  "underlying_symbol": "CU",
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.futures.get_dominant_price("CU", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #51 期货分钟线 ✅

- **优先级**: P1
- **命令**: `futures price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (225 行, 8 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "CU2406"
  ],
  "start_date": "2024-01-02 09:00:00",
  "end_date": "2024-01-02 10:00:00",
  "frequency": "1m"
}
```

**Python 对照:**
```python
rqdatac.get_price("CU2406", "2024-01-02 09:00:00", "2024-01-02 10:00:00", "1m")
```

---

### #52 多品种主力连续行情 ✅

- **优先级**: P2
- **命令**: `futures dominant-price`
- **API Method**: `futures.get_dominant_price`
- **结果**: ✅ 通过 (数据匹配 (22 行, 14 个共同字段))

**Payload:**
```json
{
  "underlying_symbols": "CU",
  "start_date": "2024-01-01",
  "end_date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.futures.get_dominant_price("CU", start_date="2024-01-01", end_date="2024-01-31")
```

---

## OPTIONS (期权) 测试用例

### #53 查询期权合约信息 ✅

- **优先级**: P2
- **命令**: `options contracts`
- **API Method**: `options.get_contracts`
- **结果**: ✅ 通过 (数据匹配 (170 行))

**Payload:**
```json
{
  "underlying": "510050.XSHG",
  "date": "2024-01-02"
}
```

**Python 对照:**
```python
rqdatac.options.get_contracts("510050.XSHG", trading_date="2024-01-02")
```

---

### #54 期权日线行情 ✅

- **优先级**: P2
- **命令**: `options price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (4 行, 15 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "10005765"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price("10005765", "2024-01-02", "2024-01-05", "1d")
```

---

### #55 查询期权希腊字母 ✅

- **优先级**: P2
- **命令**: `options greeks`
- **API Method**: `options.get_greeks`
- **结果**: ✅ 通过 (数据匹配 (4 行, 6 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "10005765"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.options.get_greeks("10005765", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #56 查询期权合约链 ✅

- **优先级**: P2
- **命令**: `options contracts`
- **API Method**: `options.get_contracts`
- **结果**: ✅ 通过 (数据匹配 (146 行))

**Payload:**
```json
{
  "underlying": "510050.XSHG",
  "date": "2024-01-31"
}
```

**Python 对照:**
```python
rqdatac.options.get_contracts("510050.XSHG", trading_date="2024-01-31")
```

---

### #57 多期权合约行情 ✅

- **优先级**: P2
- **命令**: `options price`
- **API Method**: `get_price`
- **结果**: ✅ 通过 (数据匹配 (8 行, 15 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "10005765",
    "10005766"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.get_price(["10005765", "10005766"], "2024-01-02", "2024-01-05", "1d")
```

---

### #58 查询期权合约属性 ✅

- **优先级**: P2
- **命令**: `options contract-property`
- **API Method**: `options.get_contract_property`
- **结果**: ✅ 通过 (数据匹配 (4 行, 4 个共同字段))

**Payload:**
```json
{
  "order_book_ids": [
    "10005765"
  ],
  "start_date": "2024-01-02",
  "end_date": "2024-01-05"
}
```

**Python 对照:**
```python
rqdatac.options.get_contract_property("10005765", start_date="2024-01-02", end_date="2024-01-05")
```

---

### #59 按行权价筛选 ✅

- **优先级**: P2
- **命令**: `options contracts`
- **API Method**: `options.get_contracts`
- **结果**: ✅ 通过 (数据匹配 (8 行))

**Payload:**
```json
{
  "underlying": "510050.XSHG",
  "date": "2024-01-31",
  "strike": 3.0
}
```

**Python 对照:**
```python
rqdatac.options.get_contracts("510050.XSHG", strike=3.0, trading_date="2024-01-31")
```

---

## MACRO (宏观) 测试用例

### #60 查询宏观因子数据 ✅

- **优先级**: P2
- **命令**: `macro factors`
- **API Method**: `econ.get_factors`
- **结果**: ✅ 通过 (数据匹配 (12 行, 4 个共同字段))

**Payload:**
```json
{
  "factors": "居民消费价格指数CPI_当月同比(上年同月=100)",
  "start_date": "2023-01-01",
  "end_date": "2023-12-31"
}
```

**Python 对照:**
```python
rqdatac.econ.get_factors("居民消费价格指数CPI_当月同比(上年同月=100)", "2023-01-01", "2023-12-31")
```

---

### #61 查询存款准备金率 ✅

- **优先级**: P2
- **命令**: `macro reserve-ratio`
- **API Method**: `econ.get_reserve_ratio`
- **结果**: ✅ 通过 (数据匹配 (4 行, 4 个共同字段))

**Payload:**
```json
{
  "start_date": "2023-01-01",
  "end_date": "2023-12-31"
}
```

**Python 对照:**
```python
rqdatac.econ.get_reserve_ratio(start_date="2023-01-01", end_date="2023-12-31")
```

---

### #62 查询货币供应量 ✅

- **优先级**: P2
- **命令**: `macro money-supply`
- **API Method**: `econ.get_money_supply`
- **结果**: ✅ 通过 (数据匹配 (13 行, 7 个共同字段))

**Payload:**
```json
{
  "start_date": "2023-01-01",
  "end_date": "2023-12-31"
}
```

**Python 对照:**
```python
rqdatac.econ.get_money_supply(start_date="2023-01-01", end_date="2023-12-31")
```

---

### #72 查询GDP数据 ✅

- **优先级**: P2
- **命令**: `macro gdp`
- **API Method**: `econ.get_factors`
- **结果**: ✅ 通过 (数据匹配 (4 行, 4 个共同字段))

**Payload:**
```json
{
  "start_date": "2023-01-01",
  "end_date": "2023-12-31"
}
```

**Python 对照:**
```python
rqdatac.econ.get_factors("国内生产总值GDP_累计同比", "2023-01-01", "2023-12-31")
```

---

### #73 查询PPI数据 ✅

- **优先级**: P2
- **命令**: `macro price-ppi`
- **API Method**: `econ.get_factors`
- **结果**: ✅ 通过 (数据匹配 (12 行, 4 个共同字段))

**Payload:**
```json
{
  "start_date": "2023-01-01",
  "end_date": "2023-12-31"
}
```

**Python 对照:**
```python
rqdatac.econ.get_factors("工业品出厂价格指数PPI_当月同比_(上年同月=100)", "2023-01-01", "2023-12-31")
```

---

### #74 查询CPI数据 ✅

- **优先级**: P2
- **命令**: `macro price-cpi`
- **API Method**: `econ.get_factors`
- **结果**: ✅ 通过 (数据匹配 (12 行, 4 个共同字段))

**Payload:**
```json
{
  "start_date": "2023-01-01",
  "end_date": "2023-12-31"
}
```

**Python 对照:**
```python
rqdatac.econ.get_factors("居民消费价格指数CPI_当月同比(上年同月=100)", "2023-01-01", "2023-12-31")
```

---

## 🚀 测试执行

### 运行全部测试
```bash
# 设置环境变量
export RQDATA_USERNAME=your_username
export RQDATA_PASSWORD=your_password

# 编译 CLI
cd build && cmake .. && make -j4

# 运行测试
cd .. && python3 tests/run_complete_tests.py
```

### 环境要求
- RQData 账户（设置 RQDATA_USERNAME / RQDATA_PASSWORD）
- Python 3.9+ 及依赖: `rqdatac`, `pandas`, `numpy`
- C++ 编译环境: CMake 3.14+, C++17

**文档版本**: v3.0
**最后更新**: 2026-03-12