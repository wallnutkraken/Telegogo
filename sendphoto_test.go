package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendPhoto(t *testing.T) {
	bot, err := NewClient(testBotToken)
	assert.NoError(t, err)

	expectedCaption := "Ho ho, ha ha!"
	msg, err := bot.SendPhoto(SendPhotoArgs{ChatID: testingID, PhotoPath: "testfiles/testimage.png",
		Caption: expectedCaption})
	assert.NoError(t, err)

	assert.NotZero(t, len(msg.Photo), "No photo in recieved message response.")
	assert.Equal(t, expectedCaption, msg.Caption)
}

func TestResendPhoto(t *testing.T) {
	const photoID = "AgADBAADv6cxG40-3Q19yuiXiGpU-8AmaRkABF-cEfxX8SFoWPIAAgI"
	bot, err := NewClient(testBotToken)
	assert.NoError(t, err)

	msg, err := bot.SendPhoto(SendPhotoArgs{ChatID: testingID, FileID: photoID})
	assert.NoError(t, err)
	assert.NotZero(t, len(msg.Photo), "No photo in returned Message")
}
