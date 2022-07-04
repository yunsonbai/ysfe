package conf

import (
	"fmt"
	"ysfe/tools"

	"github.com/howeyc/gopass"
)

func InputIsNo(desc string) bool {
	yes := ""
	fmt.Println(desc + ": y/[n]")
	fmt.Scanln(&yes)
	for {
		if yes == string([]byte{78, 83}) || len(yes) == 0 {
			yes = "n"
		}
		if yes == "y" || yes == "n" {
			return yes == "n"
		}
		fmt.Println("请输入: y/[n]")
		fmt.Scanln(&yes)
	}
}

func InputIsYes(desc string) bool {
	yes := ""
	fmt.Println(desc + ": [y]/n")
	fmt.Scanln(&yes)
	for {
		if yes == string([]byte{78, 83}) || len(yes) == 0 {
			yes = "y"
		}
		if yes == "y" || yes == "n" {
			return yes == "y"
		}
		fmt.Println("请输入: [y]/n")
		fmt.Scanln(&yes)
	}
}

func inputPasswd(desc string) string {
	fmt.Println(desc)
	for {
		passwdByte, err := gopass.GetPasswd()
		tools.IsErrAndExit("", err, 0)
		if len(passwdByte) == 0 {
			fmt.Println("密码不能为空, 请重新输入:")
		} else {
			return string(passwdByte)
		}

	}
}
