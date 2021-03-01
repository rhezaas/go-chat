package interfaces

// Stomp Interface
type Stomp interface {
	SendQueue(topic string, message string) string
	SendBroadcast(topic string, message string) string
	SendError(message string)
}
