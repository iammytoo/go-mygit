package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateIndexCmd = &cobra.Command{
	Use:   "update-index",
	Short: "indexを書き換えるコマンド",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ごめん作ってない")
	},
}

func init() {
	rootCmd.AddCommand(updateIndexCmd)
}
