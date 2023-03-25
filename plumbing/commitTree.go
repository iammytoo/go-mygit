package plumbing

import (
	"os"
	"strconv"

	"github.com/mirito333/go-mygit/utils"
)

func CommitTree(treeHash string, message string) string {
	refPath := "./.go-mygit/refs/main"
	refHash, _ := os.ReadFile(refPath)
	content := "tree " + treeHash + "\nbefore " + string(refHash)
	commit := "commit " + strconv.Itoa(len(content)) + "\\0"+content
	hash := utils.Sha1(commit)
	git_path := ".go-mygit/objects/"
	object_path := git_path + hash
	by := utils.EncodeZlib(content)
	os.WriteFile(object_path, by.Bytes(), 0644)
	return hash
}
