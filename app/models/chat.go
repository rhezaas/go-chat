package models

import (
	"go-chat/app/utils/dto"
	"go-chat/app/utils/helper"
	"go-chat/app/utils/interfaces"
)

// ChatModel ...
type ChatModel struct {
	Redis interfaces.Redis
}

// ChatList ...
func (Chat ChatModel) ChatList(userID string) ([]dto.ChatList, error) {
	_chatList := []dto.ChatList{}

	userIDKey := helper.KeyBuilder("user", userID)
	userChatContactListKey := helper.KeyBuilder(userIDKey, "chats")

	userChatContactList, err := Chat.Redis.SMembers(userChatContactListKey)

	if err != nil {
		return []dto.ChatList{}, err
	}

	for _, chatContact := range userChatContactList {
		contact, err := Chat.Redis.HGetAll(chatContact)
		messages, err := Chat.Redis.SMembers(helper.KeyBuilder(userIDKey, "chat", "contact", contact["id"]))
		message, err := Chat.Redis.HGetAll(helper.KeyBuilder("message", messages[len(messages)-1]))

		_chatList = append(_chatList, dto.ChatList{
			Contact: dto.User{
				ID:   contact["id"],
				Name: contact["name"],
			},
			Message: dto.Message{
				ID:        message["id"],
				Message:   message["message"],
				Date:      message["date"],
				MessageID: message["messageId"],
			},
		})

		if err != nil {
			return []dto.ChatList{}, err
		}
	}

	return _chatList, err
}

// ChatRoom ...
func (Chat ChatModel) ChatRoom() {

}
