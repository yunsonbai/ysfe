package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"ysfe/tools"

	"gopkg.in/ini.v1"
)

func confFix(cfg *ini.File) {
	Conf.PasswdMD5 = cfg.Section("").Key("passwd_md5").String()
	backupPath := cfg.Section("").Key("backup_path").String()

	if Conf.PasswdMD5 == "" {
		Conf.Passwd = CheckAndGetPassWord("密码丢失, 请恢复密码", "")
		Conf.PasswdMD5 = tools.StrMd5(Conf.Passwd)
		cfg.Section("").Key("passwd_md5").SetValue(Conf.PasswdMD5)
		cfg.SaveTo(Conf.configPath)
	}
	if backupPath == "" {
		fmt.Println("备份文件路径丢失, 请恢复备份路径")
		if !InputIsYes("使用默认备份目录" + Conf.BackupDir) {
			fmt.Println("请输入备份目录:")
			fmt.Scanln(&Conf.BackupDir)
		}
		tools.IsErrAndExit("备份目录无法创建", tools.Mkdir(Conf.BackupDir), 0)
		cfg.Section("").Key("backup_path").SetValue(Conf.BackupDir)
		cfg.SaveTo(Conf.configPath)
	} else {
		Conf.BackupDir = backupPath
	}
}

func confLoad() {
	if !tools.FileExit(Conf.configPath) {
		files, err := ioutil.ReadDir(Conf.RootDir)
		tools.IsErrAndExit("未知异常", err, 0)
		if len(files) > 0 {
			if InputIsNo("配置文件丢失, 运行目录有文件, 依然初始化应用") {
				os.Exit(0)
			}
		}
		f, err := os.Create(Conf.configPath)
		tools.IsErrAndExit("无法初始化配置文件", err, 0)
		f.Close()
		cfg, err := ini.Load(Conf.configPath)
		tools.IsErrAndExit("配置文件无法加载", err, 0)
		fmt.Println("检测到您第一使用, 请完成初始化工作")
		Conf.Passwd = CheckAndGetPassWord("", "")
		Conf.PasswdMD5 = tools.StrMd5(Conf.Passwd)
		if !InputIsYes("使用默认备份目录" + Conf.BackupDir) {
			fmt.Println("请输入备份目录:")
			fmt.Scanln(&Conf.BackupDir)
		}
		cfg.Section("").Key("passwd_md5").SetValue(Conf.PasswdMD5)
		cfg.Section("").Key("backup_path").SetValue(Conf.BackupDir)
		cfg.SaveTo(Conf.configPath)
		fmt.Println("初始化完毕, 可以使用了")
		return
	}

	cfg, err := ini.Load(Conf.configPath)
	tools.IsErrAndExit("配置文件存在但无法解析请修复, 路径: "+Conf.configPath, err, 0)
	confFix(cfg)
	if Conf.Passwd == "" {
		Conf.Passwd = CheckAndGetPassWord("", Conf.PasswdMD5)
	}

}
