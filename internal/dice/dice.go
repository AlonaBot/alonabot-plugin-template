// Package dice 是掷骰的业务逻辑，不依赖 SDK，可以直接做单元测试。
package dice

import "math/rand/v2"

const (
	minSides = 2
	maxSides = 100
)

// Roll 返回 1～sides 之间的一个点数；sides 超出 2～100 时 ok 为 false。
func Roll(sides int64) (value int64, ok bool) {
	if sides < minSides || sides > maxSides {
		return 0, false
	}
	return rand.Int64N(sides) + 1, true
}
