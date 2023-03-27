package gitstructs

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mirito333/go-mygit/utils"
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
				tree.Blob = append(tree.Blob, Blob{Hex: i.Hash, Path: path, Status: i.Status})
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

func ReadTree(hex, name string) Tree {
	treePath := "./.go-mygit/objects/" + hex
	res := Tree{Name: name, Hash: hex}
	by, err := os.ReadFile(treePath)
	if err != nil {
		fmt.Println("多分refがない")
		return Tree{}
	}
	c := bytes.NewBuffer(by)
	s := utils.DecodeZlib(c)
	header_content := strings.Split(s, "\\0")
	content := header_content[1]
	for _, line := range strings.Split(content, "\n") {
		splited_line := strings.Split(line, " ")
		if splited_line[0] == "tree" {
			res.Child = append(res.Child, ReadTree(splited_line[2], splited_line[1]))
		} else if splited_line[0] == "blob" {
			res.Blob = append(res.Blob, Blob{Path: splited_line[1], Hex: splited_line[2]})
		}
	}
	return res
}

func (t *Tree) Compare(parent *Tree) {
	for _, b := range t.Blob {
		switch b.Status {
		case "change":
			f, pb := parent.findBlob(b.Path)
			if f {
				pb.Hex = b.Hex
			}
		case "create":
			parent.Blob = append(parent.Blob, b)
		case "remove":
			f, pb := parent.findBlob(b.Path)
			if f {
				pb.Status = "remove"
			}
		}
	}
	for _, c := range t.Child {
		f, pc := parent.findChild(c.Name)
		if f {
			c.Compare(pc)
		} else {
			parent.Child = append(parent.Child, c)
			f2, pc2 := parent.findChild(c.Name)
			if f2 {
				c.Compare(pc2)
			}
		}
	}
}

func (t *Tree) containChild(name string) bool {
	for _, ta := range t.Child {
		if ta.Name == name {
			return true
		}
	}
	return false
}

func (t *Tree) findChild(name string) (bool, *Tree) {
	for _, ta := range t.Child {
		if ta.Name == name {
			return true, &ta
		}
	}
	return false, nil
}

func (t *Tree) findBlob(path string) (bool, *Blob) {
	for _, b := range t.Blob {
		if b.Path == path {
			return true, &b
		}
	}
	return false, nil
}

func (t *Tree) String() string {
	res := "tree : " + t.Name + "\n"
	for _, c := range t.Child {
		res += "child : " + c.Name + "\n"
	}
	for _, b := range t.Blob {
		res += "blob : " + b.Path + " " + b.Hex + " " + b.Status + "\n"
	}
	res += "---------\n"
	for _, c := range t.Child {
		res += c.String()
	}
	return res
}
