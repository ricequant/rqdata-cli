package validator

import (
	"fmt"
	"regexp"
)

var (
	dateRegex     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	datetimeRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
)

func ValidateDate(date string) error {
	if !dateRegex.MatchString(date) && !datetimeRegex.MatchString(date) {
		return fmt.Errorf("invalid date format: %s", date)
	}
	return nil
}

func DateToAPIFormat(date string) string {
	if len(date) >= 10 {
		return date[:4] + date[5:7] + date[8:10]
	}
	return date
}
