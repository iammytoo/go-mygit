package utils

type Dir struct{
	Path string
	Files []string
}

// 今の所何もないけど、git-ignoreとadd -aの実装用に。いい感じになったらここを育てる

func CreateDir(path string) Dir{
	// いまはpathがfilepathのみ
	return Dir{path,[]string{path}}
}
