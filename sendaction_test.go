package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendChatAction(t *testing.T) {
	ass := assert.New(t)

	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	err = bot.SendChatAction(SendChatActionArgs{ChatID: testingID, Action: Typing})
	ass.NoError(err)
}
