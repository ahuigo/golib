package fslib

import (
	"os"
	"strings"
)

func HasReadMode(fpath string) bool {
	fileInfo, err := os.Stat(fpath)
	if err != nil {
		return false
	}
	mode := fileInfo.Mode()
	return mode&(1<<2) != 0 //rwx
	// mode&os.ModePerm == os.ModePerm //0777
}

func IsValidRootDir() bool {
	rp, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_, err = os.Open(rp)
	if err != nil && os.IsNotExist(err) {
		return true
	}
	return false
}
func FixRootDir() {
	if rp, err := os.Getwd(); err == nil {
		os.Chdir(rp)
	}
}

func SafePath(path string) string {
	if strings.Contains(path, "../") {
		panic("bad path:" + path)
	}
	path = strings.TrimLeft(path, "/")
	return path
}
