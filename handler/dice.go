package handler

import (
	"github.com/AlonaBot/AlonaBot/sdk"
	"github.com/AlonaBot/AlonaBot/sdk/message"

	"github.com/AlonaBot/alonabot-plugin-template/internal/dice"
)

// RollArgs 是 roll 命令的参数；省略面数时使用配置的 dice.default_sides。
type RollArgs struct {
	Sides *int64
}

// Roll 掷一个骰子，面数超出范围时回复提示。
func Roll(c *sdk.Context, args RollArgs) (message.Content, error) {
	sides := c.Config().Object("dice").Int("default_sides")
	if args.Sides != nil {
		sides = *args.Sides
	}
	value, ok := dice.Roll(sides)
	if !ok {
		return message.Text(c.T("dice.invalid")), nil
	}
	return message.Text(c.T("dice.reply", "value", value, "sides", sides)), nil
}
