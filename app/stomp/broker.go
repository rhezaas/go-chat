package stomp

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/olahol/melody.v1"
)

// Broker ...
type Broker struct {
	connection      *melody.Melody
	session         *melody.Session
	messageReceiver func(Frame)
}

// GetMessage ...
func (Broker *Broker) GetMessage(fn func(Frame)) {
	Broker.messageReceiver = fn
}

// SendQueue ...
func (Broker *Broker) SendQueue(topic string, message string) string {
	uuid := uuid.NewString()

	stompMessage := Broker.parseResponseBody(Frame{
		Command: MESSAGE,
		Header: map[string]string{
			"destination":  topic,
			"subscription": fmt.Sprint(Broker.session.Keys[topic]),
			"content-type": "application/json",
			"message-id":   uuid,
		},
		Content: message,
	})

	Broker.session.Write(stompMessage)

	return uuid
}

// SendBroadcast ...
func (Broker *Broker) SendBroadcast(topic string, message string) string {
	uuid := uuid.NewString()

	stompMessage := Broker.parseResponseBody(Frame{
		Command: MESSAGE,
		Header: map[string]string{
			"destination":  topic,
			"subscription": fmt.Sprint(Broker.session.Keys[topic]),
			"content-type": "application/json",
			"message-id":   uuid,
		},
		Content: message,
	})

	Broker.connection.BroadcastBinaryFilter(stompMessage, func(session *melody.Session) bool {
		_, isExists := session.Get(topic)

		return isExists
	})

	return uuid
}

// SendError ...
func (Broker *Broker) SendError(message string) {
	stompMessage := Broker.parseResponseBody(Frame{
		Command: ERROR,
		Header: map[string]string{
			"content-type": "text/plain",
			"message":      message,
		},
		Content: message,
	})

	Broker.session.Write(stompMessage)
}

// PRIVATE ....
func (Broker *Broker) messageHandler(session *melody.Session, byteMessage []byte) {
	Broker.session = session
	stompMessage := Broker.parseRequestBody(byteMessage)

	switch {
	case stompMessage.Command == CONNECT:
		if stompMessage.Header["id"] == "" {
			Broker.SendError("Missing required header")
		} else {
			session.Write(Broker.parseResponseBody(Frame{
				Command: CONNECTED,
				Header: map[string]string{
					"version":    "1.1",
					"heart-beat": "0,0",
				},
				Content: "",
			}))
		}
	case stompMessage.Command == SUBSCRIBE:
		if stompMessage.Header["destination"] == "" {
			Broker.SendError("Missing required header")
		} else {
			session.Set(stompMessage.Header["destination"], stompMessage.Header["id"])
			Broker.messageReceiver(stompMessage)
		}
	case stompMessage.Command == SEND:
		if stompMessage.Header["destination"] == "" {
			Broker.SendError("Missing required header")
		} else {
			Broker.messageReceiver(stompMessage)
		}
	case stompMessage.Command == UNSUBSCRIBE:
		if stompMessage.Header["id"] == "" {
			Broker.SendError("Missing required header")
		} else {
			delete(session.Keys, stompMessage.Header["destination"])
		}
	case stompMessage.Command == DISCONNECT:
	default:
		Broker.SendError("Unhandled Stomp Frame")
	}
}

// PRIVATE ....
func (Broker *Broker) parseResponseBody(request Frame) []byte {
	body := ""
	null := []byte{0}

	body += request.Command + "\n"
	for key, value := range request.Header {
		body += key + ":" + value + "\n"
	}
	body += "\n"
	body += request.Content

	bodyBytes := []byte(body)

	return append(bodyBytes, null...)
}

// PRIVATE ....
func (Broker *Broker) parseRequestBody(response []byte) Frame {
	body := strings.Split(string(response), "\n")

	newResponse := Frame{}

	newResponse.Command = body[0]
	newResponse.Header = make(map[string]string)

	for i := 1; i < len(body); i++ {
		header := strings.Split(body[i], ":")
		if body[i] != "" && len(header) > 1 {
			newResponse.Header[header[0]] = header[1]
		}
	}

	newResponse.Content = body[len(body)-1]

	return newResponse
}
