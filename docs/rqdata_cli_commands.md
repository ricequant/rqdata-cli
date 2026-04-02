# RQData CLI 命令参考手册

> 自动生成于 commands.json

## 命令统计

- **总命令数**: 66 个
- **命令类别**: 7 个

## 目录

- [INDEX - 指数数据](#index)
- [STOCK - 股票数据](#stock)
- [CALENDAR - 交易日历](#calendar)
- [FUND - 基金数据](#fund)
- [FUTURES - 期货数据](#futures)
- [OPTIONS - 期权数据](#options)
- [MACRO - 宏观数据](#macro)

---


## INDEX - 指数数据

### `index instruments`

**描述**: 指数基础信息

**API Method**: `instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list。 中国市场的 order_book_id 通常类似'000001.XSHE'。需要注意，国内股票、ETF、指数合约代码分别应当以'.XSHG'或'.XSHE'结尾，前者代表上证，后者代表深证。 比如查询平安银行这个股票合约，则键入'000001.XSHE'，前面的数字部分为交易所内这个股票的合约代码，后半部分为对应的交易所代码。 期货则无此要求 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `index list`

**描述**: 指数列表

**API Method**: `all_instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `type` | string | array<string> | ❌ | 需要查询合约类型，例如：type='CS'代表股票。默认是所有类型 - 'CS' : Common Stock, 即股票 - 'ETF' : Exchange Traded Fund, 即交易所交易基金 - 'LOF' : Listed Open-Ended Fund，即上市型开放式基金 （以下分级基金已并入） - 'INDX' : Index, 即指数 - 'Future' : Futures，即期货，包含股指、国债和商品期货 - 'Spot' : Spot，即现货，目前包括上海黄金交易所现货合约 - 'Option' : 期权，包括目前国内已上市的全部期权合约 - 'Convertible' : 沪深两市场内有交易的可转债合约 - 'Repo' : 沪深两市交易所交易的回购合约 - 'REITs' : 不动产投资信托基金 - 'FUND' : 包括ETF, LOF, REITs后的其他基金 |
| `date` | string | ❌ | 指定日期，筛选指定日期可交易的合约 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `index price`

**描述**: 指数日行情

**API Method**: `get_price`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期 |
| `frequency` | string | ❌ | 历史数据的频率。 现在支持周/日/分钟/tick 级别的历史数据，默认为'1d'。 1m - 分钟线 1d - 日线 1w - 周线，只支持'1w' 日线和分钟可选取不同频率，例如'5m'代表 5 分钟线。 |
| `fields` | string | array<string> | ❌ | 字段名称 |
| `adjust_type` | string | ❌ | 权息修复方案，仅对股票和 ETF 有效，默认为'pre'。 不复权 - 'none'， 前复权 - 'pre'，后复权 - 'post'， 前复权 - 'pre_volume', 后复权 - 'post_volume' 两组前后复权方式仅 volume 字段处理不同，其他字段相同。其中'pre'、'post'中的 volume 采用拆分因子调整；'pre_volume'、'post_volume'中的 volume 采用复权因子调整。 |
| `skip_suspended` | boolean | ❌ | 是否跳过停牌数据。默认为 False，不跳过，用停牌前数据进行补齐。True 则为跳过停牌期。 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回原有的数据结构,周线数据需设置 expect_df=True |
| `time_slice` | string | ❌ | 开始、结束时间段。默认返回当天所有数据。 支持分钟 / tick 级别的切分，详见下方范例。 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `index constituents`

**描述**: 指数成分股

**API Method**: `index_components`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_id` | string | ✅ | 指数代码，传入 order_book_id，例如'000001.XSHG' |
| `date` | string | ❌ | 查询日期，默认为当天 |
| `start_date` | string | ❌ | 指定开始日期，不能和 date 同时指定 |
| `end_date` | string | ❌ | 指定结束日期, 需和 start_date 同时指定并且应当不小于开始日期 |
| `return_create_tm` | boolean | ❌ | 设置为 True 的时候，传入 date, 返回 tuple, 第一个元素是列表, 第二个元素是入库时间 ; 传入 start_date, end_date, 返回 dict, 其中 key 是日期, value 是 tuple, 其中第一个元素是列表, 第二个元素是入库时间 |
| `market` | string | ❌ | 默认是中国市场('cn')，目前仅支持中国市场 |

---

### `index weights`

**描述**: 成分股权重

**API Method**: `index_weights`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_id` | string | ✅ | 指数代码，可传入 order_book_id，例如'000001.XSHG'或'沪深 300'。目前所支持的指数列表可以参考指数数据表 |
| `date` | string | ❌ | 查询日期，默认为当天 |
| `start_date` | string | ❌ | 指定开始日期，不能和 date 同时指定 |
| `end_date` | string | ❌ | 指定结束日期, 需和 start_date 同时指定并且应当不小于开始日期 |
| `market` | string | ❌ | 地区代码, 如 'cn' |

---


## STOCK - 股票数据

### `stock cn list`

**描述**: 股票列表

**API Method**: `all_instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `type` | string | array<string> | ❌ | 需要查询合约类型，例如：type='CS'代表股票。默认是所有类型 - 'CS' : Common Stock, 即股票 - 'ETF' : Exchange Traded Fund, 即交易所交易基金 - 'LOF' : Listed Open-Ended Fund，即上市型开放式基金 （以下分级基金已并入） - 'INDX' : Index, 即指数 - 'Future' : Futures，即期货，包含股指、国债和商品期货 - 'Spot' : Spot，即现货，目前包括上海黄金交易所现货合约 - 'Option' : 期权，包括目前国内已上市的全部期权合约 - 'Convertible' : 沪深两市场内有交易的可转债合约 - 'Repo' : 沪深两市交易所交易的回购合约 - 'REITs' : 不动产投资信托基金 - 'FUND' : 包括ETF, LOF, REITs后的其他基金 |
| `date` | string | ❌ | 指定日期，筛选指定日期可交易的合约 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn instruments`

**描述**: A股基础信息

**API Method**: `instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list。 中国市场的 order_book_id 通常类似'000001.XSHE'。需要注意，国内股票、ETF、指数合约代码分别应当以'.XSHG'或'.XSHE'结尾，前者代表上证，后者代表深证。 比如查询平安银行这个股票合约，则键入'000001.XSHE'，前面的数字部分为交易所内这个股票的合约代码，后半部分为对应的交易所代码。 期货则无此要求 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn price`

**描述**: A股日行情

**API Method**: `get_price`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期 |
| `frequency` | string | ❌ | 历史数据的频率。 现在支持周/日/分钟/tick 级别的历史数据，默认为'1d'。 1m - 分钟线 1d - 日线 1w - 周线，只支持'1w' 日线和分钟可选取不同频率，例如'5m'代表 5 分钟线。 |
| `fields` | string | array<string> | ❌ | 字段名称 |
| `adjust_type` | string | ❌ | 权息修复方案，仅对股票和 ETF 有效，默认为'pre'。 不复权 - 'none'， 前复权 - 'pre'，后复权 - 'post'， 前复权 - 'pre_volume', 后复权 - 'post_volume' 两组前后复权方式仅 volume 字段处理不同，其他字段相同。其中'pre'、'post'中的 volume 采用拆分因子调整；'pre_volume'、'post_volume'中的 volume 采用复权因子调整。 |
| `skip_suspended` | boolean | ❌ | 是否跳过停牌数据。默认为 False，不跳过，用停牌前数据进行补齐。True 则为跳过停牌期。 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回原有的数据结构,周线数据需设置 expect_df=True |
| `time_slice` | string | ❌ | 开始、结束时间段。默认返回当天所有数据。 支持分钟 / tick 级别的切分，详见下方范例。 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn dividend`

**描述**: A股分红送股

**API Method**: `get_dividend`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可输入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，不传入 start_date ,end_date 则 默认返回全部分红数据 |
| `adjusted` | boolean | ❌ |  |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe,如果调为 False ,则返回原有的数据结构 |
| `market` | string | ❌ | 默认是中国内地市场('cn')。cn-中国内地市场，hk-中国香港市场 |

---

### `stock cn split`

**描述**: A股拆股合股

**API Method**: `get_split`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可输入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期，默认返回全部 |
| `end_date` | string | ❌ | 结束日期 ，默认返回全部 |
| `market` | string | ❌ | 默认是中国内地市场('cn')。cn-中国内地市场，hk-中国香港市场 |

---

### `stock cn shares`

**描述**: A股股本数据

**API Method**: `get_shares`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，不传入 start_date ,end_date 则 默认返回最近三个月的数据 |
| `fields` | string | array<string> | ❌ | 默认为所有字段。见下方列表 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe,如果调为 False ,则返回原有的数据结构 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn ex-factor`

**描述**: A股复权因子

**API Method**: `get_ex_factor`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可输入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期，默认返回全部 |
| `end_date` | string | ❌ | 结束日期，默认返回全部 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn industry`

**描述**: A股行业分类

**API Method**: `get_instrument_industry`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 股票合约代码，可输入 order_book_id, order_book_id list |
| `source` | string | ❌ | 分类依据。citics_2019 - 中信新分类（2019 发布）, citics - 中信旧分类（退役中）, gildata -聚源。 默认 source='citics_2019'. |
| `level` | integer | ❌ | 行业分类级别，共三级，默认返回一级分类。参数 0,1,2,3 一一对应，其中 0 返回三级分类完整情况 当 source='citics_2019' 时，level 可传入'citics_sector' 获取该股票的衍生板块及风格归属 |
| `date` | string | ❌ | 查询日期，默认为当前最新日期 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn turnover-rate`

**描述**: A股换手率

**API Method**: `get_turnover_rate`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可输入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，不传入 start_date ,end_date 则 默认返回最近三个月的数据 |
| `fields` | string | array<string> | ❌ | 默认为所有字段。当天换手率 - `today`，过去一周平均换手率 - `week`，过去一个月平均换手率 - `month`，过去一年平均换手率 - `year`，当年平均换手率 - `current_year` |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回原有的数据结构 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn margin`

**描述**: A股融资融券

**API Method**: `get_margin_stocks`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `date` | string | ❌ | 查询日期，默认为今天上一交易日 |
| `exchange` | string | ❌ | 交易所，默认为 None，返回所有字段。可选字段包括：'XSHE', 'sz' 代表深交所；'XSHG', 'sh' 代表上交所 |
| `margin_type` | string | ❌ | 'stock' 代表融券卖出，'cash'，代表融资买入，默认为'stock' |
| `market` | string | ❌ | 默认是中国内地市场('cn') |

---

### `stock cn northbound`

**描述**: 北向资金持股

**API Method**: `get_stock_connect`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | ✅ | 可输入 order_book_id 或 symbol。另， 1、输入'shanghai_connect'可返回沪股通的全部股票数据。 2、输入'shenzhen_connect'可返回深股通的全部股票数据。 3、输入'all_connect'可返回沪股通、深股通的全部股票数据。 |
| `start_date` | string | ❌ | 开始日期，默认为'2017-03-17' |
| `end_date` | string | ❌ | 结束日期，默认为'2018-03-16' |
| `fields` | string | array<string> | ❌ | 默认为所有字段。见下方列表 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回原有的数据结构 |

---

### `stock cn financial`

**描述**: A股财务数据（PIT）

**API Method**: `get_pit_financials_ex`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `fields` | array<string> | ✅ | 需要返回的财务字段。 |
| `start_quarter` | string | ✅ | 财报回溯查询的起始报告期，例如'2015q2'代表 2015 年半年报 。 |
| `end_quarter` | string | ✅ | 财报回溯查询的截止报告期，例如'2015q4'代表 2015 年年报。 |
| `date` | string | ❌ | 查询日期，默认查询日期为当前最新日期 |
| `statements` | string | ❌ | 基于查询日期，返回某一个报告期的所有记录或最新一条记录，设置 statements 为 all 时返回所有记录，statements 等于 latest 时返回最新的一条记录，默认为 latest. |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn financial-indicator`

**描述**: A股财务指标

**API Method**: `get_factor`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `factor` | string | array<string> | ✅ | 因子名称，可查询 get_all_factor_names() 得到所有有效因子字段 |
| `start_date` | string | ❌ | 开始日期。注：如使用开始日期，则必填结束日期 |
| `end_date` | string | ❌ | 结束日期。注：若使用结束日期，则开始日期必填 |
| `universe` | string | ❌ | 指定因子计算时的股票域，米筐所有公共因子均在全市场范围计算，此参数保留为 None 即可 .. deprecated:: 2 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回 原有的数据结构 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn financial-express`

**描述**: A股业绩快报

**API Method**: `current_performance`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码。 |
| `info_date` | string | ❌ | 公告日期。如果不填(info_date 和 quarter 都为空)，则返回当前日期的最新发布的快报。如果填写，则从 info_date 当天或者之前最新的报告开始抓取。 |
| `quarter` | string | ❌ | info_date 参数优先级高于 quarter。如果 info_date 填写了日期，则不查看 quarter 这个字段。 如果 info_date 没有填写而 quarter 有填写，则财报回溯查询的起始报告期，例如'2015q2', '2015q4'分别代表 2015 年半年报以及年报。默认只获取当前报告期财务信息 |
| `interval` | string | ❌ | 查询财务数据的间隔。例如，填写'5y'，则代表从报告期开始回溯 5 年，每年为相同报告期数据；'3q'则代表从报告期开始向前回溯 3 个季度。不填写默认抓取一期。 |
| `fields` | string | array<string> | ❌ | 抓取对应有效字段返回。默认返回所有字段。具体快报字段见下方。 |
| `market` | string | ❌ | 默认是中国内地市场('cn') |

---

### `stock cn forecast`

**描述**: A股业绩预告

**API Method**: `performance_forecast`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list。 |
| `info_date` | string | ❌ | 公告日期。如果不填(info_date 和 end_date 都为空)，则返回当前日期的最新发布的业绩预告。如果填写，则从 info_date 当天或者之前最新的报告开始抓取。注：info_date 优先级高于 end_date |
| `end_date` | string | ❌ | 对应财务预告期末日期，如'20150331'。 |
| `fields` | string | array<string> | ❌ | 抓取对应有效字段返回。默认返回所有字段。具体业绩预告字段见下方 |
| `market` | string | ❌ | 默认是中国内地市场('cn') |

---

### `stock cn consensus`

**描述**: A股一致预期

**API Method**: `get_factor`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `factor` | string | array<string> | ✅ | 因子名称，可查询 get_all_factor_names() 得到所有有效因子字段 |
| `start_date` | string | ❌ | 开始日期。注：如使用开始日期，则必填结束日期 |
| `end_date` | string | ❌ | 结束日期。注：若使用结束日期，则开始日期必填 |
| `universe` | string | ❌ | 指定因子计算时的股票域，米筐所有公共因子均在全市场范围计算，此参数保留为 None 即可 .. deprecated:: 2 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回 原有的数据结构 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn shareholder-top10`

**描述**: 十大股东

**API Method**: `get_main_shareholder`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期，默认为去年当日。 |
| `end_date` | string | ❌ | 结束日期，默认为查询当日。 |
| `is_total` | boolean | ❌ | 默认为 False, 即基于持有 A 股流通股。若为 True 则基于所有发行出的 A 股。 |
| `start_rank` | integer | ❌ | 排名开始值 |
| `end_rank` | integer | ❌ | 排名结束值 ,start_rank ,end_rank 不传参时返回全部的十位股东名单 |
| `market` | string | ❌ | 市场，默认'cn'为中国内地市场。 |

---

### `stock cn shareholder-top10-free`

**描述**: 十大流通股东

**API Method**: `get_main_shareholder`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期，默认为去年当日。 |
| `end_date` | string | ❌ | 结束日期，默认为查询当日。 |
| `is_total` | boolean | ❌ | 默认为 False, 即基于持有 A 股流通股。若为 True 则基于所有发行出的 A 股。 |
| `start_rank` | integer | ❌ | 排名开始值 |
| `end_rank` | integer | ❌ | 排名结束值 ,start_rank ,end_rank 不传参时返回全部的十位股东名单 |
| `market` | string | ❌ | 市场，默认'cn'为中国内地市场。 |

---

### `stock cn announcement`

**描述**: 公告信息

**API Method**: `get_announcement`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，给出单个或多个 order_book_id |
| `start_date` | string | ❌ | 开始日期。注：如使用开始日期，则必填结束日期 |
| `end_date` | string | ❌ | 结束日期。注：若使用结束日期，则开始日期必填 |
| `fields` | string | array<string> | ❌ | 可选字段见下方返回，若不指定，则默认获取所有字段 |
| `market` | string | ❌ | 默认是中国内地市场('cn') |

---

### `stock cn suspended`

**描述**: 停牌信息

**API Method**: `is_suspended`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码。传入单只或多支股票的 order_book_id |
| `start_date` | string | ❌ | 开始日期，默认为股票上市日期 |
| `end_date` | string | ❌ | 结束日期，默认为当前日期，如果股票已经退市，则为退市日期 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock cn st-status`

**描述**: ST状态

**API Method**: `is_st_stock`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期，默认为股票上市日期 |
| `end_date` | string | ❌ | 结束日期，默认为当前日期，如果股票已经退市，则为退市日期 |
| `market` | string | ❌ | 默认是中国内地市场('cn') |

---

### `stock hk list`

**描述**: 港股列表

**API Method**: `all_instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `type` | string | array<string> | ❌ | 需要查询合约类型，例如：type='CS'代表股票。默认是所有类型 - 'CS' : Common Stock, 即股票 - 'ETF' : Exchange Traded Fund, 即交易所交易基金 - 'LOF' : Listed Open-Ended Fund，即上市型开放式基金 （以下分级基金已并入） - 'INDX' : Index, 即指数 - 'Future' : Futures，即期货，包含股指、国债和商品期货 - 'Spot' : Spot，即现货，目前包括上海黄金交易所现货合约 - 'Option' : 期权，包括目前国内已上市的全部期权合约 - 'Convertible' : 沪深两市场内有交易的可转债合约 - 'Repo' : 沪深两市交易所交易的回购合约 - 'REITs' : 不动产投资信托基金 - 'FUND' : 包括ETF, LOF, REITs后的其他基金 |
| `date` | string | ❌ | 指定日期，筛选指定日期可交易的合约 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock hk instruments`

**描述**: 港股基础信息

**API Method**: `instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list。 中国市场的 order_book_id 通常类似'000001.XSHE'。需要注意，国内股票、ETF、指数合约代码分别应当以'.XSHG'或'.XSHE'结尾，前者代表上证，后者代表深证。 比如查询平安银行这个股票合约，则键入'000001.XSHE'，前面的数字部分为交易所内这个股票的合约代码，后半部分为对应的交易所代码。 期货则无此要求 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock hk price`

**描述**: 港股日行情

**API Method**: `get_price`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期 |
| `frequency` | string | ❌ | 历史数据的频率。 现在支持周/日/分钟/tick 级别的历史数据，默认为'1d'。 1m - 分钟线 1d - 日线 1w - 周线，只支持'1w' 日线和分钟可选取不同频率，例如'5m'代表 5 分钟线。 |
| `fields` | string | array<string> | ❌ | 字段名称 |
| `adjust_type` | string | ❌ | 权息修复方案，仅对股票和 ETF 有效，默认为'pre'。 不复权 - 'none'， 前复权 - 'pre'，后复权 - 'post'， 前复权 - 'pre_volume', 后复权 - 'post_volume' 两组前后复权方式仅 volume 字段处理不同，其他字段相同。其中'pre'、'post'中的 volume 采用拆分因子调整；'pre_volume'、'post_volume'中的 volume 采用复权因子调整。 |
| `skip_suspended` | boolean | ❌ | 是否跳过停牌数据。默认为 False，不跳过，用停牌前数据进行补齐。True 则为跳过停牌期。 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回原有的数据结构,周线数据需设置 expect_df=True |
| `time_slice` | string | ❌ | 开始、结束时间段。默认返回当天所有数据。 支持分钟 / tick 级别的切分，详见下方范例。 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock hk ex-factor`

**描述**: 港股复权因子

**API Method**: `get_ex_factor`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可输入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期，默认返回全部 |
| `end_date` | string | ❌ | 结束日期，默认返回全部 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock hk shares`

**描述**: 港股股本数据

**API Method**: `get_shares`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，不传入 start_date ,end_date 则 默认返回最近三个月的数据 |
| `fields` | string | array<string> | ❌ | 默认为所有字段。见下方列表 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe,如果调为 False ,则返回原有的数据结构 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock hk industry`

**描述**: 港股行业分类

**API Method**: `get_instrument_industry`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 股票合约代码，可输入 order_book_id, order_book_id list |
| `source` | string | ❌ | 分类依据。citics_2019 - 中信新分类（2019 发布）, citics - 中信旧分类（退役中）, gildata -聚源。 默认 source='citics_2019'. |
| `level` | integer | ❌ | 行业分类级别，共三级，默认返回一级分类。参数 0,1,2,3 一一对应，其中 0 返回三级分类完整情况 当 source='citics_2019' 时，level 可传入'citics_sector' 获取该股票的衍生板块及风格归属 |
| `date` | string | ❌ | 查询日期，默认为当前最新日期 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock hk turnover-rate`

**描述**: 港股换手率

**API Method**: `get_turnover_rate`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可输入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，不传入 start_date ,end_date 则 默认返回最近三个月的数据 |
| `fields` | string | array<string> | ❌ | 默认为所有字段。当天换手率 - `today`，过去一周平均换手率 - `week`，过去一个月平均换手率 - `month`，过去一年平均换手率 - `year`，当年平均换手率 - `current_year` |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回原有的数据结构 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `stock hk dividend`

**描述**: 港股分红

**API Method**: `get_dividend`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可输入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，不传入 start_date ,end_date 则 默认返回全部分红数据 |
| `adjusted` | boolean | ❌ |  |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe,如果调为 False ,则返回原有的数据结构 |
| `market` | string | ❌ | 默认是中国内地市场('cn')。cn-中国内地市场，hk-中国香港市场 |

---

### `stock hk financial`

**描述**: 港股财务数据（PIT）

**API Method**: `get_pit_financials_ex`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `fields` | array<string> | ✅ | 需要返回的财务字段。 |
| `start_quarter` | string | ✅ | 财报回溯查询的起始报告期，例如'2015q2'代表 2015 年半年报 。 |
| `end_quarter` | string | ✅ | 财报回溯查询的截止报告期，例如'2015q4'代表 2015 年年报。 |
| `date` | string | ❌ | 查询日期，默认查询日期为当前最新日期 |
| `statements` | string | ❌ | 基于查询日期，返回某一个报告期的所有记录或最新一条记录，设置 statements 为 all 时返回所有记录，statements 等于 latest 时返回最新的一条记录，默认为 latest. |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---


## CALENDAR - 交易日历

### `calendar trading-dates`

**描述**: 交易日列表

**API Method**: `get_trading_dates`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `start_date` | string | ✅ | 开始日期 |
| `end_date` | string | ✅ | 结束日期 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `calendar prev`

**描述**: 上一个交易日

**API Method**: `get_previous_trading_date`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `date` | string | ✅ | 指定日期 |
| `n` | integer | ❌ | n 代表往前第 n 个交易日。默认为 1，即前一个交易日 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `calendar next`

**描述**: 下一个交易日

**API Method**: `get_next_trading_date`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `date` | string | ✅ | 指定日期 |
| `n` | integer | ❌ | n 代表未来第 n 个交易日。默认为 1，即下一个交易日 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---


## FUND - 基金数据

### `fund list`

**描述**: 基金列表

**API Method**: `fund.all_instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `date` | date | ❌ | 该参数表示过滤掉 listed_date 晚于当日的基金，默认为 None，表示获取全部基金信息，不进行过滤。 (Default value = None) |
| `market` | string | ❌ | (Default value = "cn") |

---

### `fund instruments`

**描述**: 基金基础信息

**API Method**: `fund.instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 基金代码，str 或 list of str |
| `market` | string | ❌ | (Default value = "cn") |

---

### `fund nav`

**描述**: 基金净值

**API Method**: `fund.get_nav`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 基金代码，str 或 list of str |
| `start_date` | date | ❌ | 开始日期 (Default value = None) |
| `end_date` | date | ❌ | 结束日期 (Default value = None) |
| `fields` | string | array<string> | ❌ | str or list of str，例如：'acc_net_value', 'unit_net_value', 'subscribe_status', 'redeem_status', 'change_rate' (Default value = None) |
| `expect_df` | boolean | ❌ | 返回 MultiIndex DataFrame (Default value = False) |
| `market` | string | ❌ | (Default value = "cn") |

---

### `fund holdings`

**描述**: 基金持仓

**API Method**: `fund.get_holdings`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 基金代码，str 或 list of str |
| `date` | date | ❌ | 日期，为空则返回所有持仓 (Default value = None) |
| `market` | string | ❌ | (Default value = "cn") |

---

### `fund indicators`

**描述**: 基金指标

**API Method**: `fund.get_indicators`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 基金代码，str 或 list of str |
| `start_date` | date | ❌ | 开始日期 (Default value = None) |
| `end_date` | date | ❌ | 结束日期 (Default value = None) |
| `fields` | string | array<string> | ❌ | str or list of str (Default value = None) |
| `rule` | integer | ❌ | str, 可选：["ricequant"] (Default value = "ricequant") |
| `indicator_type` | any | ❌ | str, 可选：["value", "rank"] (Default value = "value") |
| `market` | string | ❌ | (Default value = "cn") |

---

### `fund manager`

**描述**: 基金经理

**API Method**: `fund.get_manager`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 基金代码，str 或 list of str |
| `expect_df` | boolean | ❌ | 返回 MultiIndex DataFrame (Default value = True) |
| `market` | string | ❌ | (Default value = "cn") |

---

### `fund dividend`

**描述**: 基金分红

**API Method**: `fund.get_dividend`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 基金代码，str 或 list of str |
| `market` | string | ❌ | (Default value = "cn") |

---


## FUTURES - 期货数据

### `futures list`

**描述**: 期货合约列表

**API Method**: `all_instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `type` | string | array<string> | ❌ | 需要查询合约类型，例如：type='CS'代表股票。默认是所有类型 - 'CS' : Common Stock, 即股票 - 'ETF' : Exchange Traded Fund, 即交易所交易基金 - 'LOF' : Listed Open-Ended Fund，即上市型开放式基金 （以下分级基金已并入） - 'INDX' : Index, 即指数 - 'Future' : Futures，即期货，包含股指、国债和商品期货 - 'Spot' : Spot，即现货，目前包括上海黄金交易所现货合约 - 'Option' : 期权，包括目前国内已上市的全部期权合约 - 'Convertible' : 沪深两市场内有交易的可转债合约 - 'Repo' : 沪深两市交易所交易的回购合约 - 'REITs' : 不动产投资信托基金 - 'FUND' : 包括ETF, LOF, REITs后的其他基金 |
| `date` | string | ❌ | 指定日期，筛选指定日期可交易的合约 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `futures instrument`

**描述**: 期货合约基本信息

**API Method**: `instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list。 中国市场的 order_book_id 通常类似'000001.XSHE'。需要注意，国内股票、ETF、指数合约代码分别应当以'.XSHG'或'.XSHE'结尾，前者代表上证，后者代表深证。 比如查询平安银行这个股票合约，则键入'000001.XSHE'，前面的数字部分为交易所内这个股票的合约代码，后半部分为对应的交易所代码。 期货则无此要求 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `futures price`

**描述**: 期货行情

**API Method**: `get_price`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期 |
| `frequency` | string | ❌ | 历史数据的频率。 现在支持周/日/分钟/tick 级别的历史数据，默认为'1d'。 1m - 分钟线 1d - 日线 1w - 周线，只支持'1w' 日线和分钟可选取不同频率，例如'5m'代表 5 分钟线。 |
| `fields` | string | array<string> | ❌ | 字段名称 |
| `adjust_type` | string | ❌ | 权息修复方案，仅对股票和 ETF 有效，默认为'pre'。 不复权 - 'none'， 前复权 - 'pre'，后复权 - 'post'， 前复权 - 'pre_volume', 后复权 - 'post_volume' 两组前后复权方式仅 volume 字段处理不同，其他字段相同。其中'pre'、'post'中的 volume 采用拆分因子调整；'pre_volume'、'post_volume'中的 volume 采用复权因子调整。 |
| `skip_suspended` | boolean | ❌ | 是否跳过停牌数据。默认为 False，不跳过，用停牌前数据进行补齐。True 则为跳过停牌期。 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回原有的数据结构,周线数据需设置 expect_df=True |
| `time_slice` | string | ❌ | 开始、结束时间段。默认返回当天所有数据。 支持分钟 / tick 级别的切分，详见下方范例。 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `futures dominant`

**描述**: 主力合约

**API Method**: `futures.get_dominant`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `underlying_symbol` | string | ✅ | 期货合约品种，例如沪深 300 股指期货为'IF' |
| `start_date` | string | ❌ | 开始日期，默认为期货品种最早上市日期后一交易日 |
| `end_date` | string | ❌ | 结束日期，默认为当前日期 |
| `rule` | integer | ❌ | 主力合约选取规则。 默认 rule=0，当同品种其他合约持仓量在收盘后超过当前主力合约 1.1 倍时，从第二个交易日开始进行主力合约的切换。每个合约只能做一次主力/次主力合约，不会重复出现。针对股指期货，只在当月和次月选择主力合约。 当 rule=1 时，主力/次主力合约的选取只考虑最大/第二大昨仓这个条件。 当 rule=2 时，采用昨日成交量与持仓量同为最大/第二大的合约为当日主力/次主力。 当 rule=3 时，在 rule=0 选取规则上，考虑在最后一个交易日不能成为主力/次主力合约。 |
| `rank` | integer | ❌ | 默认 rank=1。 1-主力合约（支持所有期货） 2-次主力合约（支持所有期货；针对股指期货，需满足 **rule=1** 或 **rule=2**） 3-次次主力合约（支持所有期货；针对股指期货，需满足 **rule=1** 或 **rule=2**） |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场； |

---

### `futures dominant-price`

**描述**: 主力合约行情

**API Method**: `futures.get_dominant_price`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `underlying_symbols` | string | array<string> | ✅ | 期货合约品种，可传入 underlying_symbol, underlying_symbol list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，start_date ,end_date 不传参数时默认返回最近三个月的数据 |
| `frequency` | string | ❌ | 历史数据的频率。 支持/日/分钟/tick 级别的历史数据，默认为'1d'。1m- 分钟线，1d-日线，分钟可选取不同频率，例如'5m'代表 5 分钟线 |
| `fields` | string | array<string> | ❌ | 查询字段，可选字段见下方返回，默认返回所有字段 |
| `adjust_type` | string | ❌ | 复权方式，默认为'pre'， none - 不复权 ，pre - 前复权， post - 后复权， |
| `adjust_method` | string | ❌ | 复权方法 'prev_close_spread'：基于主力合约切换前一个交易日收盘价价差进行复权 'open_spread'：基于主力合约切换当日开盘价价差进行复权 'prev_close_ratio'：基于主力合约切换前一个交易日收盘价比例进行复权 'open_ratio'：基于主力合约切换当日开盘价比例进行复权' 默认为'prev_close_spread'; adjust_type 为 None 时，adjust_method 复权方法设置无效 |
| `rule` | integer | ❌ | 主力合约选取规则。 默认 rule=0，当同品种其他合约持仓量在收盘后超过当前主力合约 1.1 倍时，从第二个交易日开始进行主力合约的切换。每个合约只能做一次主力/次主力合约，不会重复出现。针对股指期货，只在当月和次月选择主力合约。 当 rule=1 时，主力/次主力合约的选取只考虑最大/第二大昨仓这个条件。 当 rule=2 时，采用昨日成交量与持仓量同为最大/第二大的合约为当日主力/次主力。 当 rule=3 时，在 rule=0 选取规则上，考虑在最后一个交易日不能成为主力/次主力合约。 |
| `rank` | integer | ❌ | 默认 rank=1 1-主力合约（支持所有期货） 2-次主力合约（支持所有期货；针对股指期货，需满足 **rule=1** 或 **rule=2**） 3-次次主力合约（支持所有期货；针对股指期货，需满足 **rule=1** 或 **rule=2**） |
| `time_slice` | string | ❌ | 开始、结束时间段。默认返回当天所有数据。 |

---


## OPTIONS - 期权数据

### `options list`

**描述**: 期权合约列表

**API Method**: `all_instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `type` | string | array<string> | ❌ | 需要查询合约类型，例如：type='CS'代表股票。默认是所有类型 - 'CS' : Common Stock, 即股票 - 'ETF' : Exchange Traded Fund, 即交易所交易基金 - 'LOF' : Listed Open-Ended Fund，即上市型开放式基金 （以下分级基金已并入） - 'INDX' : Index, 即指数 - 'Future' : Futures，即期货，包含股指、国债和商品期货 - 'Spot' : Spot，即现货，目前包括上海黄金交易所现货合约 - 'Option' : 期权，包括目前国内已上市的全部期权合约 - 'Convertible' : 沪深两市场内有交易的可转债合约 - 'Repo' : 沪深两市交易所交易的回购合约 - 'REITs' : 不动产投资信托基金 - 'FUND' : 包括ETF, LOF, REITs后的其他基金 |
| `date` | string | ❌ | 指定日期，筛选指定日期可交易的合约 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `options instruments`

**描述**: 期权合约信息

**API Method**: `instruments`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list。 中国市场的 order_book_id 通常类似'000001.XSHE'。需要注意，国内股票、ETF、指数合约代码分别应当以'.XSHG'或'.XSHE'结尾，前者代表上证，后者代表深证。 比如查询平安银行这个股票合约，则键入'000001.XSHE'，前面的数字部分为交易所内这个股票的合约代码，后半部分为对应的交易所代码。 期货则无此要求 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `options price`

**描述**: 期权行情

**API Method**: `get_price`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码，可传入 order_book_id, order_book_id list |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期 |
| `frequency` | string | ❌ | 历史数据的频率。 现在支持周/日/分钟/tick 级别的历史数据，默认为'1d'。 1m - 分钟线 1d - 日线 1w - 周线，只支持'1w' 日线和分钟可选取不同频率，例如'5m'代表 5 分钟线。 |
| `fields` | string | array<string> | ❌ | 字段名称 |
| `adjust_type` | string | ❌ | 权息修复方案，仅对股票和 ETF 有效，默认为'pre'。 不复权 - 'none'， 前复权 - 'pre'，后复权 - 'post'， 前复权 - 'pre_volume', 后复权 - 'post_volume' 两组前后复权方式仅 volume 字段处理不同，其他字段相同。其中'pre'、'post'中的 volume 采用拆分因子调整；'pre_volume'、'post_volume'中的 volume 采用复权因子调整。 |
| `skip_suspended` | boolean | ❌ | 是否跳过停牌数据。默认为 False，不跳过，用停牌前数据进行补齐。True 则为跳过停牌期。 |
| `expect_df` | boolean | ❌ | 默认返回 pandas dataframe。如果调为 False，则返回原有的数据结构,周线数据需设置 expect_df=True |
| `time_slice` | string | ❌ | 开始、结束时间段。默认返回当天所有数据。 支持分钟 / tick 级别的切分，详见下方范例。 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场；'hk' - 香港市场 |

---

### `options contracts`

**描述**: 期权合约列表

**API Method**: `options.get_contracts`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `underlying` | string | ✅ | 期权标的。可以填写 'M' 代表期货品种的字母；也可填写'M1901' 这种具体合约代码。只支持单一品种或合约代码输入 |
| `option_type` | string | ❌ | 'C' 代表认购期权；'P' 代表认沽期权合约。默认返回全部类型 |
| `maturity` | string | ❌ | 到期月份。例如 1811 代表期权 18 年 11 月到期（而不是标的期货的到期时间）。默认返回全部到期月份 |
| `strike` | number | ❌ | 行权价。查询时向左靠档匹配（例如，当前最高行权价是 1000，则输入大于 1000 的行权价都会向左靠档至 1000）。默认返回全部行权价 |
| `trading_date` | string | ❌ | 查询日期。默认返回全部数据 |

---

### `options greeks`

**描述**: 期权希腊字母

**API Method**: `options.get_greeks`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码 |
| `start_date` | string | ❌ | 查询的开始日期 |
| `end_date` | string | ❌ | 查询的结束日期，默认为当天 |
| `fields` | string | array<string> | ❌ | 查询字段 - 'iv': 从场内期权价格反推对应的标的物收益波动率 - 'delta': 期权价格对于标的物价格变化的一阶敏感度 - 'gamma': 期权价格对于标的物价格变化的二阶敏感度 - 'vega': 期权价格对于隐含波动率的一阶敏感度 - 'theta': 期权价格对于合约待偿期变化的一阶敏感度，为每日历年的 Theta - 'rho': 期权价格对于无风险利率变化的一阶敏感度 |
| `model` | string | ❌ | 计算方法,默认为 'implied_forward'。 针对 BS 模型中的标的物远期利率，提供两种计算方法： last - 以国债同期限收益率作为无风险利率，标的物分红定为 0 计算远期利率； implied_forward - 基于期权的风险平价关系（put-call parity），推算市场数据隐含的标的物远期利率 |
| `price_type` | string | ❌ | 计算使用价格，默认为'close' 。 close - 使用期权收盘价计算衍生指标。 settlement - 使用期权结算价计算衍生指标 |
| `frequency` | string | ❌ | 数据的频率。 默认为'1d'。 1m - 分钟级别（仅支持股指期货,且 price_type 需为'close'） 1d - 日级别 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场； |

---

### `options indicators`

**描述**: 期权指标

**API Method**: `options.get_indicators`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `underlying_symbols` | string | array<string> | ✅ | 期权标的代码，例 'CU' |
| `maturity` | string | ✅ | 到期月份，例 2503 代表期权 25 年 3 月到期 |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，start_date ,end_date 不传参数时默认返回最近三个月的数据 |
| `fields` | string | array<string> | ❌ | 查询字段，可选字段见下方返回，默认返回所有字段 - 'AM_PCR': 成交额 PCR (每日看跌期权成交额 / 每日看涨期权成交额) - 'OI_PCR': 持仓量 PCR (每日看跌期权持仓量 / 每日看涨期权持仓量) - 'VL_PCR': 成交量 PCR (每日看跌期权成交量 / 每日看涨期权成交量) - 'skew': 期权偏度 ((( delta 为 0.25 的认购合约 IV ) - ( delta 为 -0.25 的认沽合约 IV )) / ( delta 为 -0.25 的认沽合约 IV )) - 'iv_025_dela': 0.25delta 隐含波动率 - 'iv_minus_025_dela': -0.25delta 隐含波动率 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场； |

---

### `options dominant-month`

**描述**: 期权主力月份

**API Method**: `options.get_dominant_month`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `underlying_symbol` | string | array<string> | ✅ | 期权标的代码，例'CU' |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，start_date ,end_date 不传参数时默认返回所有数据 |
| `rule` | integer | ❌ | 默认 rule=0，每个月份只能做一次主力月份，不会重复出现。 当 rule=1 时，主力月份的选取只考虑持仓量与成交量条件。 |
| `rank` | integer | ❌ | 默认 rank=1。 1-主力月份，2-次主力月份 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场； |

---

### `options contract-property`

**描述**: 期权合约属性

**API Method**: `options.get_contract_property`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `order_book_ids` | string | array<string> | ✅ | 合约代码 |
| `start_date` | string | ❌ | 开始日期 |
| `end_date` | string | ❌ | 结束日期，start_date ,end_date 不传参数时默认返回所有数据 |
| `fields` | string | array<string> | ❌ | 查询字段，可选字段见下方返回，默认返回所有字段 |
| `market` | string | ❌ | 默认是中国内地市场('cn') 。可选'cn' - 中国内地市场； |

---


## MACRO - 宏观数据

### `macro reserve-ratio`

**描述**: 存款准备金率

**API Method**: `econ.get_reserve_ratio`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `reserve_type` | string | array<string> | ❌ | 存款准备金详细类别，默认为 'all'，目前支持 'all', 'major', 'other' |
| `start_date` | string | ❌ | 开始查找时间，如 '20180501'，默认为上一年的当天 |
| `end_date` | string | ❌ | 截止查找时间，如 '20180501'，默认为当天 |
| `market` | string | ❌ | (Default value = "cn") |

---

### `macro money-supply`

**描述**: 货币供应量

**API Method**: `econ.get_money_supply`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `start_date` | string | ❌ | 开始日期，默认为去年的查询当日（基准为信息公布日）。 |
| `end_date` | string | ❌ | 结束日期，默认为查询当日。 |
| `market` | string | ❌ | (Default value = "cn") |

---

### `macro price-cpi`

**描述**: CPI居民消费价格指数

**API Method**: `econ.get_factors`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `factors` | string | array<string> | ✅ | 宏观因子名称，<https://assets.ricequant.com/vendor/rqdata/econ_get_factors.xlsx>（见文档） |
| `start_date` | string | ✅ | 起始日期 |
| `end_date` | string | ✅ | 截止日期 |
| `market` | string | ❌ | (Default value = "cn") |

---

### `macro price-ppi`

**描述**: PPI工业品出厂价格指数

**API Method**: `econ.get_factors`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `factors` | string | array<string> | ✅ | 宏观因子名称，<https://assets.ricequant.com/vendor/rqdata/econ_get_factors.xlsx>（见文档） |
| `start_date` | string | ✅ | 起始日期 |
| `end_date` | string | ✅ | 截止日期 |
| `market` | string | ❌ | (Default value = "cn") |

---

### `macro gdp`

**描述**: GDP国内生产总值

**API Method**: `econ.get_factors`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `factors` | string | array<string> | ✅ | 宏观因子名称，<https://assets.ricequant.com/vendor/rqdata/econ_get_factors.xlsx>（见文档） |
| `start_date` | string | ✅ | 起始日期 |
| `end_date` | string | ✅ | 截止日期 |
| `market` | string | ❌ | (Default value = "cn") |

---

### `macro pmi`

**描述**: PMI制造业采购经理指数

**API Method**: `econ.get_factors`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `factors` | string | array<string> | ✅ | 宏观因子名称，<https://assets.ricequant.com/vendor/rqdata/econ_get_factors.xlsx>（见文档） |
| `start_date` | string | ✅ | 起始日期 |
| `end_date` | string | ✅ | 截止日期 |
| `market` | string | ❌ | (Default value = "cn") |

---

### `macro interest-rate`

**描述**: Shibor利率

**API Method**: `econ.get_factors`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `factors` | string | array<string> | ✅ | 宏观因子名称，<https://assets.ricequant.com/vendor/rqdata/econ_get_factors.xlsx>（见文档） |
| `start_date` | string | ✅ | 起始日期 |
| `end_date` | string | ✅ | 截止日期 |
| `market` | string | ❌ | (Default value = "cn") |

---

### `macro get`

**描述**: 宏观数据查询

**API Method**: `econ.get_factors`

**Payload 参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| `factors` | string | array<string> | ✅ | 宏观因子名称，<https://assets.ricequant.com/vendor/rqdata/econ_get_factors.xlsx>（见文档） |
| `start_date` | string | ✅ | 起始日期 |
| `end_date` | string | ✅ | 截止日期 |
| `market` | string | ❌ | (Default value = "cn") |

---

