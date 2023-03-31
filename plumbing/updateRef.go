package plumbing

import "os"

func UpdateRef(hash string) error{
	refPath := "./.go-mygit/refs/main"
	os.WriteFile("./.go-mygit/index",[]byte(""),0644)
	index, _ := os.Create(refPath)
	index.Write([]byte(hash))
	return nil
}
