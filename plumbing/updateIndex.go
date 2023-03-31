package plumbing

import (
	"fmt"
	"log"
	"os"

	"github.com/mirito333/go-mygit/gitstructs"
)

func UpdateIndex(path string, blob gitstructs.Blob) error {
	indexPath := "./.go-mygit/index"
	f, err := os.OpenFile(indexPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	index_formated := blob.Status + " " + path + " " + blob.Hex+"\n"
	fmt.Fprint(f, index_formated)
	return nil
}
