package TeleGogo

import (
	"testing"
)

func TestSendMessage_Basic(t *testing.T) {
	bot, err := NewClient(testBotToken)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	text := "TestSendMessage"
	args := SendMessageArgs{ChatID: testingID, Text: text}
	response, err := bot.SendMessage(args)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if response.Text != text {
		t.Error("Response message is not the same as sent\n")
		t.Error("Expected:", text)
		t.Error("Actual:", response.Text)
		t.FailNow()
	} else {
		t.Log(response.From.Username, "with ID", response.From.ID, "sent message", response.Text)
	}
}

func TestSendMessage_Markdown(t *testing.T) {
	bot, err := NewClient(testBotToken)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	text := "Hello this is `TestSendMessage_Markdown`"
	args := SendMessageArgs{ChatID: testingID, Text: text, ParseMode: "Markdown"}
	response, err := bot.SendMessage(args)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(response.Entities) != 0 {
		t.Log("Response has entities:")
		for _, entity := range response.Entities {
			t.Log("Type:", entity.Type, "Starts at index:", entity.Offset, "of length:", entity.Length)
			t.Log("Relevant text:", response.Text[entity.Offset:entity.Offset+entity.Length], "from total message:",
				response.Text)
		}
	} else {
		t.Error("Reply has no entities, therefore no code.")
		t.FailNow()
	}
}

func TestSendMessage_HTML(t *testing.T) {
	bot, err := NewClient(testBotToken)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	text := "Hello this is <b>TestSendMessage_HTML</b>, <i>ironically</i>"
	args := SendMessageArgs{ChatID: testingID, Text: text, ParseMode: "HTML"}
	response, err := bot.SendMessage(args)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log("Response:", response.Text)
	if len(response.Entities) != 0 {
		t.Log("Response has entities:")
		for _, entity := range response.Entities {
			t.Log("Type:", entity.Type, "Starts at index:", entity.Offset, "of length:", entity.Length)
			t.Log("Relevant text:", response.Text[entity.Offset:entity.Offset+entity.Length])
		}
	} else {
		t.Error("Reply has no entities, therefore no code.")
		t.FailNow()
	}
}
