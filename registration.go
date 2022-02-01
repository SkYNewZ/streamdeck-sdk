package sdk

// Info is sent to the plugin on registration process.
type Info struct {
	Application struct {
		Font            string `json:"font"`
		Language        string `json:"language"`
		Platform        string `json:"platform"`
		PlatformVersion string `json:"platformVersion"`
		Version         string `json:"version"`
	} `json:"application"`
	Plugin struct {
		UUID    string `json:"uuid"`
		Version string `json:"version"`
	} `json:"plugin"`
	DevicePixelRatio int `json:"devicePixelRatio"`
	Colors           struct {
		ButtonPressedBackgroundColor string `json:"buttonPressedBackgroundColor"`
		ButtonPressedBorderColor     string `json:"buttonPressedBorderColor"`
		ButtonPressedTextColor       string `json:"buttonPressedTextColor"`
		DisabledColor                string `json:"disabledColor"`
		HighlightColor               string `json:"highlightColor"`
		MouseDownColor               string `json:"mouseDownColor"`
	} `json:"colors"`
	Devices []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Size struct {
			Columns int `json:"columns"`
			Rows    int `json:"rows"`
		} `json:"size"`
		Type int `json:"type"`
	} `json:"devices"`
}

// register send a registration event to StreamDeck SDK.
func (s *StreamDeck) register(inRegisterEvent string) error {
	event := &SendEvent{
		Event: EventName(inRegisterEvent),
		UUID:  s.UUID,
	}

	return s.conn.WriteJSON(event)
}
