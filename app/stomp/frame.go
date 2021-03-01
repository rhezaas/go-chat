package stomp

// Frame ...
type Frame struct {
	Command string
	Header  map[string]string
	Content string
}
