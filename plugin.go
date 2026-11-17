// Package plugin 是 AlonaBot 插件模板的入口：New 只负责把各功能模块组装到同一个插件上。
//
// 每个功能放在 internal/ 下自己的包里，用子插件登记命令，配置是 config_schema.json 中与子插件 ID 同名的 object。
// 新增功能时照 internal/greet 新建一个包，在这里调用它的 Register。
package plugin

import (
	"github.com/AlonaBot/alonabot-plugin-template/internal/dice"
	"github.com/AlonaBot/alonabot-plugin-template/internal/greet"

	"github.com/AlonaBot/AlonaBot/sdk"
)

// New 创建插件并登记各模块。每次调用返回独立的插件对象。
func New() *sdk.Plugin {
	p := sdk.NewPlugin()
	greet.Register(p)
	dice.Register(p)
	return p
}
