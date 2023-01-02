package init

import (
	"cmdb/command"
	"cmdb/user"
)

func init() {
	command.Registry("添加用户", user.AddUser)
	command.Registry("删除用户", user.Del)
	command.Registry("修改用户", user.Modify)
	command.Registry("查询用户", user.Quary)

}
