package gitstructs

import (
	"bytes"
	"strconv"

	"github.com/mirito333/go-mygit/utils"
)

type Blob struct {
	Size       int
	Content    string
	Hex        string
	Compressed bytes.Buffer
	Status     string
	Path       string
}

func CreateBlob(size int, content string) (Blob, error) {
	format_content := "blob " + strconv.Itoa(size) + "\\0" + content
	hex := utils.Sha1(format_content)
	compressed := utils.EncodeZlib(format_content)
	return Blob{Size: size, Content: content, Hex: hex, Compressed: compressed}, nil
}

func (b *Blob) String() string {
	return strconv.Itoa(b.Size) + "," + b.Content
}
