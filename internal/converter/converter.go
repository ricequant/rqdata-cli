package converter

import (
	"encoding/csv"
	"strings"
)

func Convert(csvData, format string, fields []string) (string, error) {
	records, err := parseCSV(csvData)
	if err != nil {
		return "", err
	}

	if len(fields) > 0 {
		records = filterFields(records, fields)
	}

	switch format {
	case "csv":
		return toCSV(records), nil
	case "json":
		return toJSON(records)
	case "ndjson":
		return toNDJSON(records)
	default:
		return toNDJSON(records)
	}
}

// ConvertRecords converts records directly without parsing CSV
func ConvertRecords(records []map[string]string, format string, fields []string) (string, error) {
	if len(fields) > 0 {
		records = filterFields(records, fields)
	}

	switch format {
	case "csv":
		return toCSV(records), nil
	case "json":
		return toJSON(records)
	case "ndjson":
		return toNDJSON(records)
	default:
		return toNDJSON(records)
	}
}

func parseCSV(data string) ([]map[string]string, error) {
	r := csv.NewReader(strings.NewReader(data))
	rows, err := r.ReadAll()
	if err != nil || len(rows) == 0 {
		return nil, err
	}

	headers := rows[0]
	var records []map[string]string
	for _, row := range rows[1:] {
		record := make(map[string]string)
		for i, val := range row {
			if i < len(headers) {
				record[headers[i]] = val
			}
		}
		records = append(records, record)
	}
	return records, nil
}
