/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mirito333/go-mygit/cmd"
	"github.com/mirito333/go-mygit/utils"
)

func main() {
	b := utils.EncodeZlib("hoge")
	os.WriteFile("./a", b.Bytes(), 0644)
	by, _ := os.ReadFile("./a")
	c := bytes.NewBuffer(by)
	s := utils.DecodeZlib(c)
	fmt.Println(s)
	cmd.Execute()
}
