package loader

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/ricequant/rqdata-cli/internal/configs"
)

func Load() error {
	// 加载顺序：当前目录 -> ~/.rqdata/ -> 嵌入配置
	configData, err := loadConfigFile()
	if err != nil {
		configData = configs.EmbeddedCommands
	}

	if err := json.Unmarshal(configData, &config); err != nil {
		return err
	}

	schemaData, err := loadSchemaFile()
	if err != nil {
		schemaData = configs.EmbeddedSchema
	}
	json.Unmarshal(schemaData, &schema)

	return nil
}

func loadConfigFile() ([]byte, error) {
	if data, err := os.ReadFile("commands.json"); err == nil {
		return data, nil
	}
	home, _ := os.UserHomeDir()
	return os.ReadFile(filepath.Join(home, ".rqdata", "commands.json"))
}

func loadSchemaFile() ([]byte, error) {
	if data, err := os.ReadFile("schema.json"); err == nil {
		return data, nil
	}
	home, _ := os.UserHomeDir()
	return os.ReadFile(filepath.Join(home, ".rqdata", "schema.json"))
}
