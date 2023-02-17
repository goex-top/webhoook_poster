package webhoook_poster

import "testing"

func TestNew(t *testing.T) {
	poster := New(PostType_Feishu, Param{
		WebhookURL: "aaa",
		Keyword:    "aaa",
		ChatID:     0,
	})

	poster.Send("what are you doing?")
}
