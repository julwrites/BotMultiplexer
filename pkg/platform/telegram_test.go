package platform

import (
	"encoding/json"
	"testing"

	"github.com/julwrites/BotMultiplexer/pkg/def"
)

func GenerateTestData() []byte {
	var req TelegramRequest
	req.Message.Id = 9876
	req.Message.Chat.Id = 4567
	req.Message.Text = "/Command Text"
	req.Message.Sender.Id = 1234
	req.Message.Sender.Firstname = "First"
	req.Message.Sender.Lastname = "Last"
	req.Message.Sender.Username = "User"

	data, _ := json.Marshal(req)

	return data
}

func TestTelegramTranslate(t *testing.T) {
	data := GenerateTestData()
	env := TelegramTranslate(data)

	if env.Type != def.TYPE_TELEGRAM {
		t.Errorf("Failed TestTelegramTranslate, Type is wrong")
	}
	if env.Msg.Command != "Command" {
		t.Errorf("Failed TestTelegramTranslate, Msg Command is wrong")
	}
	if env.Msg.Message != "Text" {
		t.Errorf("Failed TestTelegramTranslate, Msg Text is wrong")
	}
	if env.Msg.Id != "9876" {
		t.Errorf("Failed TestTelegramTranslate, Msg ID is wrong")
	}
	if env.Channel != "4567" {
		t.Errorf("Failed TestTelegramTranslate, Channel ID is wrong")
	}
	if env.User.Id != "1234" {
		t.Errorf("Failed TestTelegramTranslate, User ID is wrong")
	}
	if env.User.Firstname != "First" {
		t.Errorf("Failed TestTelegramTranslate, User First name is wrong")
	}
	if env.User.Lastname != "Last" {
		t.Errorf("Failed TestTelegramTranslate, User Last name is wrong")
	}
	if env.User.Username != "User" {
		t.Errorf("Failed TestTelegramTranslate, User Username is wrong")
	}
}
