package dto

// ChatRoom ...
type ChatRoom struct {
	Chat     Chat      `json:"chat"`
	Messages []Message `json:"messages"`
}
