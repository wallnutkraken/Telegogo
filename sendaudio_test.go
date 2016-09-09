package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendAudio(t *testing.T) {
	bot, err := NewClient(testBotToken)
	assert.NoError(t, err)

	expectedTitle := "testerino"
	msg, err := bot.SendAudio(SendAudioArgs{ChatID: testingID, AudioPath: "testfiles/testaudio.mp3",
		Title: expectedTitle})
	assert.NoError(t, err)
	assert.Equal(t, expectedTitle, msg.Audio.Title)
	assert.NotZero(t, msg.Audio.FileSize, "File size of audio is 0")
}

func TestResendAudio(t *testing.T) {
	bot, err := NewClient(testBotToken)
	assert.NoError(t, err)

	updates, err := bot.GetUpdates(GetUpdatesOptions{})
	assert.NoError(t, err, "Could not get updates; test inconclusive")

	var audio *Audio
	for _, update := range updates {
		if update.Message.Audio.Duration != 0 {
			audio = &update.Message.Audio
		}
	}
	assert.NotNil(t, audio, "No audio files in updates; inconclusive")
	msg, err := bot.SendAudio(SendAudioArgs{ChatID: testingID, AudioFileID: audio.FileID})
	assert.NoError(t, err)
	assert.NotZero(t, msg.Audio.FileSize, "No audio in returned message")
}
