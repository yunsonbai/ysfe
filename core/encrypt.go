package core

import (
	"io/ioutil"
	"ysfe/conf"
	"ysfe/tools"
)

func encrypt(dst, src string) (bool, string) {
	c, err := ioutil.ReadFile(src)
	tools.IsErrAndExit("打开原始文件异常", err, 0)
	res := tools.ByteXOR(c, []byte(conf.Conf.Passwd))
	if dst != "" {

		err = ioutil.WriteFile(dst, res, 0644)
		tools.IsErrAndExit("加密异常", err, 0)
		return true, ""
	} else {
		return true, string(res)
	}
}

func Encrypt(dstFile, srcFile string) bool {
	res, _ := encrypt(dstFile, srcFile)
	return res
}

func Decrypt(dstFile, srcFile string) bool {
	res, _ := encrypt(dstFile, srcFile)
	return res
}

func DecryptPrint(srcFile string) string {
	_, c := encrypt("", srcFile)
	return c
}
