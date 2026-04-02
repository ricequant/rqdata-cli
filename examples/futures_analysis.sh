#!/bin/bash
# 期货分析示例：股指期货

set -e

echo "=== 期货分析：股指期货 ==="
echo

echo "1. 沪深300期货主力合约"
rqdata futures dominant --payload '{
  "underlying_symbol": "IF",
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}' --format json
echo

echo "2. 主力连续合约行情"
rqdata futures dominant-price --payload '{
  "underlying_symbols": "IF",
  "start_date": "2024-01-01",
  "end_date": "2024-12-31",
  "fields": ["open", "close", "volume", "open_interest"]
}' --format json
echo

echo "=== 分析完成 ==="
