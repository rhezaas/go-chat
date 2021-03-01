package stomp

const (
	// CONNECT stomp client command
	CONNECT = "CONNECT"
	// SEND stomp client command
	SEND = "SEND"
	// SUBSCRIBE stomp client command
	SUBSCRIBE = "SUBSCRIBE"
	// UNSUBSCRIBE stomp command
	UNSUBSCRIBE = "UNSUBSCRIBE"
	// DISCONNECT stomp client command
	DISCONNECT = "DISCONNECT"
	// CONNECTED stomp server command
	CONNECTED = "CONNECTED"
	// MESSAGE stomp server command
	MESSAGE = "MESSAGE"
	// ERROR stomp server command
	ERROR = "ERROR"
)
