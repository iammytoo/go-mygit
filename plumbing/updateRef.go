package plumbing

import "os"

func UpdateRef(hash string) error{
	refPath := "./.go-mygit/refs/main"
	os.WriteFile(refPath, []byte(hash), 0644)
	return nil
}
