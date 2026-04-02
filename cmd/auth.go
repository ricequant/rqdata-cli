package cmd

import (
	"fmt"

	"github.com/ricequant/rqdata-cli/internal/auth"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "认证管理",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "交互式登录",
	RunE: func(cmd *cobra.Command, args []string) error {
		return auth.Login()
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "登出并清除凭证",
	RunE: func(cmd *cobra.Command, args []string) error {
		return auth.Logout()
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "查看认证状态",
	RunE: func(cmd *cobra.Command, args []string) error {
		status, err := auth.Status()
		if err != nil {
			return err
		}
		fmt.Println(status)
		return nil
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)
	authCmd.AddCommand(statusCmd)
}
