package plumbing

import (
	"os"

	"github.com/mirito333/go-mygit/gitstructs"
)

func CreateHashObject(path string) gitstructs.Blob {
	content_byte, _ := os.ReadFile(path)
	content := string(content_byte)
	blob, _ := gitstructs.CreateBlob(len(content), content)
	return blob
}

func SaveHashObject(blob gitstructs.Blob) error {
	git_path := ".go-mygit/objects/"
	object_path := git_path + blob.Hex
	os.WriteFile(object_path, blob.Compressed.Bytes(), 0644)
	return nil
}
