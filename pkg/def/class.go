package def

import "github.com/julwrites/BotPlatform/pkg/secrets"

// Struct definitions for bot

type UserData struct {
	Firstname string `firestore:"first_name"`
	Lastname  string `firestore:"last_name"`
	Username  string `firestore:"user_name"`
	Id        string `firestore:"id"`
	Type      string `firestore:"type"`   // Group/Individual
	Action    string `firestore:"action"` // Current action if any
	Config    string `firestore:"config"`
}

type MessageData struct {
	Id      string
	Chat    string
	Command string
	Message string
}

type Option struct {
	Text string
	Link string
}

type ResponseOptions struct {
	Inline  bool
	Options []Option
	Remove  bool
}

type ResponseData struct {
	Message     string
	Affordances ResponseOptions
}

type SessionData struct {
	Secrets      secrets.SecretsData
	Type         string
	Channel      string
	User         UserData
	Msg          MessageData
	Res          ResponseData
	ResourcePath string
}
