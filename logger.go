package sdk

import (
	"fmt"
)

// Logf send formatted log message to StreamDeck SDK.
func (s *StreamDeck) Logf(format string, a ...interface{}) {
	s.send(&SendEvent{
		Event:   LogMessage,
		Payload: &SendEventPayload{Message: fmt.Sprintf(format, a...)},
	})
}

// Log send log message to StreamDeck SDK.
func (s *StreamDeck) Log(message string) {
	s.send(&SendEvent{
		Event:   LogMessage,
		Payload: &SendEventPayload{Message: message},
	})
}
