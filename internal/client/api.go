package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrUnauthorized = errors.New("401 unauthorized")

func (c *Client) CallAPI(token, method string, params map[string]interface{}) (string, error) {
	payload := map[string]interface{}{
		"method": method,
	}
	for k, v := range params {
		payload[k] = v
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", dataURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == 401 {
		return "", ErrUnauthorized
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API 错误: HTTP %d, %s", resp.StatusCode, string(data))
	}

	return string(data), nil
}

func CallAPI(token, method string, params map[string]interface{}) (string, error) {
	return New().CallAPI(token, method, params)
}
