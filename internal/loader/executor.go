package loader

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ricequant/rqdata-cli/internal/auth"
	"github.com/ricequant/rqdata-cli/internal/client"
	"github.com/ricequant/rqdata-cli/internal/converter"
)

func ExecuteCommand(cmd Command, payloadStr, format string, fields []string) error {
	// 1. 解析 payload
	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(payloadStr), &payload); err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}

	// 2. 验证 order_book_ids
	if err := validateIDs(payload, cmd.Validator); err != nil {
		return err
	}

	// 3. 验证并转换日期
	if err := validateDates(payload); err != nil {
		return err
	}

	// 4. 注入默认值
	for k, v := range cmd.InjectDefaults {
		if _, exists := payload[k]; !exists {
			payload[k] = v
		}
	}

	// 5. 构建请求参数
	params := make(map[string]interface{})
	for k, v := range payload {
		if k == "fields" && !cmd.APISupportsFields {
			continue
		}
		params[k] = v
	}

	// 6. 调用 API（token 过期时自动重新认证重试）
	token, err := auth.GetToken()
	if err != nil {
		return err
	}
	apiMethod := cmd.APIMethod

	csvData, err := client.CallAPI(token, apiMethod, params)
	if errors.Is(err, client.ErrUnauthorized) {
		token, err = auth.GetFreshToken()
		if err != nil {
			return err
		}
		csvData, err = client.CallAPI(token, apiMethod, params)
	}
	if err != nil {
		return err
	}

	// 8. 转换输出
	clientFields := []string{}
	if !cmd.APISupportsFields {
		clientFields = fields
	}

	output, err := converter.Convert(csvData, format, clientFields)
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}
