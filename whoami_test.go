package TeleGogo

import (
	"testing"
)

const (
	testBotToken = "232603277:AAFkfNN5LrwHvIiX39K2n3sZbmSW3-AkrA0"
	testingID    = "118857134"
	guyToKick    = 224584816
	groupID      = "-178279414"
)

func TestWhoAmI(t *testing.T) {
	bot, err := NewBot(testBotToken)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	who, err := bot.WhoAmI()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("I am", who.Username, "and my ID is", who.ID)
}
