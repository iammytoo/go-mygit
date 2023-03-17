package utils

import (
	"bytes"
	"compress/zlib"
)

func EncodeZlib(content string) bytes.Buffer {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte(content))
	w.Close()
	return b
}

func DecodeZlib(b *bytes.Buffer) string {
	r, _ := zlib.NewReader(b)
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	r.Close()
	return buf.String()
}
