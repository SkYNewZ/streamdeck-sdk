package sdk

// Option describes a single option.
type Option func(*StreamDeck)

// WithDebug enables debug mode.
func WithDebug(debug bool) Option {
	return func(deck *StreamDeck) {
		deck.debug = debug
	}
}
