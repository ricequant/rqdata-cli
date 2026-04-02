package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ricequant/rqdata-cli/internal/client"
)

func Login() error {
	// 尝试从环境变量获取凭证
	username, password, err := loadCredentials()
	if err != nil {
		return err
	}

	// 验证凭证
	token, err := client.Authenticate(username, password)
	if err != nil {
		return fmt.Errorf("认证失败: %w", err)
	}

	// 保存凭证（优先 keyring，失败时用文件）
	if err := saveToKeyring(username, password); err != nil {
		if err := saveToFile(username, password); err != nil {
			return fmt.Errorf("保存凭证失败: %w", err)
		}
	}

	// 缓存 token
	saveTokenCache(TokenCache{
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	})

	fmt.Println("登录成功")
	return nil
}

func Logout() error {
	deleteFromKeyring()
	deleteFromFile()
	clearTokenCache()
	fmt.Println("已登出")
	return nil
}

func Status() (string, error) {
	status := make(map[string]interface{})

	// 检查凭证来源
	for _, envKey := range []string{"RQDATAC_CONF", "RQDATAC2_CONF"} {
		if conf := os.Getenv(envKey); conf != "" {
			if username, _, err := parseRqdatacConf(conf); err == nil {
				status["credentials"] = envKey
				status["username"] = username
				goto checkToken
			}
		}
	}

	if os.Getenv("RQDATA_USERNAME") != "" {
		status["credentials"] = "RQDATA_USERNAME"
		status["username"] = os.Getenv("RQDATA_USERNAME")
	} else if username, _, err := loadFromKeyring(); err == nil {
		status["credentials"] = "keyring"
		status["username"] = username
	} else if username, _, err := loadFromFile(); err == nil {
		status["credentials"] = "file"
		status["username"] = username
	} else {
		status["credentials"] = "none"
	}

checkToken:
	// 检查 token 缓存
	if cache, err := loadTokenCache(); err == nil {
		status["token_cached"] = true
		status["token_valid"] = isValid(cache)
		status["expires_at"] = cache.ExpiresAt.Format(time.RFC3339)
	} else {
		status["token_cached"] = false
	}

	data, _ := json.MarshalIndent(status, "", "  ")
	return string(data), nil
}
