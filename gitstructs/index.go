package gitstructs

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Index struct {
	Path   string
	Hash   string
	Status string
}

func ReadIndex() []Index {
	indexPath := "./.go-mygit/index"
	res := []Index{}
	f, _ := os.Open(indexPath)
	defer f.Close()
	bu := bufio.NewScanner(f)
	for bu.Scan() {
		line := bu.Text()
		splited_line := strings.Split(line, " ")
		res = append(res, Index{Path: splited_line[1], Hash: splited_line[2], Status: splited_line[0]})
	}
	// 多分sortしているので、順序はいい感じに守られる。多分。
	sort.SliceStable(res, func(i,j int) bool {return res[i].Path < res[j].Path})
	return res
}
