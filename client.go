package TeleGogo

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type client struct {
	token      string
	httpClient *http.Client
}

// GetUpdates receives incoming updates using long polling.
func (c *client) GetUpdates(options GetUpdatesOptions) ([]Update, TelegramError) {
	response, err := c.sendJSON(options)
	if err != nil {
		return nil, errToTelegramErr(err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, responseToTgError(response)
	}

	upd := updateResponse{}
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&upd); err != nil {
		return nil, errToTelegramErr(err)
	}
	return upd.Result, nil
}

// SetWebhook NOT TESTED. Use this method to specify a url and receive incoming updates via an outgoing
// webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
// containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a
// reasonable amount of attempts.
func (c *client) SetWebhook(args SetWebhookArgs) TelegramError {
	response, err := c.sendFile(args)
	if err != nil {
		return errToTelegramErr(err)
	}

	if response.StatusCode != http.StatusOK {
		return responseToTgError(response)
	}
	response.Body.Close()

	return nil
}

// DownloadFile downloads the specified file
func (c *client) DownloadFile(fileID, path string) TelegramError {
	/* Get a File object by ID */
	args := getFileArgs{FileID: fileID}
	response, err := c.sendJSON(args)
	if err != nil {
		return responseToTgError(response)
	}
	fileResp := fileResponse{}
	fileDecoder := json.NewDecoder(response.Body)
	defer response.Body.Close()
	if err = fileDecoder.Decode(&fileResp); err != nil {
		return errToTelegramErr(err)
	}

	get, err := createFileGet(c.token, fileResp.Result.FilePath)
	if err != nil {
		return errToTelegramErr(err)
	}
	resp, err := c.httpClient.Do(get)
	if err != nil {
		return errToTelegramErr(err)
	}
	physicalFile, err := os.Create(path)
	if err != nil {
		return errToTelegramErr(err)
	}
	defer physicalFile.Close()
	defer resp.Body.Close()

	_, err = io.Copy(physicalFile, resp.Body)
	if err != nil {
		return errToTelegramErr(err)
	}
	return nil
}

