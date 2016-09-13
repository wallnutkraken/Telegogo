package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserProfilePhotos(t *testing.T) {
	ass := assert.New(t)
	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	photos, err := bot.GetUserProfilePhotos(GetUserProfilePhotosArgs{UserID: testingID})
	ass.NoError(err)

	ass.NotZero(photos.Count, "No photos returned")
}
