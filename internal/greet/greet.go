// Package greet 提供 hello 命令：用配置的问候语问候。
package greet

import (
	"github.com/AlonaBot/AlonaBot/sdk"
	"github.com/AlonaBot/AlonaBot/sdk/message"
)

// ID 是子插件 ID，也是 config_schema.json 中本模块配置的键。
const ID = "greet"

// args 是 hello 命令的参数；指针字段是可选参数，省略时为 nil。
type args struct {
	Name *string
}

// Register 在插件上登记本模块的子插件与命令。
func Register(p *sdk.Plugin) {
	s := p.Sub(ID).Permission(sdk.User)
	cfg := s.Config()
	s.Command("hello").Description("Greet someone").Handle(func(c *sdk.Context, a args) (message.Content, error) {
		name := c.T("greet.everyone")
		if a.Name != nil {
			name = *a.Name
		}
		return message.Text(c.T("greet.reply", "greeting", cfg.String("greeting"), "name", name)), nil
	})
}
