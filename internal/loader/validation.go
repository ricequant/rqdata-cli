package loader

import (
	"fmt"

	"github.com/ricequant/rqdata-cli/internal/validator"
)

func validateIDs(payload map[string]interface{}, validatorType string) error {
	if validatorType == "" || validatorType == "null" {
		return nil
	}

	ids, ok := payload["order_book_ids"]
	if !ok {
		return nil
	}

	var idList []string
	switch v := ids.(type) {
	case string:
		idList = []string{v}
	case []interface{}:
		for _, id := range v {
			if s, ok := id.(string); ok {
				idList = append(idList, s)
			}
		}
	}

	for _, id := range idList {
		var err error
		switch validatorType {
		case "cn_stock_code":
			err = validator.ValidateCNStock(id)
		case "hk_stock_code":
			err = validator.ValidateHKStock(id)
		case "index_code":
			err = validator.ValidateIndex(id)
		case "futures_code":
			err = validator.ValidateFutures(id)
		case "fund_code":
			err = validator.ValidateFund(id)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func validateDates(payload map[string]interface{}) error {
	dateFields := []string{"start_date", "end_date", "date"}
	for _, field := range dateFields {
		if val, ok := payload[field]; ok {
			if dateStr, ok := val.(string); ok {
				if err := validator.ValidateDate(dateStr); err != nil {
					return err
				}
				payload[field] = validator.DateToAPIFormat(dateStr)
			}
		}
	}

	// 检查日期范围
	if start, ok := payload["start_date"].(string); ok {
		if end, ok := payload["end_date"].(string); ok {
			if start > end {
				return fmt.Errorf("start_date must be <= end_date")
			}
		}
	}
	return nil
}
