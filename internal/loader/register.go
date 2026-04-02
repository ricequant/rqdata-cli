package loader

import (
	"github.com/spf13/cobra"
)

func RegisterAll(root *cobra.Command) error {
	if err := Load(); err != nil {
		return err
	}

	for _, group := range config.Groups {
		registerGroup(root, &group)
	}
	return nil
}

func registerGroup(parent *cobra.Command, group *Group) {
	groupCmd := &cobra.Command{
		Use:   group.Name,
		Short: group.Description,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// 注册子命令
	for _, cmd := range group.Commands {
		registerCommand(groupCmd, cmd)
	}

	// 特殊处理：为 macro 组添加 indicators 命令
	if group.Name == "macro" {
		registerIndicatorsCommand(groupCmd)
	}

	// 递归注册子分组
	for _, subgroup := range group.Groups {
		registerGroup(groupCmd, &subgroup)
	}

	parent.AddCommand(groupCmd)
}
