package validator

import (
	"fmt"
	"regexp"
)

var (
	cnStockRegex  = regexp.MustCompile(`^\d{6}\.(SH|SZ|XSHG|XSHE)$`)
	hkStockRegex  = regexp.MustCompile(`^\d{5}\.(HK|XHKG)$`)
	futuresRegex  = regexp.MustCompile(`^[A-Z]{1,2}\d{4}$`)
	fundRegex     = regexp.MustCompile(`^\d{6}$`)
)

func ValidateCNStock(code string) error {
	if !cnStockRegex.MatchString(code) {
		return fmt.Errorf("invalid CN stock code: %s", code)
	}
	return nil
}

func ValidateHKStock(code string) error {
	if !hkStockRegex.MatchString(code) {
		return fmt.Errorf("invalid HK stock code: %s", code)
	}
	return nil
}

func ValidateIndex(code string) error {
	return nil
}

func ValidateFutures(code string) error {
	if !futuresRegex.MatchString(code) {
		return fmt.Errorf("invalid futures code: %s", code)
	}
	return nil
}

func ValidateFund(code string) error {
	if !fundRegex.MatchString(code) {
		return fmt.Errorf("invalid fund code: %s", code)
	}
	return nil
}
