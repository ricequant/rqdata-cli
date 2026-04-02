#!/bin/bash
# 行业分析示例：银行板块

set -e

echo "=== 行业分析：银行板块 ==="
echo

echo "1. 获取中证银行指数成分股"
rqdata index constituents --payload '{
  "order_book_id": "000300.XSHG",
  "date": "2024-01-02"
}' --format json
echo

echo "2. 银行股估值对比"
rqdata stock cn financial-indicator --payload '{
  "order_book_ids": ["000001.XSHE", "600036.XSHG", "601398.XSHG"],
  "factor": ["pe_ratio_ttm", "pb_ratio", "roe_ttm", "market_cap"],
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}' --format json
echo

echo "3. 行业指数走势对比"
rqdata index price --payload '{
  "order_book_ids": ["000300.XSHG", "000905.XSHG"],
  "start_date": "2024-01-01",
  "end_date": "2024-12-31",
  "fields": ["close"]
}' --format json
echo

echo "=== 分析完成 ==="
