package fileutil

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

func GetContentType(file io.ReadSeeker) (string, error) {
	decByte := make([]byte, 512)
	if _, err := file.Read(decByte); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	contentType := http.DetectContentType(decByte)
	return contentType, nil
}

func GetFileExtension(fileName string) string {
	index := strings.LastIndex(fileName, ".")
	if index == -1 {
		return ""
	}
	return fileName[index:]
}

// GetContentType2 从file中读出最多512字节用于确定类型 并所回一个新的reader供后续使用
func GetContentType2(file io.Reader) (mType string, newReader io.Reader, err error) {
	decByte := make([]byte, 512)
	if readLen, err := file.Read(decByte); err != nil {
		return "", file, err
	} else {
		contentType := http.DetectContentType(decByte[0:readLen])
		return contentType, io.MultiReader(bytes.NewReader(decByte[0:readLen]), file), nil
	}
}
