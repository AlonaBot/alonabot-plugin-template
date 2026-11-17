// Package plugin 是 AlonaBot 插件模板：一个读取配置的问候命令和一个带可选参数的掷骰命令。
package plugin

import (
	"math/rand/v2"

	"github.com/AlonaBot/AlonaBot/sdk"
	"github.com/AlonaBot/AlonaBot/sdk/message"
)

const (
	minSides = 2
	maxSides = 100
)

// helloArgs 是 hello 命令的参数；指针字段是可选参数，省略时为 nil。
type helloArgs struct {
	Name *string
}

// rollArgs 是 roll 命令的参数；省略面数时使用配置的 dice_sides。
type rollArgs struct {
	Sides *int64
}

// New 创建插件并登记命令。每次调用返回独立的插件对象。
func New() *sdk.Plugin {
	p := sdk.NewPlugin()
	p.Command("hello").Description("Greet someone").Handle(hello)
	p.Command("roll").Alias("dice").Description("Roll a die").Handle(roll)
	return p
}

// hello 用配置的问候语回复，省略名字时问候所有人。
func hello(c *sdk.Context, args helloArgs) (message.Content, error) {
	name := c.T("hello.everyone")
	if args.Name != nil {
		name = *args.Name
	}
	return message.Text(c.T("hello.reply", "greeting", c.Config().String("greeting"), "name", name)), nil
}

// roll 掷一个骰子，面数超出范围时回复提示。
func roll(c *sdk.Context, args rollArgs) (message.Content, error) {
	sides := c.Config().Int("dice_sides")
	if args.Sides != nil {
		sides = *args.Sides
	}
	if sides < minSides || sides > maxSides {
		return message.Text(c.T("roll.invalid")), nil
	}
	value := rand.Int64N(sides) + 1
	return message.Text(c.T("roll.reply", "value", value, "sides", sides)), nil
}
