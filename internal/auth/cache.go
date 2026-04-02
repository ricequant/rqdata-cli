package auth

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func loadTokenCache() (*TokenCache, error) {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, tokenFile)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cache TokenCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, err
	}

	return &cache, nil
}

func saveTokenCache(cache TokenCache) error {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".rqdata")
	os.MkdirAll(dir, 0700)

	path := filepath.Join(home, tokenFile)
	data, _ := json.Marshal(cache)
	return os.WriteFile(path, data, 0600)
}

func isValid(cache *TokenCache) bool {
	return time.Now().Before(cache.ExpiresAt)
}

func clearTokenCache() error {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, tokenFile)
	return os.Remove(path)
}
