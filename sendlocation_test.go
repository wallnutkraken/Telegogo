package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendLocation(t *testing.T) {
	const latitude float32 = 35.852124
	const longitude float32 = -95.317949
	ass := assert.New(t)

	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendLocation(SendLocationArgs{ChatID: testingID,
		Longitude: longitude, Latitude: latitude})
	ass.NoError(err)
	ass.NotEmpty(msg.Location.Latitude, "No latitude")
	ass.NotEmpty(msg.Location.Longitude, "No longitude")
}
