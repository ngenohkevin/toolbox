package cmd

import (
	"fmt"
	"github.com/ngenohkevin/toolbox/cmd/info"
	"github.com/ngenohkevin/toolbox/cmd/net"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
//init package 

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "help message for toggle")

	addSubCommandPalettes()
}
