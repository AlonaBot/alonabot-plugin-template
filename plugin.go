// Package plugin 是 AlonaBot 插件模板的入口。
//
// 全部命令都在 New 中集中注册，一眼就能看到插件提供了什么；处理函数放在 handler/ 中按功能分文件。
// 新增命令时在 handler/ 写处理函数，再在这里加一行注册。
package plugin

import (
	"github.com/AlonaBot/AlonaBot/sdk"

	"github.com/AlonaBot/alonabot-plugin-template/handler"
)

// New 创建插件并注册全部命令。每次调用返回独立的插件对象。
func New() *sdk.Plugin {
	p := sdk.NewPlugin()

	p.Command("hello").Description("Greet someone").Handle(handler.Hello)

	return p
}
