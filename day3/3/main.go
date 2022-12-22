package main

import (
	"fmt"
	"strconv"
)

var users = []map[string]string{
	{
		"id":   "1",
		"name": "arun",
		"addr": "北京",
		"tel":  "xxx1231213",
	},
}

func input(msg string) (userItem string) {

	fmt.Print(msg)
	_, err := fmt.Scan(&userItem)
	if err != nil {
		fmt.Println(err)
	}
	return userItem
}

func list() {
	fmt.Println("----------用户列表-------------")
	for _, v2 := range users {
		printUser(v2)
	}
	fmt.Println("-----------------------")
}

func addUser() {
	user := make(map[string]string)
	user["name"] = input("输入名字：")
	user["addr"] = input("输入地址：")
	user["tel"] = input("输入电话：")
	currentUsersLength := len(users)
	for _, v := range users {
		tmp, _ := strconv.Atoi(v["id"])
		if tmp > currentUsersLength {
			currentUsersLength = tmp
		}
	}
	user["id"] = strconv.Itoa(currentUsersLength + 1)
	users = append(users, user)
	fmt.Printf("成功增加一个用户：\n")
	list()
}

func printUser(v map[string]string) {
	fmt.Printf("ID:%s\tName:%s\tTel:%s\tAddr:%s\n", v["id"], v["name"], v["tel"], v["addr"])
}

func del() {
	user := make(map[string]string)
	user["id"] = input("输入要删除的ID：")
	for k, v := range users {
		if v["id"] == user["id"] {
			printUser(v)

			switch input("确认是否删除上述用户，输入y/n确认是否删除") {
			case "y", "Y":
				users = append(users[:k], users[k+1:]...)
				fmt.Printf("成功删除一个用户：")
				printUser(v)
				list()
				return

			case "n", "N":
				fmt.Println("取消删除，返回")
				return
			}
		}
	}
	fmt.Printf("没有找到对应的ID:%v，请重新输入\n", user["id"])
}

func modify() {
	user := make(map[string]string)
	user["id"] = input("输入要修改的ID：")
	for k, v := range users {
		if v["id"] == user["id"] {
			printUser(v)
			switch input("确认是否修改上述用户，输入y/n确认是否修改") {
			case "y", "Y":
				user["name"] = input("输入名字：")
				user["addr"] = input("输入地址：")
				user["tel"] = input("输入电话：")
				users[k] = user
				fmt.Printf("成功修改一个用户：")
				printUser(v)
				list()
				return

			case "n", "N":
				fmt.Println("取消修改，返回")
				return
			}
		}
	}
	fmt.Printf("没有找到对应的ID:%v，请重新输入\n", user["id"])
}

func menu() {
	fmt.Println("1.添加用户")
	fmt.Println("2.删除用户")
	fmt.Println("3.修改用户")
	fmt.Println("4.查询用户")
	fmt.Println("5.退出")

	key := input("请输入序号：")
	switch key {
	case "1":
		addUser()
	case "2":
		del()
	case "3":
		fmt.Println("修改用户")
		modify()
	case "4":
		list()
	case "5":
		fmt.Println("退出")
		return
	default:
		fmt.Println("输入错误，请重新输入")
	}

}

func main() {
	for {

		menu()
	}
}
