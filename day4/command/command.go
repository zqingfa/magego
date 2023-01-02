package command

import (
	"strconv"
)

var Commands = map[string]func(){
	//"1": user.AddUser,
	//"2": user.Del,
	//"3": user.Modify,
	//"4": user.Quary,
}

// 打印欢迎信息
var Prompts = [][2]string{
	//"0": "退出",
	//"1": "添加用户",
	//"2": "删除用户",
	//"3": "修改用户",
	//"4": "查询用户",
}

var id = 1

func Registry(desc string, callbak func()) {
	Commands[strconv.Itoa(id)] = callbak
	//Prompts[strconv.Itoa(id)] = desc
	Prompts = append(Prompts, [2]string{strconv.Itoa(id), desc})
	id++

}
