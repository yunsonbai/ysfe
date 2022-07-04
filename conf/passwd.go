package conf

import (
	"fmt"
	"os"
	"ysfe/tools"
)

func CheckAndGetPassWord(desc string, passwdMD5 string) string {
	if desc != "" {
		fmt.Println(desc)
	}
	passwd := ""
	if passwdMD5 == "" {
		passwd = inputPasswd("请输入初始化密码:")
		passwd2 := inputPasswd("请再次输入初始化密码:")
		if passwd2 != passwd {
			fmt.Println("两次密码输入不一致")
			os.Exit(0)
		}
		return passwd
	} else {
		passwd = inputPasswd("请输入密码:")
		if tools.StrMd5(passwd) != passwdMD5 {
			fmt.Println("密码错误")
			os.Exit(0)
		}
	}
	return passwd
}
