package webhoook_poster

import "testing"

func TestNew(t *testing.T) {
	poster := New(PostType_Telegram, Param{
		WebhookURL: "xxx",
		Keyword:    "aaa",
		ChatID:     -1001741261665,
	})

	poster.Send("what are you doing?")
}
