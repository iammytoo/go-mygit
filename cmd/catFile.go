/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mirito333/go-mygit/utils"
	"github.com/spf13/cobra"
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "ハッシュを指定して、情報を見れるよ",
	Long:  `ほげ`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "" {
			object_path := "./.go-mygit/objects/" + args[0]
			by, _ := os.ReadFile(object_path)
			c := bytes.NewBuffer(by)
			s := utils.DecodeZlib(c)
			header_content := strings.Split(s,"\\0")
			header := header_content[0]
			contentType := strings.Split(header," ")[0]
			contentSize := strings.Split(header," ")[1]
			content := header_content[1]
			fmt.Println("type: "+contentType)
			fmt.Println("size: "+contentSize)
			fmt.Println(content)
		}
	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)
}
