# Unit testing

## Does not exist, but needs some way to be tested:

 - SetWebhook
 - LeaveChat

## Exist, but could be improved

 - TestSendAudio
 - TestSendContact
 - TestSendDocument
 - TestSendPhoto
 - TestSendVideo

All of the above would be improved by making the scope of the test use more of the arguments.

 - TestKickChatMember

Needs a reliable way to kick somebody. Maybe in the future it may have it's own megagroup of 5000 bots to kick.

 - TestSendChatAction
Since all tests run all at once, it's really hard to see if this one works.
