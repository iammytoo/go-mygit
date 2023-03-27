package gitstructs

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
	root_tree := WriteTree(c.Index, 1,"/")
	c.CurrentTree = root_tree
}

func(c *CommitHelper) ReadPreTree() {

}

func (c *CommitHelper) CompareTree() Tree {
	return Tree{}
}