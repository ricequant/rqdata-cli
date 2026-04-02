# Changelog

## [1.0.0] - 2026-03-27

### 首次发布

**核心特性**
- Go 实现的高性能 CLI 工具
- 支持 68 个金融数据 API 命令
- NDJSON/JSON/CSV 多格式输出
- 系统 Keyring 集成，安全存储凭证
- 自动 Token 管理和刷新
- Schema 自动生成，含中文参数描述

**数据覆盖**
- 指数数据（8 个命令）
- A股数据（23 个命令）
- 港股数据（8 个命令）
- 交易日历（5 个命令）
- 基金数据（9 个命令）
- 期货数据（8 个命令）
- 期权数据（7 个命令）
- 宏观数据（6 个命令）

**技术亮点**
- 单文件可执行，静态编译
- 跨平台支持（Linux/macOS/Windows）
- 类型安全的参数验证
- AI Agent 友好的 NDJSON 输出

**文档**
- 完整的快速上手指南
- 68 个命令的详细参考文档
- Go 构建指南


### Added

#### 核心功能
- ✨ 完整的 RQData HTTP API 支持
- 🚀 CURL 连接复用优化（QPS 提升 44.4%）
- ⚡ Daemon 模式支持（长连接服务）
- 📊 多种输出格式（NDJSON、JSON、CSV）
- 🔒 严格的 Schema 验证
- 🎯 字段过滤功能（`--fields` 参数）

#### 数据模块
- 📈 **Index（指数）** - 指数列表、行情、成分股、权重
- 📊 **Stock CN（A股）** - 行情、财务、分红、融资融券、北向资金、停牌
- 🇭🇰 **Stock HK（港股）** - 行情、财务、分红
- 📅 **Calendar（交易日历）** - 交易日列表、前后交易日查询
- 📉 **Macro（宏观）** - Shibor、CPI、PPI、GDP

#### 认证系统
- 🔑 环境变量认证（推荐）
- 🔐 交互式登录
- 💾 Token 缓存机制
- ⏰ 自动 Token 刷新

#### 性能优化
- 🔄 CURL 连接复用
- 🌐 TCP Keep-Alive
- 📡 HTTP 持久连接
- ⚡ Daemon 模式（批量查询性能提升 25%）

#### 测试覆盖
- ✅ 74 个功能测试（100% 通过率）
- 📊 5 个性能测试场景
- 🔬 三方性能对比（CLI vs Daemon vs Python）
- 🧪 全模块性能验证

### Performance

#### 性能基准
- **单次请求**: 0.140s（优化前 0.154s，提升 9.1%）
- **连续请求**: 0.150s/次（优化前 0.200s/次，提升 25%）
- **平均 QPS**: 7.22 req/s（优化前 5 req/s，提升 44.4%）
- **大数据吞吐**: 7,652 行/秒

#### 场景性能
- **1-3次查询**: CLI 直接模式最快（比 Python 快 68-78%）
- **3-10次查询**: CLI Daemon 模式最快（比 Python 快 60-78%）
- **20+次高频**: Python API 最快（比 CLI 快 12倍）
- **大数据量**: Python API 最快（比 CLI 快 2.5倍）

### Documentation

- 📖 完整的 README.md
- 📚 API 文档
- 🔧 开发指南
- 📊 性能测试报告
- 🤝 贡献指南
- 📝 更新日志

### Technical Details

#### 架构
- C++17 标准
- CMake 构建系统
- vcpkg 依赖管理
- 模块化设计

#### 依赖库
- libcurl - HTTP 客户端
- nlohmann-json - JSON 处理
- CLI11 - 命令行解析
- csv-parser - CSV 解析
- fmt - 格式化输出

#### 平台支持
- ✅ macOS (Apple Silicon & Intel)
- ✅ Linux (x86_64)
- 🚧 Windows (计划中)

### Testing

#### 功能测试
- Index 模块: 8/8 通过 ✅
- Stock CN 模块: 23/23 通过 ✅
- Stock HK 模块: 8/8 通过 ✅
- Calendar 模块: 5/5 通过 ✅
- Fund 模块: 9/9 通过 ✅
- Futures 模块: 8/8 通过 ✅
- Options 模块: 7/7 通过 ✅
- Macro 模块: 6/6 通过 ✅
- **总计**: 74/74 通过（100%）✅

