package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/zalando/go-keyring"
	"golang.org/x/term"
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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	passwordBytes, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	return username, string(passwordBytes), nil
}
