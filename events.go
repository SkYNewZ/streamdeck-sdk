package sdk

// Alert sends an alert on the StreamDeck of action with given context.
func (s *StreamDeck) Alert(context string) {
	s.writeCh <- &SendEvent{Event: ShowAlert, Context: context}
}

// OpenURL tell the Stream Deck application to open an URL in the default browser.
func (s *StreamDeck) OpenURL(u string) {
	s.writeCh <- &SendEvent{Event: OpenURL, Payload: &SendEventPayload{URL: u}}
}

// SetTitle tell StreamDeck to dynamically change the title of action with given context.
func (s *StreamDeck) SetTitle(context string, title string, target Target) {
	s.writeCh <- &SendEvent{
		Event:   SetTitle,
		Context: context,
		Payload: &SendEventPayload{
			Title:  title,
			Target: target,
			//State: 0, TODO: state is not supported yet
		},
	}
}

// ShowOK temporarily show an OK checkmark icon of action with given context.
func (s *StreamDeck) ShowOK(context string) {
	s.writeCh <- &SendEvent{Event: ShowOk, Context: context}
}

// SetState change the state of an action supporting multiple states.
func (s *StreamDeck) SetState(context string, state uint8) {
	s.writeCh <- &SendEvent{
		Event:   SetState,
		Context: context,
		Payload: &SendEventPayload{State: state},
	}
}

// SetImage change the image of an action with given context.
func (s *StreamDeck) SetImage(context string, image string) {
	s.writeCh <- &SendEvent{
		Event:   SetImage,
		Context: context,
		Payload: &SendEventPayload{Image: image},
	}
}

// SetTriggerDescription change the trigger description of an action with given context.
func (s *StreamDeck) SetTriggerDescription(context string, payload *SendEventSetTriggerDescriptionPayload) {
	s.writeCh <- &SendEvent{
		Event:   SetTriggerDescription,
		Context: context,
		Payload: payload,
	}
}

// SetFeedback change the feedback of an action with given context.
func (s *StreamDeck) SetFeedback(context string, payload *SendEventSetFeedbackPayload) {
	s.writeCh <- &SendEvent{
		Event:   SetFeedback,
		Context: context,
		Payload: payload,
	}
}

// SetFeedbackLayout change the feedback layout of an action with given context.
func (s *StreamDeck) SetFeedbackLayout(context string, layout string) {
	s.writeCh <- &SendEvent{
		Event:   SetFeedbackLayout,
		Context: context,
		Payload: &SendEventSetFeedbackLayoutPayload{Layout: layout},
	}
}

// SendToPropertyInspector sends a payload to the property inspector for an action instance.
func (s *StreamDeck) SendToPropertyInspector(context string, payload map[string]interface{}) {
	s.writeCh <- &SendEvent{
		Event:   SendToPropertyInspector,
		Context: context,
		Payload: payload,
	}
}

// SetSettings change the settings of an action with given context.
func (s *StreamDeck) SetSettings(context string, settings map[string]interface{}) {
	s.writeCh <- &SendEvent{
		Event:   SetSettings,
		Context: context,
		Payload: &settings,
	}
}

// GetSettings get the settings of an action with given context.
// Settings will be sent back to the plugin as a ReceivedEvent with the event name DidReceiveSettings.
func (s *StreamDeck) GetSettings(context string) {
	s.writeCh <- &SendEvent{Event: GetSettings, Context: context}
}

// SetGlobalSettings change the global settings of an action with given context.
func (s *StreamDeck) SetGlobalSettings(context string, settings map[string]interface{}) {
	s.writeCh <- &SendEvent{
		Event:   SetGlobalSettings,
		Context: context,
		Payload: &settings,
	}
}

// GetGlobalSettings get the global settings of an action with given context.
// Global settings will be sent back to the plugin as a ReceivedEvent with the event name DidReceiveGlobalSettings.
func (s *StreamDeck) GetGlobalSettings(context string) {
	s.writeCh <- &SendEvent{Event: GetGlobalSettings, Context: context}
}

// SwitchProfile switches to a pre-defined profile distributed with the plugin on the specified device.
// The payload should contain "profile" (string) and optionally "page" (int, 0-indexed).
// When profile is omitted, Stream Deck switches back to the previous profile.
func (s *StreamDeck) SwitchProfile(context, device string, payload map[string]interface{}) {
	s.writeCh <- &SendEvent{
		Event:   SwitchToProfile,
		Context: context,
		Device:  device,
		Payload: payload,
	}
}
