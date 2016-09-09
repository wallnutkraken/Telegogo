package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResendPhoto(t *testing.T) {
	bot, err := NewClient(testBotToken)
	assert.NoError(t, err)
	updates, err := bot.GetUpdates(GetUpdatesOptions{})
	assert.NoError(t, err)

	inconclusive := true /* Changed later if we find a photo to resend */
	var photoUpdate Update
	for _, update := range updates {
		if len(update.Message.Photo) > 0 {
			inconclusive = false
			photoUpdate = update
		}
	}
	if inconclusive {
		assert.FailNow(t, "Test inconclusive: no photo to forward")
	}
	msg, err := bot.SendPhoto(SendPhotoArgs{ChatID: testingID,
		FileID: photoUpdate.Message.Photo[0].FileID})
	assert.NoError(t, err)
	assert.NotZero(t, len(msg.Photo), "No photo in returned Message")
}

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
