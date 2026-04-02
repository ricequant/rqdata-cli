package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	authURL = "https://rqdata.ricequant.com/auth"
	dataURL = "https://rqdata.ricequant.com/api"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func Authenticate(username, password string) (string, error) {
	payload := map[string]string{
		"user_name": username,
		"password":  password,
	}
	body, _ := json.Marshal(payload)

	resp, err := http.Post(authURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("认证失败: HTTP %d", resp.StatusCode)
	}

	token, _ := io.ReadAll(resp.Body)
	return string(token), nil
}
