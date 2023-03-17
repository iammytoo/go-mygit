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
	Use:   "hash_object",
	Short: "object作るよ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		is_save, _ := cmd.Flags().GetBool("write")
		if args[0] != "" {
			blob := plumbing.CreateHashObject(args[0])
			fmt.Println(blob.String())
			fmt.Println(blob.Hex)
			if is_save {
				plumbing.SaveHashObject(blob)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)
	hashObjectCmd.Flags().BoolP("write", "w", false, "送信先のメールアドレス")
}
