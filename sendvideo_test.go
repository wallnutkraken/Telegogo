package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendVideo(t *testing.T) {
	ass := assert.New(t)
	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendVideo(SendVideoArgs{ChatID: testingID, VideoPath: "testfiles/testvideo.mp4"})
	ass.NoError(err)
	ass.NotEmpty(msg.Video.FileID, "No video")
}

func TestResendVideo(t *testing.T) {
	const videoID = "BAADBAADWQADjT7dDeoL8w4AAezcHwI"
	ass := assert.New(t)
	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	msg, err := bot.SendVideo(SendVideoArgs{ChatID: testingID, VideoID: videoID})
	ass.NoError(err)
	ass.Equal(videoID, msg.Video.FileID, "File ID is not the same")
}
