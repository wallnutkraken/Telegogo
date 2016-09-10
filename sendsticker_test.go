package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendSticker(t *testing.T) {
	ass := assert.New(t)
	bot, err := NewClient(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendSticker(SendStickerArgs{ChatID: testingID,
		StickerPath: "testfiles/teststicker.webp"})
	ass.NoError(err)

	ass.NotEqual("", msg.Sticker.FileID, "No sticker in response")
}

func TestResendSticker(t *testing.T) {
	const stickerID = "BQADBAADrQIAAq6dFQcC08UCtLq4AAEC"
	ass := assert.New(t)
	bot, err := NewClient(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendSticker(SendStickerArgs{ChatID: testingID, StickerID: stickerID})
	ass.NoError(err)
	ass.Equal(msg.Sticker.FileID, stickerID, "Sticker file IDs are not equal")
}
