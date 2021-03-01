package stomp

import (
	"net/http"

	"gopkg.in/olahol/melody.v1"
)

// Socket ...
type Socket struct {
	connection *melody.Melody
	Broker     *Broker
}

// Init ...
func (Socket Socket) Init() Socket {
	Socket.connection = melody.New()
	Socket.connection.Upgrader.Subprotocols = []string{"v11.stomp"}
	Socket.Broker = &Broker{connection: Socket.connection, messageReceiver: func(Frame) {}}

	Socket.connection.HandleMessage(Socket.Broker.messageHandler)

	return Socket
}

// Upgrade ...
func (Socket Socket) Upgrade(writer http.ResponseWriter, request *http.Request) {
	Socket.connection.HandleRequest(writer, request)
}
