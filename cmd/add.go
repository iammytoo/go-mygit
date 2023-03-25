/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mirito333/go-mygit/plumbing"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "後ろについてるpathのファイルのblobを作ってindexにのせる",
	Long:  `まだ-aを足してないよ`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "" {
			blob := plumbing.CreateHashObject(args[0])
			plumbing.SaveHashObject(blob)
			plumbing.UpdateIndex(args[0],blob,"initial")
		} else {
			fmt.Println("path指定してくれ")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
