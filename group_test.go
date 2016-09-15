package TeleGogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKickChatMember(t *testing.T) {
	const expectedAlreadyKicked = "Bad Request: USER_NOT_PARTICIPANT"
	ass := assert.New(t)
	bot, err := NewBot(testBotToken)
	ass.NoError(err)

	tgerr := bot.KickChatMember(KickChatMemberArgs{ChatID: groupID, UserID: guyToKick})
	/* Since bots can't add people to groups, this test will simply check if the user it's trying to
	 * kick, which used to be in the given group, cannot be kicked since he's not a member of the group.
	 * which, for all intents and purposes of the test, seems fine to me. */
	ass.Equal(expectedAlreadyKicked, tgerr.Error(), `Error is not "user is already kicked"`)
}
