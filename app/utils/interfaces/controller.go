package interfaces

import "go-chat/app/utils/types"

// Controller ...
type Controller interface {
	TopicRoute() types.TopicRoute
}
