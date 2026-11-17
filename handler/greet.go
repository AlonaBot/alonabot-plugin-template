// Package handler 是命令的处理函数，按功能分文件；注册集中在插件入口的 New 中。
package handler

import (
	"github.com/AlonaBot/AlonaBot/sdk"
	"github.com/AlonaBot/AlonaBot/sdk/message"
)

// HelloArgs 是 hello 命令的参数；指针字段是可选参数，省略时为 nil。
type HelloArgs struct {
	Name *string
}

// Hello 用配置的问候语回复，省略名字时问候所有人。
func Hello(c *sdk.Context, args HelloArgs) (message.Content, error) {
	name := c.T("greet.everyone")
	if args.Name != nil {
		name = *args.Name
	}
	greeting := c.Config().Object("greet").String("greeting")
	return message.Text(c.T("greet.reply", "greeting", greeting, "name", name)), nil
}
