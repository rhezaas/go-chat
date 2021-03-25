package controllers

import (
	"go-chat/app/models"
	"go-chat/app/utils/dto"
	"go-chat/app/utils/helper"
	"go-chat/app/utils/interfaces"
	"go-chat/app/utils/types"
)

// UserController ...
type UserController struct {
	Stomp interfaces.Stomp
	Redis interfaces.Redis
}

// TopicRoute ...
func (User UserController) TopicRoute() types.TopicRoute {
	return types.TopicRoute{
		helper.TopicBuilder("/login", "id", "name", "role"): User.login,
	}
}

func (User UserController) login(params types.TopicParams, message string) {
	// check required params
	if params["id"] == "" {
		User.Stomp.SendError(`"id" is required`)
	}

	// check required params
	if params["name"] == "" {
		User.Stomp.SendError(`"name" is required`)
	}

	// create user object
	user := dto.User{
		DefaultUserGroup: dto.DefaultUserGroup{
			ID:   params["id"],
			Name: params["name"],
		},
	}

	// call user model
	userModel := models.UserModel{Redis: User.Redis}

	// insert to redis
	err := userModel.CreateUser(user)

	// send error if something happens
	if err != nil {
		User.Stomp.SendError(err.Error())
	}

	User.Stomp.SendQueue(
		helper.TopicBuilder("/login", types.TopicParams{
			"id":   params["id"],
			"name": params["name"],
		}),
		"",
	)
}
