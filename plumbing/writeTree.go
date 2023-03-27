package plumbing

import (
	"os"
	"strconv"

	"github.com/mirito333/go-mygit/gitstructs"
	"github.com/mirito333/go-mygit/utils"
)

func CreateTree() gitstructs.Tree {
	comitter := gitstructs.CreateCommitHelper()
	comitter.Index = gitstructs.ReadIndex()
	comitter.CreateCurrentTree()
	comitter.ReadPreTree("./.go-mygit/refs/main")
	tree := comitter.CompareTree()
	return tree
}

func SaveTree(tree *gitstructs.Tree) error {
	inner := ""
	for _, c := range tree.Child {
		SaveTree(&c)
		inner += "tree " + c.Name + " " + c.Hash + "\n"
	}
	for _, b := range tree.Blob {
		if b.Status == "remove" {
			continue
		}
		inner += "blob " + b.Path + " " + b.Hex + "\n"
	}
	header := "tree " + strconv.Itoa(len(inner)) + "\\0"
	content := header + inner
	hash := utils.Sha1(content)
	tree.Hash = hash
	git_path := "./.go-mygit/objects/"
	object_path := git_path + hash
	by := utils.EncodeZlib(content)
	os.WriteFile(object_path, by.Bytes(), 0644)
	return nil
}
