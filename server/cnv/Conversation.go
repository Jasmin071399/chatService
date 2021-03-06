package cnv

import "../msg"

// Conversations is a slice of Conversation
var Conversations []Conversation

// Conversation is a private conversation between two Users
type Conversation struct {
	User1    string        `json:"user1"`
	User2    string        `json:"user2"`
	Messages []msg.Message `json:"messages"`
}
