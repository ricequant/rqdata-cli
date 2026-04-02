package cmd

import (
	"fmt"

	"github.com/ricequant/rqdata-cli/internal/loader"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Schema 自省",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		commands, err := loader.ListCommands()
		if err != nil {
			return err
		}
		fmt.Println(commands)
		return nil
	},
}

func init() {
	schemaCmd.AddCommand(listCmd)
}
