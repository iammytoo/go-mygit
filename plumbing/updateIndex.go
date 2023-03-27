package plumbing

import (
	"os"

	"github.com/mirito333/go-mygit/gitstructs"
)

func UpdateIndex(path string, blob gitstructs.Blob) error {
	indexPath := "./.go-mygit/index"
	index, _ := os.Create(indexPath)
	index_formated := blob.Status + " " + path + " " + blob.Hex + "\n"
	index.Write([]byte(index_formated))
	return nil
}
