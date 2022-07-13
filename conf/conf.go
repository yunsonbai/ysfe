package conf

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"ysfe/tools"
)

type Config struct {
	ActionOpt string
	ListOpt   bool
	DelOpt    bool

	RootDir        string
	SrcFilePath    string
	SrcFileName    string
	DstFilePath    string
	TmpFilePath    string
	BackupDir      string
	BackupFilePath string
	EFileDir       string // 加密文件存放目录
	TmpDir         string
	configPath     string

	Passwd    string
	PasswdMD5 string
}

const VERSION = "0.9.1"

var Conf = Config{}
var usage = `Usage: ysfe [Options]

Options:
  -l  查看加密文件列表, 输入该值其他选项失效
  -d  删除目标加密文件
  -a  动作, e:加密 u:更新 v:查看解密内容 p:终端查看解密内容
    e -- 加密目标文件
    u -- 目标文件解密后放入临时文件, 关闭程序时加密临时文件并覆盖原加密文件
    v -- 目标文件解密后放入临时文件, 120秒后删除临时文件
    p -- 目标文件解密后直接从终端输出
  -f  要操作的目标文件

a为e时f为原始文件, 其他动作f为加密文件(通过-l获取加密文件列表);
当a不为e时, -f后边只需要输入文件名即可;

`

func arrangeOptions() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}
	help := flag.Bool("h", false, "")
	list := flag.Bool("l", false, "")
	d := flag.Bool("d", false, "")
	version := flag.Bool("v", false, "")
	a := flag.String("a", "", "")
	f := flag.String("f", "", "")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if *version {
		fmt.Println("version is", VERSION)
		os.Exit(0)
	}
	if *list {
		Conf.ListOpt = *list
	} else {
		Conf.SrcFilePath = *f
		if *d || *a != "" {
			if *f == "" {
				fmt.Println("请使用-f告知我你要操作的目标文件")
				os.Exit(0)
			}
		}
		srcDir, fileName := filepath.Split(Conf.SrcFilePath)
		if *d {
			Conf.DelOpt = *d
			if srcDir != "" && srcDir != Conf.EFileDir {
				fmt.Println("只能删除某个加密文件, 可以使用-l查看文件列表")
				os.Exit(0)
			}
			Conf.SrcFilePath = Conf.EFileDir + "/" + fileName
			if !tools.FileExit(Conf.SrcFilePath) {
				fmt.Println("只能删除某个加密文件, 可以使用-l查看文件列表")
				os.Exit(0)
			}
		} else if *a != "" {
			acs := [4]string{"e", "u", "v", "p"}
			errStr := "!!!! -a 只能为e/d/u/v/p\n"
			Conf.TmpFilePath = Conf.TmpDir + "/" + fileName
			Conf.SrcFileName = fileName
			if *a == "e" {
				if srcDir == Conf.EFileDir {
					fmt.Println("不可对加密路径下的文件加密")
					os.Exit(0)
				}
				if tools.FileExit(Conf.DstFilePath) {
					if InputIsNo(fileName + "文件已在加密文件列表中, 是否覆盖加密") {
						os.Exit(0)
					}
				}
				Conf.DstFilePath = Conf.EFileDir + "/" + fileName
			} else {
				if srcDir != "" && srcDir != Conf.EFileDir {
					fmt.Println("加密文件不存在, 可以使用-l查看文件列表")
					os.Exit(0)
				}
				Conf.SrcFilePath = Conf.EFileDir + "/" + fileName
				Conf.DstFilePath = Conf.EFileDir + "/" + fileName
			}

			if !tools.FileExit(Conf.SrcFilePath) {
				fmt.Println("源文件不存在")
				os.Exit(0)
			}

			for _, v := range acs {
				if v == *a {
					errStr = ""
					break
				}
			}
			if errStr != "" {
				fmt.Println(errStr)
				os.Exit(0)
			}
			Conf.ActionOpt = *a
		} else {
			fmt.Println("\n请告诉我你的操作\n")
			flag.Usage()
			os.Exit(0)
		}
	}
}

func ConfFileLoad() {
	confLoad()
	tools.Mkdir(Conf.BackupDir)
	tools.Mkdir(Conf.EFileDir)
	tools.IsErrAndExit("临时目录无法创建", tools.Mkdir(Conf.TmpDir), 0)
}

func Init() {
	u, err := user.Current()
	tools.IsErrAndExit("无法获得本用户目录", err, 0)
	Conf.RootDir = u.HomeDir + "/ysfe"
	Conf.EFileDir = Conf.RootDir + "/efile"
	Conf.TmpDir = u.HomeDir + "/ysfe/tmp"
	Conf.BackupDir = Conf.RootDir + "/backup"
	Conf.configPath = Conf.RootDir + "/conf.ini"

	tools.IsErrAndExit("运行目录无法创建", tools.Mkdir(Conf.RootDir), 0)
	arrangeOptions()
	fmt.Println("软件运行目录为:", Conf.RootDir)
	ConfFileLoad()
	Conf.BackupFilePath = Conf.BackupDir + "/" + Conf.SrcFileName
}
