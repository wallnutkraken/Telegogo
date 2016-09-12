package TeleGogo

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForwardMessage(t *testing.T) {
	bot, err := NewBot(testBotToken)
	assert.NoError(t, err)
	sentMessage, err := bot.SendMessage(SendMessageArgs{ChatID: testingID, Text: "messageToForward"})
	assert.NoError(t, err)
	assert.NotZero(t, sentMessage.ID, "Returned message ID is 0")
	id := strconv.Itoa(sentMessage.Chat.ID)
	forwardedMsg, err := bot.ForwardMessage(ForwardMessageArgs{ChatID: id, FromChatID: id, MessageID: sentMessage.ID})
	assert.NoError(t, err)
	assert.NotEqual(t, forwardedMsg, Message{})
}
