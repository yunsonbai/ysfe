package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"time"
	"ysfe/conf"
	"ysfe/core"
	"ysfe/tools"
)

func init() {
	conf.Init()
}

func actionV() {
	time.Sleep(10 * time.Second)
	os.Remove(conf.Conf.TmpFilePath)
	fmt.Println("临时文件已回收")
	os.Exit(0)
}

func main() {
	if conf.Conf.ListOpt {
		files, _ := ioutil.ReadDir(conf.Conf.EFileDir)
		for _, file := range files {
			fmt.Printf("%-30s \t %s\n", file.Name()+"", file.ModTime().Local().Format("2006-01-02 15:04:05"))
		}
		return
	}

	if conf.Conf.DelOpt {
		if conf.InputIsYes("确认删除加密文件" + conf.Conf.SrcFileName) {
			os.Remove(conf.Conf.SrcFilePath)
		}
		return
	}

	action := conf.Conf.ActionOpt
	backup := false
	if action == "e" {
		if core.Encrypt(conf.Conf.DstFilePath, conf.Conf.SrcFilePath) {
			if conf.InputIsYes("加密完成, 是否删除原始文件") {
				os.Remove(conf.Conf.SrcFilePath)
			}
		}
		backup = true
	} else if action == "p" {
		fmt.Println("解密后内容为:")
		fmt.Println(core.DecryptPrint(conf.Conf.SrcFilePath))
	} else {
		c := make(chan os.Signal)
		// signal.Notify(c, os.Interrupt, os.Kill)
		signal.Notify(c)
		core.Decrypt(conf.Conf.TmpFilePath, conf.Conf.SrcFilePath)
		fmt.Println("已输出到临时文件:", conf.Conf.TmpFilePath)

		if action == "v" {
			fmt.Println("临时文件将在120秒后回收, 请尽快查看")
			go actionV()
			<-c
		} else if action == "u" {
			fmt.Println("编辑完临时文件, 请关闭(ctrl+c)进程来加密保存文件")
			<-c
			core.Encrypt(conf.Conf.DstFilePath, conf.Conf.TmpFilePath)
			fmt.Println("已更新加密文件" + conf.Conf.SrcFileName)
			backup = true
		}
		os.Remove(conf.Conf.TmpFilePath)
		fmt.Println("临时文件已回收")
	}
	fmt.Println("backup:", backup, conf.Conf.DstFilePath, conf.Conf.BackupFilePath)
	if backup {
		tools.FileCopy(conf.Conf.DstFilePath, conf.Conf.BackupFilePath)
	}
}
