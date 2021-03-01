package controllers

import (
	"go-chat/app/utils/helper"
	"go-chat/app/utils/interfaces"
	"go-chat/app/utils/types"
)

// ChatController ...
type ChatController struct {
	Stomp interfaces.Stomp
	Redis interfaces.Redis
}

// TopicRoute ...
func (Chat ChatController) TopicRoute() types.TopicRoute {
	return types.TopicRoute{
		helper.TopicBuilder("/chat", "userId"):              Chat.chatList,
		helper.TopicBuilder("/chat", "userId", "contactId"): Chat.chatRoom,
	}
}

func (Chat ChatController) chatList(params types.TopicParams, message string) {
	if params["userId"] == "" {
		Chat.Stomp.SendError("\"name\" is required")
	}
}

func (Chat ChatController) chatRoom(params types.TopicParams, message string) {
	if params["userId"] == "" {
		Chat.Stomp.SendError("\"userId\" is required")
	}

	if params["contactId"] == "" {
		Chat.Stomp.SendError("\"contactId\" is required")
	}
}
