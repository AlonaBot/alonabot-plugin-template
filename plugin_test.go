package plugin

import (
	"testing"

	"github.com/AlonaBot/AlonaBot/plugintest"
)

func TestCommands(t *testing.T) {
	p := New()
	h := plugintest.Start(t, plugintest.Plugin(p), plugintest.Config(p, `{"greeting":"Hi","dice_sides":2}`))
	user := plugintest.User{ID: "alice"}
	chat := plugintest.Private(user.ID)

	reply := func(text string) string {
		t.Helper()
		result := h.Text(chat, user, text)
		if len(result.Sends) != 1 {
			t.Fatalf("%s: sends = %+v, want 1", text, result)
		}
		return result.Sends[0].Text()
	}

	if got := reply("/hello Alona"); got != "Hi, Alona!" {
		t.Errorf("/hello Alona = %q", got)
	}
	if got := reply("/hello"); got != "Hi, everyone!" {
		t.Errorf("/hello = %q", got)
	}
	// 省略面数时使用配置的 dice_sides，别名 dice 与 roll 相同。
	if got := reply("/dice"); got != "You rolled 1 (1-2)." && got != "You rolled 2 (1-2)." {
		t.Errorf("/dice = %q", got)
	}
	if got := reply("/roll 1"); got != "The number of sides must be between 2 and 100." {
		t.Errorf("/roll 1 = %q", got)
	}
}
