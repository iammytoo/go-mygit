/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "イニシャライズコマンド",
	Long:  `.go-mygit,.go-mygit/index,.go-mygit/refs,.go-mygit/refs/main,.go-mygit/objectsを作る`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.Mkdir("./.go-mygit", 0777); err != nil {
			fmt.Println(err)
		}
		if err := os.Mkdir("./.go-mygit/refs", 0777); err != nil {
			fmt.Println(err)
		}
		if err := os.Mkdir("./.go-mygit/objects", 0777); err != nil {
			fmt.Println(err)
		}
		if _,err := os.Create("./.go-mygit/index"); err != nil {
			fmt.Println(err)
		}
		if _,err := os.Create("./.go-mygit/refs/main"); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
