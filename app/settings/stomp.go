package settings

import (
	"net/http"
	"regexp"
	"strings"

	"go-chat/app/stomp"
	"go-chat/app/utils/interfaces"
	"go-chat/app/utils/types"
)

// Stomp Class
type Stomp struct {
	socket stomp.Socket
	broker *stomp.Broker
	topic  types.TopicRoute
}

// Initialize ...
func (Stomp Stomp) Initialize() Stomp {
	Stomp.socket = stomp.Socket{}.Init()
	Stomp.broker = Stomp.socket.Broker

	return Stomp
}

// SetupControllers ...
func (Stomp Stomp) SetupControllers(controllers []interfaces.Controller) {
	Stomp.topic = Stomp.mergeControllers(controllers)
	Stomp.broker.GetMessage(Stomp.messageHandler)
}

// Upgrade ...
func (Stomp Stomp) Upgrade(writer http.ResponseWriter, request *http.Request) {
	Stomp.socket.Upgrade(writer, request)
}

// SendQueue ...
func (Stomp Stomp) SendQueue(topic string, message string) string {
	return Stomp.broker.SendQueue(topic, message)
}

// SendBroadcast ...
func (Stomp Stomp) SendBroadcast(topic string, message string) string {
	return Stomp.broker.SendBroadcast(topic, message)
}

// SendError ...
func (Stomp Stomp) SendError(message string) {
	Stomp.broker.SendError(message)
}

// PRIVATE ...
func (Stomp Stomp) messageHandler(frame stomp.Frame) {
	route := frame.Header["destination"]
	params := make(types.TopicParams)

	// process to get endpoint
	regEndp := regexp.MustCompile("[^&?]*")
	endpoint := regEndp.FindAllString(route, -1)[0]

	// process to get route parameters
	regRoute := regexp.MustCompile("[^&?]*?=[^&?]*")
	for i, param := range regRoute.FindAllString(route, -1) {
		p := strings.Split(param, "=")

		// add param
		params[p[0]] = p[1]

		// build new endpoint to match stomp topic
		if len(params) > 0 {
			if i == 0 {
				endpoint += "?" + p[0] + "=" + "{" + p[0] + "}"
			} else {
				endpoint += "&" + p[0] + "=" + "{" + p[0] + "}"
			}
		}
	}

	if method, isDestinationExists := Stomp.topic[endpoint]; isDestinationExists {
		method(params, frame.Content)
	} else {
		Stomp.broker.SendError("Unknown Destination")
	}
}

// PRIVATE ...
func (Stomp Stomp) mergeControllers(controllers []interfaces.Controller) types.TopicRoute {
	topic := make(types.TopicRoute)

	for _, controller := range controllers {
		topics := controller.TopicRoute()

		for route, fn := range topics {
			topic[route] = fn
		}
	}

	return topic
}
