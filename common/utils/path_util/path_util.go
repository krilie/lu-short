package path_util

import (
	"errors"
	"os"
	"strings"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetWd 获取当前目录的上级上上级目录
func GetRootWd(rootDep int) string {
	dir, err := os.Getwd()
	panicOnErr(err)
	for i := 0; i < rootDep; i++ {
		index := strings.LastIndex(dir, string(os.PathSeparator))
		if index == -1 {
			panic(errors.New("err path"))
		}
		dir = dir[:index]
	}
	return dir
}

// GetWd 获取当前目录的上级上上级目录
func GetRootWdBefore(beforeModuleName string) string {
	dir, err := os.Getwd()
	panicOnErr(err)
	index := strings.LastIndex(dir, beforeModuleName)
	if index == -1 {
		panic(errors.New("bad path string"))
	}
	return dir[:index-1]
}
