package gitstructs

type Tree struct {
	Name  string
	Hash  string
	Blob  []Blob
	Child []Tree
}

func CreateTree(name string, hash string) Tree {
	return Tree{Name: name, Hash: hash}
}

