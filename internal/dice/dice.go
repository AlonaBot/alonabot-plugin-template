// Package dice 提供 roll 命令（别名 dice）：掷一个骰子。
//
// 业务逻辑 Roll 不依赖 SDK，可以直接做单元测试；Register 只负责把它接到命令上。
package dice

import (
	"math/rand/v2"

	"github.com/AlonaBot/AlonaBot/sdk"
	"github.com/AlonaBot/AlonaBot/sdk/message"
)

// ID 是子插件 ID，也是 config_schema.json 中本模块配置的键。
const ID = "dice"

const (
	minSides = 2
	maxSides = 100
)

// args 是 roll 命令的参数；省略面数时使用配置的 default_sides。
type args struct {
	Sides *int64
}

// Roll 返回 1～sides 之间的一个点数；sides 超出 2～100 时 ok 为 false。
func Roll(r *rand.Rand, sides int64) (value int64, ok bool) {
	if sides < minSides || sides > maxSides {
		return 0, false
	}
	return r.Int64N(sides) + 1, true
}

// Register 在插件上登记本模块的子插件与命令。
func Register(p *sdk.Plugin) {
	s := p.Sub(ID).Permission(sdk.User)
	cfg := s.Config()
	r := rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))
	s.Command("roll").Alias("dice").Description("Roll a die").Handle(func(c *sdk.Context, a args) (message.Content, error) {
		sides := cfg.Int("default_sides")
		if a.Sides != nil {
			sides = *a.Sides
		}
		value, ok := Roll(r, sides)
		if !ok {
			return message.Text(c.T("dice.invalid")), nil
		}
		return message.Text(c.T("dice.reply", "value", value, "sides", sides)), nil
	})
}
