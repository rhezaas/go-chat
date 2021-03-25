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
	chatLists := []dto.ChatList{}

	userIDKey := helper.KeyBuilder("user", userID)
	userChatIDsKey := helper.KeyBuilder(userIDKey, "chats")

	userChatIDLists, err := Chat.Redis.SMembers(userChatIDsKey)

	if err != nil {
		return []dto.ChatList{}, err
	}

	for _, chatID := range userChatIDLists {
		chat, err := Chat.Redis.HGetAll(chatID)
		messages, err := Chat.Redis.SMembers(helper.KeyBuilder(userIDKey, "chat", chat["id"]))
		message, err := Chat.Redis.HGetAll(helper.KeyBuilder("message", messages[len(messages)-1]))

		chatLists = append(chatLists, dto.ChatList{
			Chat: dto.Chat{
				ID:    chat["id"],
				Title: chat["name"],
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

	return chatLists, err
}

// ChatRoom ...
func (Chat ChatModel) ChatRoom() {

}
