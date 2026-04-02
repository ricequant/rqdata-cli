package loader

import (
	"bytes"
	"encoding/json"
)

func ListCommands() (string, error) {
	if err := Load(); err != nil {
		return "", err
	}

	var paths []string
	for _, group := range config.Groups {
		collectPaths(&group, "", &paths)
	}

	data, _ := json.MarshalIndent(paths, "", "  ")
	return string(data), nil
}

func collectPaths(group *Group, prefix string, paths *[]string) {
	groupPrefix := prefix
	if prefix == "" {
		groupPrefix = group.Name
	} else {
		groupPrefix = prefix + " " + group.Name
	}

	for _, cmd := range group.Commands {
		*paths = append(*paths, groupPrefix+" "+cmd.Name)
	}

	for _, subgroup := range group.Groups {
		collectPaths(&subgroup, groupPrefix, paths)
	}
}

func GetSchema(apiMethod string) (string, error) {
	if err := Load(); err != nil {
		return "", err
	}

	if schema == nil {
		return `{"error":"schema.json not found"}`, nil
	}

	if s, ok := schema[apiMethod]; ok {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		enc.Encode(s)
		return buf.String(), nil
	}

	return `{"error":"schema not found for ` + apiMethod + `"}`, nil
}
