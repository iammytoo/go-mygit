package gitstructs

import (
	"strings"
)

type Tree struct {
	Name  string
	Hash  string
	Blob  []Blob
	Child []Tree
}

func CreateTree(name string, hash string) Tree {
	return Tree{Name: name, Hash: hash}
}

func WriteTree(indexes []Index, depth int, name string) Tree {
	tree := Tree{Name: name}
	for _, i := range indexes {
		path := i.Path
		splited_path := strings.Split(path, "/")
		if len(splited_path) == depth {
			if depth == 1 || name == splited_path[depth-2] {
				tree.Blob = append(tree.Blob, Blob{Hex: i.Hash, Path: splited_path[depth-1]})
			}
		}
		if len(splited_path) > depth {
			if !tree.containChild(splited_path[depth-1]) {
				tree.Child = append(tree.Child, WriteTree(indexes, depth+1, splited_path[depth-1]))
			}
		}
	}
	return tree
}

func (t *Tree) containChild(name string) bool {
	for _, t := range t.Child {
		if t.Name == name {
			return true
		}
	}
	return false
}

func (t *Tree) String() string {
	res := "tree : " + t.Name + "\n"
	for _, c := range t.Child {
		res += "child : " + c.Name + "\n"
	}
	for _, b := range t.Blob {
		res += "blob : " + b.Path + " " + b.Hex + "\n"
	}
	res += "---------\n"
	for _, c := range t.Child {
		res += c.String()
	}
	return res
}
