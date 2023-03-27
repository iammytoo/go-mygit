package plumbing

import (
	"fmt"

	"github.com/mirito333/go-mygit/gitstructs"
)

func CreateTree() gitstructs.Tree {
	comitter := gitstructs.CreateCommitHelper()
	comitter.Index = gitstructs.ReadIndex()
	comitter.CreateCurrentTree()
	fmt.Println(comitter.CurrentTree.String())
	comitter.ReadPreTree()
	tree := comitter.CompareTree()
	return tree
}

func SaveTree(tree gitstructs.Tree) error{
	return nil
}
