package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflection(t *testing.T) {
	ass := assert.New(t)

	args := SendPhotoArgs{ChatID: testingID, PhotoPath: "testfiles/testimage.png", Caption: "something"}

	_, _, err := toMultiPart(args)
	ass.NoError(err)
}