// WhoAmI A simple method for testing your bot's auth token. Requires no parameters.
// Returns basic information about the bot in form of a User object.
func (c *client) WhoAmI() (User, TelegramError) {
	request, err := createBotGet(c.token, "getMe")

	if err != nil {
		return User{}, errToTelegramErr(err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return User{}, errToTelegramErr(err)
	}

	if response.StatusCode != http.StatusOK {
		return User{}, responseToTgError(response)
	}
	tgResp := userResponse{}
	decoder := json.NewDecoder(response.Body)

	if err = decoder.Decode(&tgResp); err != nil {
		return tgResp.Result, errToTelegramErr(err)
	}
	response.Body.Close()

	return tgResp.Result, nil
}

// SendMessage sends a message with the specified arguments. On success returns the sent Message.
func (c *client) SendMessage(args SendMessageArgs) (Message, TelegramError) {
	return c.sendJSONMessage(args)
}

func (c *client) ForwardMessage(args ForwardMessageArgs) (Message, TelegramError) {
	response, err := c.sendJSON(args)
	if err != nil {
		return Message{}, errToTelegramErr(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return Message{}, responseToTgError(response)
	}

	return responseToMessage(response)
}

// SendPhoto Use this method to send photos. On success, the sent Message is returned.
func (c *client) SendPhoto(args SendPhotoArgs) (Message, TelegramError) {
	/* Decide whether this is a newly uploaded file or an old one. */
	if args.FileID == "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

func (c *client) SendAudio(args SendAudioArgs) (Message, TelegramError) {
	/* Decide if it's a new or existing file, based on user intent */
	if args.AudioPath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

// SendDocument Use this method to send general files. On success, the sent Message is returned.
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (c *client) SendDocument(args SendDocumentArgs) (Message, TelegramError) {
	/* Decide if the document is a resend or a new file, based on args */
	if args.DocumentPath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

// SendSticker Use this method to send .webp stickers. On success, the sent Message is returned.
func (c *client) SendSticker(args SendStickerArgs) (Message, TelegramError) {
	/* Decide if the sticker is a existing or a new sticker, based on args */
	if args.StickerPath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

// SendVideo Use this method to send video files,
// Telegram clients support mp4 videos (other formats may be sent as Document).
// On success, the sent Message is returned.
// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (c *client) SendVideo(args SendVideoArgs) (Message, TelegramError) {
	/* Decide if the video is an existing or a new video, based on args */
	if args.VideoPath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

// SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a
// playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS
// (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can
// currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (c *client) SendVoice(args SendVoiceArgs) (Message, TelegramError) {
	/* Decide if the voice is an existing or a new voice, based on args */
	if args.VoicePath != "" {
		return c.sendFileMessage(args)
	}
	return c.sendJSONMessage(args)
}

// SendLocation Use this method to send point on the map. On success, the sent Message is returned.
func (c *client) SendLocation(args SendLocationArgs) (Message, TelegramError) {
	return c.sendJSONMessage(args)
}

// SendVenue Use this method to send information about a venue. On success, the sent Message is returned.
func (c *client) SendVenue(args SendVenueArgs) (Message, TelegramError) {
	return c.sendJSONMessage(args)
}

// SendContact Use this method to send phone contacts. On success, the sent Message is returned.
func (c *client) SendContact(args SendContactArgs) (Message, TelegramError) {
	return c.sendJSONMessage(args)
}

// SendChatAction Use this method when you need to tell the user that something is happening
// on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot,
// Telegram clients clear its typing status).
func (c *client) SendChatAction(args SendChatActionArgs) TelegramError {
	response, err := c.sendJSON(args)
	if err != nil {
		return errToTelegramErr(err)
	}
	if response.StatusCode != http.StatusOK {
		return responseToTgError(response)
	}
	return nil
}

// GetUserProfilePhotos Use this method to get a list of profile pictures for a user.
func (c *client) GetUserProfilePhotos(args GetUserProfilePhotosArgs) (UserProfilePhotos, TelegramError) {
	response, err := c.sendJSON(args)
	if err != nil {
		return UserProfilePhotos{}, errToTelegramErr(err)
	}
	if response.StatusCode != http.StatusOK {
		return UserProfilePhotos{}, responseToTgError(response)
	}
	photos := responseProfilePhotos{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&photos)
	if err != nil {
		return photos.Result, errToTelegramErr(err)
	}
	return photos.Result, nil
}

// KickChatMember Use this method to kick a user from a group or a supergroup.
// In the case of supergroups, the user will not be able to return to the group on their own using invite
// links, etc., unless unbanned first. The bot must be an administrator in the group for this to work.
func (c *client) KickChatMember(args KickChatMemberArgs) TelegramError {
	response, err := c.sendJSON(args)
	if err != nil {
		return errToTelegramErr(err)
	}
	if response.StatusCode != http.StatusOK {
		return responseToTgError(response)
	}
	return nil
}

// LeaveChat Use this method for your bot to leave a group, supergroup or channel.
func (c *client) LeaveChat(chatToLeave Chat) TelegramError {
	response, err := c.sendJSON(leaveChatArgs{ChatID: chatToLeave.ID})
	if err != nil {
		return errToTelegramErr(err)
	}
	if response.StatusCode != http.StatusOK {
		return responseToTgError(response)
	}
	return nil
}

func responseToMessage(response *http.Response) (Message, TelegramError) {
	msg := messageReply{}
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		return msg.Result, errToTelegramErr(err)
	}
	defer response.Body.Close()
	return msg.Result, nil
}

// NewBot Creates a new Client
func NewBot(token string) (Client, error) {
	c := new(client)
	c.token = token
	c.httpClient = &http.Client{}
	return c, nil
}

// Client represents a bot in Telegram.
type Client interface {
	DownloadFile(fileID string, pathToSaveTo string) TelegramError
	WhoAmI() (User, TelegramError)
	GetUpdates(GetUpdatesOptions) ([]Update, TelegramError)
	SendMessage(SendMessageArgs) (Message, TelegramError)
	ForwardMessage(ForwardMessageArgs) (Message, TelegramError)
	SetWebhook(SetWebhookArgs) TelegramError
	SendPhoto(SendPhotoArgs) (Message, TelegramError)
	SendAudio(SendAudioArgs) (Message, TelegramError)
	SendDocument(SendDocumentArgs) (Message, TelegramError)
	SendSticker(SendStickerArgs) (Message, TelegramError)
	SendVideo(SendVideoArgs) (Message, TelegramError)
	SendVoice(SendVoiceArgs) (Message, TelegramError)
	SendLocation(SendLocationArgs) (Message, TelegramError)
	SendVenue(SendVenueArgs) (Message, TelegramError)
	SendContact(SendContactArgs) (Message, TelegramError)
	SendChatAction(SendChatActionArgs) TelegramError
	GetUserProfilePhotos(GetUserProfilePhotosArgs) (UserProfilePhotos, TelegramError)
	KickChatMember(KickChatMemberArgs) TelegramError
}
