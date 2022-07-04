package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func IsErrAndExit(desc string, err error, code int) {
	if err == nil {
		return
	}
	fmt.Println(desc, " ", err)
	os.Exit(code)
}

func StrMd5(str string) (retMd5 string) {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func ByteXOR(message []byte, keywords []byte) (result []byte) {
	messageLen := len(message)
	keywordsLen := len(keywords)

	for i := 0; i < messageLen; i++ {
		result = append(result, message[i]^keywords[i%keywordsLen])
	}
	return result
}

func DirExit(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			if s.IsDir() {
				return true
			}
		}
		return false
	}
	if s.IsDir() {
		return true
	}
	return false
}

func FileExit(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			if s.IsDir() {
				return false
			}
		}
		return false
	}
	if s.IsDir() {
		return false
	}
	return true
}

func Mkdir(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			if s.IsDir() {
				return nil
			} else {
				return os.MkdirAll(path, 0711)
			}
		}
		return os.MkdirAll(path, 0711)
	}
	if s.IsDir() {
		return nil
	} else {
		return os.MkdirAll(path, 0711)
	}
}

func FileCopy(src, dst string) {
	source, err := os.Open(src)
	if err != nil {
		return
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return
	}

	defer destination.Close()
	io.Copy(destination, source)
}
