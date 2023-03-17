package gitstructs

import (
	"strconv"

	"github.com/mirito333/go-mygit/utils"
)

type Blob struct {
	size    int
	content string
	hex     string
}

func CreateBlob(size int, content string) Blob {
	format_content := "blob " + strconv.Itoa(size) + "\\0" + content
	hex := utils.Sha1(format_content)
	
	return Blob{size, content, hex}
}

func (b *Blob) String() string {
	return strconv.Itoa(b.size) + "," + b.content
}
