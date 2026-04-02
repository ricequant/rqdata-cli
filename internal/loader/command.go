package loader

import (
	"fmt"

	"github.com/spf13/cobra"
)

func registerCommand(parent *cobra.Command, cmd Command) {
	var payload string
	var format string
	var fields []string
	var showSchema bool

	cmdObj := &cobra.Command{
		Use:   cmd.Name,
		Short: cmd.Description,
		RunE: func(c *cobra.Command, args []string) error {
			if showSchema {
				schema, err := GetSchema(cmd.APIMethod)
				if err != nil {
					return err
				}
				fmt.Println(schema)
				return nil
			}

			if payload == "" {
				return fmt.Errorf(`{"error":{"message":"--payload is required"}}`)
			}

			return ExecuteCommand(cmd, payload, format, fields)
		},
	}

	cmdObj.Flags().StringVar(&payload, "payload", "", "JSON payload")
	cmdObj.Flags().StringVar(&format, "format", "ndjson", "输出格式：ndjson/json/csv")
	cmdObj.Flags().StringSliceVar(&fields, "fields", nil, "限制返回字段")
	cmdObj.Flags().BoolVar(&showSchema, "schema", false, "显示命令 Schema")

	parent.AddCommand(cmdObj)
}
