package main

import (
	"cmdb/command"
	_ "cmdb/init"
	"cmdb/user"
	"crypto/sha256"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

// 打印提示信息
func prompt() {
	prompts := command.Prompts

	fmt.Println(strings.Repeat("*", 10))
	fmt.Println("欢迎使用xxx系统")
	fmt.Println("0. 退出")
	for _, v := range prompts {
		fmt.Printf("%s. %s\n", v[0], v[1])
	}

	fmt.Println(strings.Repeat("*", 10))
}

const Password = "admin"

func main() {
	//Password使用sum256加密
	fmt.Printf("%x\n", sha256.Sum256([]byte(Password)))

	//校对密码
	var initPassword []byte

	//隐藏输入的密码
	fmt.Print("请输入密码：")
	initPassword, _ = term.ReadPassword(int(os.Stdin.Fd()))

	if sha256.Sum256(initPassword) != sha256.Sum256([]byte(Password)) {
		fmt.Println("密码错误")
		return
	}

	commands := command.Commands
END:
	for {
		prompt()
		cmd := user.Input("请输入指令的序号：")
		if command, ok := commands[cmd]; ok {
			command()
		} else if cmd == "0" {
			break END
		} else {
			fmt.Println("指令错误，请重新输入")
		}
	}
}
