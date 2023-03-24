package fslib

import (
	"os"
	"runtime"
	"strings"
	"syscall"
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

func IsValidRootDirLinux() bool {
	// linux: return error if root dir is deleted
	// mac: return nil(inode is valid even if it is deleted)
	_, err := os.Getwd()
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

var lastCwdInode uint64

func IsCwdChangedDarwin() bool {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cwdStat, err := os.Stat(cwd)
	if err != nil {
		// linux error
		if os.IsNotExist(err) {
			return true
		}
		panic(err)
	}
	// mac
	stat, ok := cwdStat.Sys().(*syscall.Stat_t)
	if !ok {
		panic("Not a syscall.Stat_t")
	}
	if lastCwdInode != stat.Ino {
		lastCwdInode = stat.Ino
		return true
	}
	return false
}
func IsCwdChanged() bool {
	// if runtime.GOOS == "windows" {
	if runtime.GOOS == "darwin" {
		return IsCwdChangedDarwin()
	} else if runtime.GOOS == "linux" {
		return !IsValidRootDirLinux()
	} else {
		panic(runtime.GOOS + " not supported")
	}

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
