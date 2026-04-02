package cmd

import (
	"github.com/ricequant/rqdata-cli/internal/loader"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "rqdata",
	Short:   "RQData CLI - 面向 AI Agent 的金融数据工具",
	Version: Version,
}

func Execute() error {
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(schemaCmd)
	loader.RegisterAll(rootCmd)
	return rootCmd.Execute()
}
