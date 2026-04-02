#!/bin/bash
# 个股研究示例：平安银行（000001.XSHE）

set -e

STOCK="000001.XSHE"

echo "=== 个股研究：平安银行 ==="
echo

echo "1. 行情概览（近 1 年）"
rqdata stock cn price --payload "{
  \"order_book_ids\": [\"$STOCK\"],
  \"start_date\": \"2024-01-01\",
  \"end_date\": \"2024-12-31\",
  \"fields\": [\"close\", \"volume\", \"total_turnover\"],
  \"adjust_type\": \"pre\"
}" --format json
echo

echo "2. 财务分析（近 5 年三大表）"
rqdata stock cn financial --payload "{
  \"order_book_ids\": [\"$STOCK\"],
  \"fields\": [\"revenue\", \"net_profit\", \"total_assets\", \"total_liabilities\"],
  \"start_quarter\": \"2020Q1\",
  \"end_quarter\": \"2024Q3\"
}" --format json
echo

echo "3. 估值定位（PE/PB/ROE TTM）"
rqdata stock cn financial-indicator --payload "{
  \"order_book_ids\": [\"$STOCK\"],
  \"factor\": [\"pe_ratio_ttm\", \"pb_ratio\", \"roe_ttm\"],
  \"start_date\": \"2024-01-01\",
  \"end_date\": \"2024-12-31\"
}" --format json
echo

echo "4. 股东结构（最新季度）"
rqdata stock cn shareholder-top10 --payload "{
  \"order_book_ids\": [\"$STOCK\"],
  \"start_date\": \"2024-09-30\",
  \"end_date\": \"2024-09-30\"
}" --format json
echo

echo "5. 一致预期（分析师预测）"
rqdata stock cn consensus --payload "{
  \"order_book_ids\": [\"$STOCK\"],
  \"factor\": [\"eps_consensus\", \"target_price\", \"analyst_count\"],
  \"start_date\": \"2024-01-01\",
  \"end_date\": \"2024-12-31\"
}" --format json
echo

echo "=== 研究完成 ==="
