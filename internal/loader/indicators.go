package loader

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ricequant/rqdata-cli/internal/converter"
	"github.com/spf13/cobra"
)

//go:embed econ_factors.txt
var econFactors string

func registerIndicatorsCommand(parent *cobra.Command) {
	var format string
	var fields []string

	cmdObj := &cobra.Command{
		Use:   "indicators",
		Short: "宏观指标列表",
		RunE: func(c *cobra.Command, args []string) error {
			lines := strings.Split(strings.TrimSpace(econFactors), "\n")
			var records []map[string]string
			for _, line := range lines {
				if line != "" {
					records = append(records, map[string]string{"indicator": line})
				}
			}
			output, err := converter.ConvertRecords(records, format, fields)
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		},
	}

	cmdObj.Flags().StringVar(&format, "format", "ndjson", "输出格式：ndjson/json/csv")
	cmdObj.Flags().StringSliceVar(&fields, "fields", nil, "限制返回字段")

	parent.AddCommand(cmdObj)
}
