package auth

import (
	"time"

	"github.com/ricequant/rqdata-cli/internal/client"
)

const (
	serviceName = "rqdata-cli"
	tokenFile   = ".rqdata/token.cache"
)

type TokenCache struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func GetToken() (string, error) {
	// 1. 检查缓存
	if cache, err := loadTokenCache(); err == nil && isValid(cache) {
		return cache.Token, nil
	}
	return GetFreshToken()
}

// GetFreshToken 跳过缓存，重新认证并更新缓存。
func GetFreshToken() (string, error) {
	clearTokenCache()

	username, password, err := loadCredentials()
	if err != nil {
		return "", err
	}

	token, err := client.Authenticate(username, password)
	if err != nil {
		return "", err
	}

	saveTokenCache(TokenCache{
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	})

	return token, nil
}
