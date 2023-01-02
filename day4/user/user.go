package user

import (
	//_ "cmdb/init"
	"fmt"
	"strconv"
)

var users []map[string]string

// 键盘输入功能
func input(msg string) (userItem string) {

	fmt.Print(msg)
	_, err := fmt.Scan(&userItem)
	if err != nil {
		fmt.Println(err)
	}
	return userItem
}

func Input(msg string) (userItem string) {

	fmt.Print(msg)
	_, err := fmt.Scan(&userItem)
	if err != nil {
		fmt.Println(err)
	}
	return userItem
}

// 打印用户列表
func printUser(v map[string]string) {
	fmt.Printf("ID:%s\tName:%s\tTel:%s\tAddr:%s\n", v["id"], v["name"], v["tel"], v["addr"])
}

func list() {
	fmt.Println("----------用户列表-------------")
	for _, v2 := range users {
		printUser(v2)
	}
	fmt.Println("-----------------------")
}

// AddUser 添加用户
func AddUser() {
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

// Del 删除用户
func Del() {
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

// Modify 修改用户信息
func Modify() {
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

// Quary 查询用户
func Quary() {
	user := make(map[string]string)
	user["id"] = input("输入要查询的ID：")
	for _, v := range users {
		if v["id"] == user["id"] {
			printUser(v)
			return
		}
	}
	fmt.Printf("没有找到对应的ID:%v，请重新输入\n", user["id"])
}
