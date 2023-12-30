package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

// doesn't work yet
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "prints disk usage of the current directory",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		usage := du.NewDiskUsage(".")
		fmt.Printf("%v\n", usage)
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)
}
