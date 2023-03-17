/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mirito333/go-mygit/plumbing"
	"github.com/spf13/cobra"
)

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hashObject",
	Short: "object作るよ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "" {
			blob := plumbing.CreateHashObject(args[0])
			fmt.Println(blob.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)
}
