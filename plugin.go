// Package plugin 是 AlonaBot 插件模板的入口，命令注册与处理函数都放在本文件中。
package plugin

import (
	"github.com/AlonaBot/AlonaBot/sdk"
	"github.com/AlonaBot/AlonaBot/sdk/message"
)

// New 创建插件并注册全部命令。每次调用返回独立的插件对象。
func New() *sdk.Plugin {
	p := sdk.NewPlugin()
	p.Command("hello").Description("Greet someone").Handle(Hello)
	return p
}

// HelloArgs 是 hello 命令的参数；指针字段是可选参数，省略时为 nil。
type HelloArgs struct {
	Name *string
}

// Hello 用配置的问候语回复，省略名字时问候所有人。
func Hello(c *sdk.Context, args HelloArgs) (message.Content, error) {
	name := "everyone"
	if args.Name != nil {
		name = *args.Name
	}
	greeting := c.Config().Object("greet").String("greeting")
	return message.Text(greeting + ", " + name + "!"), nil
}