#### 性能测试
- 单次请求性能测试 ✅
- 连续请求性能测试 ✅
- 大数据量性能测试 ✅
- 高频请求性能测试 ✅
- 三方对比测试 ✅

### Known Issues

- Stock HK 模块性能相对较慢（QPS 3.17）- 受 API 响应速度限制
- 高频场景（20+次/秒）性能不如 Python API - 受进程启动开销限制

### Recommendations

#### 使用建议
- **临时查询（1-3次）**: 使用 CLI 直接模式
- **批量脚本（3-20次）**: 使用 CLI Daemon 模式
- **高频查询（20+次）**: 使用 Python API
- **大数据量（10K+行）**: 使用 Python API

## [1.0.1] - 2026-03-10

### Fixed

#### 财务数据功能修复
- 🐛 **修复 `stock cn financial` 命令 HTTP 400 错误**
  - 问题：`get_pit_financials_ex` API 要求 `fields` 必填，但 CLI 在构建请求时把 `fields` 从 payload 中剔除
  - 修复：针对 `get_pit_financials_ex` 方法，保留 payload 中的 `fields` 参数传给 API
  - 影响：现在可以正常获取三大表财务数据（资产负债表、利润表、现金流量表）

- 🐛 **修复 `stock cn financial-indicator` 命令 HTTP 400 错误**
  - 问题：使用了错误的 method `get_factor`（参数格式不匹配）
  - 修复：改用 `get_pit_financials_ex` + 衍生指标字段
  - 影响：现在可以正常获取 ROE、毛利率、净利率等财务衍生指标

#### 认证问题修复
- 🐛 **修复 Token 过期导致的 HTTP 401 错误**
  - 问题：缓存的 token 过期后，CLI 继续使用导致认证失败
  - 解决方案：删除 `~/.rqdata/token.cache` 后，CLI 自动重新认证获取新 token
  - 建议：如遇 401 错误，运行 `rm -f ~/.rqdata/token.cache` 清除过期缓存

### Added

#### 完整的 HALO 策略支持
- ✅ 现在可以用纯 CLI 完成 HALO 策略筛选
- ✅ 支持获取三大表财务数据
- ✅ 支持获取财务衍生指标（ROE、毛利率等）
- ✅ 支持批量查询 20 只股票的财务数据

### Documentation

- 📝 添加财务数据 API 使用说明
- 📝 添加衍生指标字段名参考
- 📝 添加 Token 过期问题排查指南

### Technical Details

#### 代码修改
- `src/commands/stock_cn.cpp` 第 71-84 行：修改 `execute_stock_cn_command` 函数，针对 `get_pit_financials_ex` 保留 `fields` 参数
- `src/commands/stock_cn.cpp` 第 426 行：`financial-indicator` 命令改用 `get_pit_financials_ex` 方法

#### 测试验证
- ✅ `stock cn financial` 命令测试通过
- ✅ `stock cn financial-indicator` 命令测试通过
- ✅ HALO 策略 20 只股票批量查询测试通过

### Notes

#### 衍生指标字段名
- 使用完整名称：`return_on_equity_weighted_average`（不是 `roe`）
- TTM 字段使用大写：`net_profitTTM`（不是 `net_profit_ttm`）
- Quarter 格式使用小写：`2024q3`（不是 `2024Q3`）

## [Unreleased]

### Planned Features

- 🚧 Windows 平台支持
- 🚧 批量 API 支持
- 🚧 并行请求支持
- 🚧 MCP Server 支持
- 🚧 更多输出格式（Parquet、Arrow）

### Future Optimizations

- 🔮 CSV 解析优化
- 🔮 二进制数据格式支持
- 🔮 内存池优化
- 🔮 并发连接池

---

## Version History

- **v1.0.1** (2026-03-10) - 财务数据修复
- **v1.0.0** (2026-03-10) - 初始版本发布

---

## Links

- [Repository](http://git.ricequant.com/projects/RQAI/repos/rqdata-cli)
- [Documentation](http://git.ricequant.com/projects/RQAI/repos/rqdata-cli/browse/docs)
- [Issues](https://jira.ricequant.com)（RiceQuant 内部 Jira）
