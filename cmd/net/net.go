package net

import (
	"github.com/spf13/cobra"
)

var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "net is a palette that contain network based commands",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func init() {

}
