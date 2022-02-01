package sdk

// HandlerFunc receive event from StreamDeck SDK.
type HandlerFunc func(event *ReceivedEvent) error

// Handler same as HandlerFunc but with interface.
type Handler interface {
	Handle(event *ReceivedEvent) error
}

// Handler register given handlers.
func (s *StreamDeck) Handler(h ...Handler) {
	for _, handler := range h {
		s.handlers = append(s.handlers, handler.Handle)
	}
}

// HandlerFunc register given HandlerFunc.
func (s *StreamDeck) HandlerFunc(h ...HandlerFunc) {
	s.handlers = append(s.handlers, h...)
}
