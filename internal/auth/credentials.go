package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/zalando/go-keyring"
)

type credentialsFile struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loadCredentials() (string, string, error) {
	// 1. RQDATAC_CONF / RQDATAC2_CONF 环境变量（tcp://user:pass@host:port）
	for _, envKey := range []string{"RQDATAC_CONF", "RQDATAC2_CONF"} {
		if u, p, err := parseRqdatacConf(os.Getenv(envKey)); err == nil {
			return u, p, nil
		}
	}

	// 2. RQDATA_USERNAME / RQDATA_PASSWORD 环境变量
	if u := os.Getenv("RQDATA_USERNAME"); u != "" {
		if p := os.Getenv("RQDATA_PASSWORD"); p != "" {
			return u, p, nil
		}
	}

	// 3. Keyring（优先）
	if u, p, err := loadFromKeyring(); err == nil {
		return u, p, nil
	}

	// 4. 凭证文件（降级）
	if u, p, err := loadFromFile(); err == nil {
		return u, p, nil
	}

	// 5. 提示输入
	return promptCredentials()
}

func parseRqdatacConf(conf string) (string, string, error) {
	if conf == "" {
		return "", "", fmt.Errorf("empty")
	}
	u, err := url.Parse(conf)
	if err != nil || u.User == nil {
		return "", "", fmt.Errorf("invalid")
	}
	username := u.User.Username()
	password, _ := u.User.Password()
	if username == "" || password == "" {
		return "", "", fmt.Errorf("missing credentials")
	}
	return username, password, nil
}

func loadFromKeyring() (string, string, error) {
	username, err := keyring.Get(serviceName, "username")
	if err != nil {
		return "", "", err
	}
	password, err := keyring.Get(serviceName, username)
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}

func saveToKeyring(username, password string) error {
	if err := keyring.Set(serviceName, "username", username); err != nil {
		return err
	}
	return keyring.Set(serviceName, username, password)
}

func deleteFromKeyring() {
	username, err := keyring.Get(serviceName, "username")
	if err == nil {
		keyring.Delete(serviceName, username)
	}
	keyring.Delete(serviceName, "username")
}

func loadFromFile() (string, string, error) {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".rqdata", "credentials")

	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}

	var creds credentialsFile
	if err := json.Unmarshal(data, &creds); err != nil {
		return "", "", err
	}

	return creds.Username, creds.Password, nil
}

func saveToFile(username, password string) error {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".rqdata")
	os.MkdirAll(dir, 0700)

	path := filepath.Join(dir, "credentials")
	creds := credentialsFile{Username: username, Password: password}
	data, _ := json.Marshal(creds)
	return os.WriteFile(path, data, 0600)
}

func deleteFromFile() error {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".rqdata", "credentials")
	return os.Remove(path)
}

func promptCredentials() (string, string, error) {
	input := os.Stdin
	shouldClose := false

	// When stdin is redirected, read from the controlling terminal instead so
	// interactive login still blocks for user input.
	if stat, err := os.Stdin.Stat(); err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		tty, err := os.Open("/dev/tty")
		if err != nil {
				return "", "", fmt.Errorf("无法读取交互式输入: stdin 不是终端，且打开 /dev/tty 失败: %w", err)
		}
		input = tty
		shouldClose = true
	}
	if shouldClose {
		defer input.Close()
	}

	reader := bufio.NewReader(input)

	fmt.Fprint(os.Stderr, "license: ")
	license, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("读取 license 失败: %w", err)
	}
	license = strings.TrimSpace(license)
	fmt.Fprintln(os.Stderr)

	return "license", license, nil
}
