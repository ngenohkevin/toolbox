package info

import (
	"fmt"
	"github.com/spf13/cobra"
)

var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "diskusage information",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diskusage called")
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)
}
