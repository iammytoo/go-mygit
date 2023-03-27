package gitstructs

import "os"

type CommitHelper struct {
	Hash        string
	Index       []Index
	CurrentTree Tree
	ParentTree  Tree
}

func CreateCommitHelper() CommitHelper {
	return CommitHelper{}
}

func (c *CommitHelper) CreateCurrentTree() {
	root_tree := WriteTree(c.Index, 1, "/")
	c.CurrentTree = root_tree
}

func (c *CommitHelper) ReadPreTree(refPath string) {
	refHash, _ := os.ReadFile(refPath)
	pre_tree := ReadTree(string(refHash), "/")
	c.ParentTree = pre_tree
}

func (c *CommitHelper) CompareTree() Tree {
	tree := c.ParentTree
	c.CurrentTree.Compare(&tree)
	return tree
}
