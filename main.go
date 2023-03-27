/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/mirito333/go-mygit/cmd"
	"github.com/mirito333/go-mygit/plumbing"
)

func main() {
	tree := plumbing.CreateTree()
	plumbing.SaveTree(&tree)
	cmd.Execute()
}
