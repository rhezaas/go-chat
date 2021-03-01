package dto

// ChatList ...
type ChatList struct {
	Contact User    `json:"contact"`
	Message Message `json:"message"`
}
