package sdk

// Option describes a single option.
type Option func(*StreamDeck)

// WithDebug enables debug mode.
func WithDebug(debug bool) Option {
	return func(deck *StreamDeck) {
		deck.debug = debug
	}
}

// WithErrorHandler receives errors instead of the Stream Deck log, which is
// unreachable once the connection is lost. It may be called concurrently.
func WithErrorHandler(h func(error)) Option {
	return func(deck *StreamDeck) {
		deck.onError = h
	}
}
