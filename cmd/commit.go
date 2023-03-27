/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mirito333/go-mygit/plumbing"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "indexの内容をcommitするよ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("message")
		tree := plumbing.CreateTree()
		plumbing.SaveTree(&tree)
		plumbing.CommitTree(tree.Hash, message)
		plumbing.UpdateRef(tree.Hash)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringP("message", "m", "commit!", "コミットメッセージ")
}
