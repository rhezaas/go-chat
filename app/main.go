package main

import (
	"go-chat/app/controllers"
	"go-chat/app/settings"
	"go-chat/app/utils/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	server := settings.Server{}.Initialize()
	stomp := settings.Stomp{}.Initialize()
	redis := settings.Redis{}.Initialize()

	server.GET("/", func(c *gin.Context) {
		stomp.Upgrade(c.Writer, c.Request)
	})

	stomp.SetupControllers([]interfaces.Controller{
		controllers.UserController{Stomp: stomp, Redis: redis},
		controllers.ChatController{Stomp: stomp, Redis: redis},
	})

	server.Run()
}
