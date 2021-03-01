package settings

import (
	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct{}

// Initialize ...
func (Server Server) Initialize() *gin.Engine {
	server := gin.Default()

	return server
}
