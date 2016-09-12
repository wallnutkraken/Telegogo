package TeleGogo

import (
	"testing"
)

func TestGetUpdates(t *testing.T) {
	bot, err := NewBot(testBotToken)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	opts := GetUpdatesOptions{Offset: 0, Limit: 100, Timeout: 0}
	updates, err := bot.GetUpdates(opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(updates) > 0 {
		for _, update := range updates {
			t.Log("Recieved message from", update.Message.From.Username, "with ID",
				update.Message.From.ID, "he said", update.Message.Text)
		}
	} else {
		t.Log("Inconclusive: no messages.")
	}
}
