package converter

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"strings"
)

func filterFields(records []map[string]string, fields []string) []map[string]string {
	var filtered []map[string]string
	for _, record := range records {
		newRecord := make(map[string]string)
		for _, field := range fields {
			if val, ok := record[field]; ok {
				newRecord[field] = val
			}
		}
		filtered = append(filtered, newRecord)
	}
	return filtered
}

func toCSV(records []map[string]string) string {
	if len(records) == 0 {
		return ""
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)

	// 写入表头
	var headers []string
	for k := range records[0] {
		headers = append(headers, k)
	}
	w.Write(headers)

	// 写入数据
	for _, record := range records {
		var row []string
		for _, h := range headers {
			row = append(row, record[h])
		}
		w.Write(row)
	}
	w.Flush()
	return buf.String()
}

func toJSON(records []map[string]string) (string, error) {
	data, err := json.MarshalIndent(records, "", "  ")
	return string(data), err
}

func toNDJSON(records []map[string]string) (string, error) {
	var lines []string
	for _, record := range records {
		data, _ := json.Marshal(record)
		lines = append(lines, string(data))
	}
	return strings.Join(lines, "\n"), nil
}
