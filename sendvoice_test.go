package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendVoice(t *testing.T) {
	ass := assert.New(t)

	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendVoice(SendVoiceArgs{ChatID: testingID, VoicePath: "testfiles/testvoice.ogg",
		Duration: 500})
	ass.NoError(err)
	ass.NotEmpty(msg.Voice.FileID, "No audio fileID")
}

func TestResendVoice(t *testing.T) {
	const voiceID = "AwADBAADlwADjT7dDUiBdYAXI_x3Ag"

	ass := assert.New(t)
	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendVoice(SendVoiceArgs{ChatID: testingID, VoiceID: voiceID})
	ass.NoError(err)
	ass.Equal(voiceID, msg.Voice.FileID, "Resent voice ID does not match")
}
