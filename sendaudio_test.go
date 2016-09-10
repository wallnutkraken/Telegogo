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
	const audioID = "BQADBAADKAADjT7dDeoWcahVii-dAg"
	bot, err := NewClient(testBotToken)
	assert.NoError(t, err)

	msg, err := bot.SendAudio(SendAudioArgs{ChatID: testingID, AudioFileID: audioID})
	assert.NoError(t, err)
	assert.NotZero(t, msg.Audio.FileSize, "No audio in returned message")
}
