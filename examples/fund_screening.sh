#!/bin/bash
# 基金筛选示例：权益类公募基金

set -e

echo "=== 基金筛选：权益类公募基金 ==="
echo

echo "1. 基金净值（近 1 年）"
rqdata fund nav --payload '{
  "order_book_ids": ["000001"],
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}' --format json
echo

echo "2. 基金持仓明细"
rqdata fund holdings --payload '{
  "order_book_ids": ["000001"],
  "date": "2024-06-30"
}' --format json
echo

echo "3. 基金衍生指标（夏普/最大回撤）"
rqdata fund indicators --payload '{
  "order_book_ids": ["000001"],
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}' --format json
echo

echo "4. 基金经理信息"
rqdata fund manager --payload '{
  "order_book_ids": ["000001"]
}' --format json
echo

echo "=== 筛选完成 ==="
