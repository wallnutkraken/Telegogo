package TeleGogo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFile(t *testing.T) {
	const fileID = "AgADBAADnrkxG66dFQcjyYWumARVqK6HXhkABA6b5RipUWPSlyEBAAEC"
	const pathToSaveTo = "getfile_test_file"

	ass := assert.New(t)

	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	err = bot.DownloadFile(fileID, pathToSaveTo)
	if err != nil {
		ass.FailNow("Could not download file because of", err.Error())
	}

	stat, err := os.Stat(pathToSaveTo)
	ass.NoError(err)
	ass.NotZero(stat.Size(), "File size is 0")

	/* Cleanup */
	err = os.Remove(pathToSaveTo)
	ass.NoError(err, "Could not remove file")
}
