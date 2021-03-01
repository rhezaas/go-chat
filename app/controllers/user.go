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
		helper.TopicBuilder(`/login`, `id`, `name`, `role`): User.login,
	}
}

func (User UserController) login(params types.TopicParams, message string) {
	// check required params
	if params[`id`] == `` {
		User.Stomp.SendError(`id is required`)
	}

	// check required params
	if params[`name`] == `` {
		User.Stomp.SendError(`name is required`)
	}

	// check required params
	if params[`role`] == `` {
		User.Stomp.SendError(`role is required`)
	}

	// create user object
	user := dto.User{
		ID:   params[`id`],
		Name: params[`name`],
		Role: params[`role`],
	}

	// prep userRedis
	userRedis := models.UserModel{Redis: User.Redis}

	// insert to redis
	err := userRedis.CreateUser(user)

	// send error if something happens
	if err != nil {
		User.Stomp.SendError(err.Error())
	}

	// send message if user is agent
	if user.Role == `agent` {
		User.Stomp.SendQueue(
			helper.TopicBuilder(`/login`, types.TopicParams{
				`id`:   params[`id`],
				`name`: params[`name`],
				`role`: params[`role`],
			}),
			"",
		)
	}

	// send message if user is client
	if user.Role == `client` {
		// get random contact

		User.Stomp.SendQueue(
			helper.TopicBuilder(`/login`, types.TopicParams{
				`id`:   params[`id`],
				`name`: params[`name`],
				`role`: params[`role`],
			}),
			"",
		)
	}
}
