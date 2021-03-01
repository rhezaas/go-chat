package dto

// ChatRoom ...
type ChatRoom struct {
	Contact  User      `json:"contact"`
	Messages []Message `json:"messages"`
}
