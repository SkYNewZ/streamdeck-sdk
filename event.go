package sdk

// EventName is simply the name of the received event.
type EventName string

const (
	// [Events Received] (https://docs.elgato.com/sdk/plugins/events-received)

	// DidReceiveSettings Event received after calling the GetSettings API
	// to retrieve the persistent data stored for the action.
	DidReceiveSettings EventName = "didReceiveSettings"

	// DidReceiveGlobalSettings Event received after calling the GetGlobalSettings API
	// to retrieve the global persistent data.
	DidReceiveGlobalSettings EventName = "didReceiveGlobalSettings"

	// KeyDown When the user presses a key, the plugin will receive the KeyDown event.
	KeyDown EventName = "keyDown"

	// KeyUp When the user releases a key, the plugin will receive the keyUp event.
	KeyUp EventName = "keyUp"

	// When the user touches the display, the plugin will receive the touchTap event (SD+).
	TouchTap EventName = "touchTap"

	// When the user presses the encoder down, the plugin will receive the dialDown event (SD+).
	DialDown EventName = "dialDown"

	// When the user releases a pressed encoder, the plugin will receive the dialUp event (SD+).
	DialUp EventName = "dialUp"

	// When the user rotates the encoder, the plugin will receive the dialRotate event (SD+).
	DialRotate EventName = "dialRotate"

	// WillAppear When an instance of an action is displayed on the Stream Deck,
	// for example when the hardware is first plugged in,
	// or when a folder containing that action is entered, the plugin will receive a willAppear event.
	WillAppear EventName = "willAppear"

	// WillDisappear When an instance of an action ceases to be displayed on Stream Deck,
	// for example when switching profiles or folders,
	// the plugin will receive a willDisappear event.
	WillDisappear EventName = "willDisappear"

	// TitleParametersDidChange When the user changes the title or title parameters,
	// the plugin will receive a titleParametersDidChange event.
	TitleParametersDidChange EventName = "titleParametersDidChange"

	// DeviceDidConnect When a device is plugged to the computer,
	// the plugin will receive a deviceDidConnect event.
	DeviceDidConnect EventName = "deviceDidConnect"

	// DeviceDidDisconnect When a device is unplugged from the computer,
	// the plugin will receive a deviceDidDisconnect event.
	DeviceDidDisconnect EventName = "deviceDidDisconnect"

	// ApplicationDidLaunch When a monitored application is launched,
	// the plugin will be notified and will receive the applicationDidLaunch event.
	ApplicationDidLaunch EventName = "applicationDidLaunch"

	// ApplicationDidTerminate When a monitored application is terminated,
	// the plugin will be notified and will receive the applicationDidTerminate event.
	ApplicationDidTerminate EventName = "applicationDidTerminate"

	// SystemDidWakeUp When the computer is wake up,
	// the plugin will be notified and will receive the systemDidWakeUp event.
	SystemDidWakeUp EventName = "systemDidWakeUp"

	// Occurs when Stream Deck receives a deep-link message intended for the plugin.
	// The message is re-routed to the plugin, and provided as part of the payload.
	DidReceiveDeepLink EventName = "didReceiveDeepLink"

	// PropertyInspectorDidAppear Event received when the Property Inspector appears
	// in the Stream Deck software user interface,
	// for example when selecting a new instance.
	PropertyInspectorDidAppear EventName = "propertyInspectorDidAppear"

	// PropertyInspectorDidDisappear Event received when the Property Inspector for an instance is removed
	// from the Stream Deck software user interface,
	// for example when selecting a different instance.
	PropertyInspectorDidDisappear EventName = "propertyInspectorDidDisappear"

	// SendToPlugin (from the Property Inspector) Send a payload to the plugin.
	SendToPlugin EventName = "sendToPlugin"

	// [Events Sent] (https://docs.elgato.com/sdk/plugins/events-sent)

	// SetSettings Save data persistently for the action's instance.
	SetSettings EventName = "setSettings"

	// GetSettings Request the persistent data for the action's instance.
	GetSettings EventName = "getSettings"

	// SetGlobalSettings Save data securely and globally for the plugin.
	SetGlobalSettings EventName = "setGlobalSettings"

	// GetGlobalSettings Request the global persistent data.
	GetGlobalSettings EventName = "getGlobalSettings"

	// OpenURL Open a URL in the default browser.
	OpenURL EventName = "openUrl"

	// LogMessage Write a debug log to the logs file.
	LogMessage EventName = "logMessage"

	// SetTitle Dynamically change the title of an instance of an action.
	SetTitle EventName = "setTitle"

	// SetImage Dynamically change the image displayed by an instance of an action.
	SetImage EventName = "setImage"

	// Dynamically change properties of items on the Stream Deck + touch display.
	SetFeedback EventName = "setFeedback"

	// Dynamically change the current layout for the Stream Deck + touch display.
	SetFeedbackLayout EventName = "setFeedbackLayout"

	// Sets the trigger descriptions associated with an encoder (touch display + dial) action instance.
	// All descriptions are optional; when one or more descriptions are defined all descriptions are updated,
	// with undefined values having their description hidden in Stream Deck.
	// To reset the descriptions to the default values defined within the manifest,
	// an empty payload can be sent as part of the event.
	SetTriggerDescription EventName = "setTriggerDescription"

	// ShowAlert Temporarily show an alert icon on the image displayed by an instance of an action.
	ShowAlert EventName = "showAlert"

	// ShowOk Temporarily show an OK checkmark icon on the image displayed by an instance of an action.
	ShowOk EventName = "showOk"

	// SetState Change the state of the action's instance supporting multiple states.
	SetState EventName = "setState"

	// SwitchToProfile Switch to one of the preconfigured read-only profiles.
	SwitchToProfile EventName = "switchToProfile"

	// SendToPropertyInspector Send a payload to the Property Inspector.
	SendToPropertyInspector EventName = "sendToPropertyInspector"
)

// Device used with StreamDeck.
type Device uint8

const (
	// KESDSDKDeviceTypeStreamDeck Device type: Stream Deck.
	KESDSDKDeviceTypeStreamDeck Device = iota

	// KESDSDKDeviceTypeStreamDeckMini Device type: Stream Deck Mini.
	KESDSDKDeviceTypeStreamDeckMini

	// KESDSDKDeviceTypeStreamDeckXL Device type: Stream Deck XL.
	KESDSDKDeviceTypeStreamDeckXL

	// KESDSDKDeviceTypeStreamDeckMobile Device type: Stream Deck Mobile.
	KESDSDKDeviceTypeStreamDeckMobile

	// KESDSDKDeviceTypeCorsairGKeys Device type: Corsair G-Keys.
	KESDSDKDeviceTypeCorsairGKeys
)

// Controller used with StreamDeck.
type Controller string

const (
	Encoder Controller = "Encoder"
	KeyPad  Controller = "KeyPad"
)

// Layouts used with StreamDeck.
const (
	// IconLayout is a layout that displays an icon on the Stream Deck + touch display.
	// See [Icon Layout](https://docs.elgato.com/sdk/plugins/layouts-sd+#icon-layout-usdx1)
	IconLayout = "$X1"

	// CanvasLayout is a layout that displays a canvas on the Stream Deck + touch display.
	// See [Canvas Layout](https://docs.elgato.com/sdk/plugins/layouts-sd+#canvas-layout-usda0)
	CanvasLayout = "$A0"

	// ValueLayout is a layout that displays a value on the Stream Deck + touch display.
	// See [Value Layout](https://docs.elgato.com/sdk/plugins/layouts-sd+#value-layout-usda1)
	ValueLayout = "$A0"

	// IndicatorLayout is a layout that displays an indicator on the Stream Deck + touch display.
	// See [Indicator Layout](https://docs.elgato.com/sdk/plugins/layouts-sd+#indicator-layout-usdb1)
	IndicatorLayout = "$B1"

	// GradientIndicatorLayout is a layout that displays a gradient indicator on the Stream Deck + touch display.
	// See [Gradient Indicator Layout](https://docs.elgato.com/sdk/plugins/layouts-sd+#gradient-indicator-layout-usdb2)
	GradientIndicatorLayout = "$B2"

	// DoubleIndicatorLayout is a layout that displays a double indicator on the Stream Deck + touch display.
	// See [Double Indicator Layout](https://docs.elgato.com/sdk/plugins/layouts-sd+#double-indicator-layout-usdc1)
	DoubleIndicatorLayout = "$C1"
)

// ReceivedEvent describes a received event from StreamDeck SDK.
type ReceivedEvent struct {
	// The action's unique identifier. If your plugin supports multiple actions,
	// you should use this value to see which action was triggered.
	Action string `json:"action"`

	// The event name
	Event EventName `json:"event"`

	// An opaque value identifying the instance's action
	Context string `json:"context"`

	// An opaque value identifying the device.
	Device string `json:"device"`

	// A json object containing information about the device. Used on deviceDidConnect
	DeviceInfo struct {
		// The name of the device set by the user.
		Name string `json:"name"`

		// Type of device.
		Type Device `json:"type"`

		// The number of columns and rows of keys that the device owns.
		Size struct {
			Columns uint8 `json:"columns"`
			Rows    uint8 `json:"rows"`
		} `json:"size"`
	} `json:"deviceInfo"`

	// A json object containing context about received event
	Payload *ReceivedEventPayload `json:"payload"`
}

// ReceivedEventPayload describes a payload received from StreamDeck SDK.
type ReceivedEventPayload struct {
	// This json object contains data that you can set and is stored persistently.
	Settings map[string]interface{} `json:"settings"`

	// Contain value: 'Encoder' or 'KeyPad'.
	// Used on events: touchTap, dialDown, dialUp, dialRotate
	Controller Controller `json:"controller"`

	// The coordinates of the action triggered.
	Coordinates struct {
		Column uint8 `json:"column"`
		Row    uint8 `json:"row"`
	} `json:"coordinates"`

	// The array which holds (x, y) coordinates as a position of tap inside of LCD slot associated with action.
	// Used on events: touchTap
	TapPos [2]int `json:"tapPos"`

	// Boolean which is true when long tap happened.
	// Used on events: touchTap
	Hold bool `json:"hold"`

	// The integer which holds the number of "ticks" on encoder rotation.
	// Positive values are for clockwise rotation, negative values are for counterclockwise rotation, zero value is never happen
	// Used on events: dialRotate
	Ticks int `json:"ticks"`

	// Boolean which is true on rotation when encoder pressed
	// Used on events: dialRotate
	Pressed bool `json:"pressed"`

	// This is a parameter that is only set when the action has multiple states defined in its manifest.json.
	// The 0-based value contains the current state of the action.
	State uint8 `json:"state"`

	// Boolean indicating if the action is inside a Multi Action.
	IsInMultiAction bool `json:"isInMultiAction"`

	// This is a parameter that is only set when the action is triggered with a specific value from a Multi Action.
	// For example if the user sets the Game Capture Record action to be disabled in a Multi Action, you would see the value 1.
	// Only the value 0 and 1 are valid.
	UserDesiredState uint8 `json:"userDesiredState"`

	// The new title. Used on events: titleParametersDidChange
	Title string `json:"title"`

	// A json object describing the new title parameters. Used on events: titleParametersDidChange
	TitleParameters struct {
		FontFamily     string `json:"fontFamily"`
		FontSize       int    `json:"fontSize"`
		FontStyle      string `json:"fontStyle"`
		FontUnderline  bool   `json:"fontUnderline"`
		ShowTitle      bool   `json:"showTitle"`
		TitleAlignment string `json:"titleAlignment"`
		TitleColor     string `json:"titleColor"`
	} `json:"titleParameters"`

	// The identifier of the application that has been launched.
	// Used on events: applicationDidLaunch, applicationDidTerminate
	Application string `json:"application"`
}

// Target is where you want to display the title.
type Target uint8

const (
	// HardwareAndSoftware displays the title on Streamdeck software and hardware.
	HardwareAndSoftware Target = iota

	// Hardware only displays the title on Streamdeck hardware.
	Hardware

	// Software only displays the title on Streamdeck software.
	Software
)

// SendEvent describes a event to send to StreamDeck SDK.
type SendEvent struct {
	// The action unique identifier.
	// Used on events: sendToPropertyInspector, sendToPlugin
	Action string `json:"action,omitempty"`

	// Used to register this plugin
	UUID string `json:"uuid,omitempty"`

	// The event name
	Event EventName `json:"event"`

	// An opaque value identifying the instance's action
	Context string `json:"context,omitempty"`

	// An opaque value identifying the device.
	Device string `json:"device,omitempty"`

	// A json object
	Payload interface{} `json:"payload,omitempty"`
}

// SendEventPayload describes a payload to send to StreamDeck SDK.
type SendEventPayload struct {
	// A URL to open in the default browser.
	// Used on events: openUrl
	URL string `json:"url,omitempty"`

	// A string to write to the logs file.
	// Used on events: logMessage
	Message string `json:"message,omitempty"`

	// The title to display. If there is no title parameter, the title is reset to the title set by the user.
	// Used on events: setTitle
	Title string `json:"title,omitempty"`

	// Specify where you want to display the title.
	// Used on events: setTitle, setImage
	Target Target `json:"target,omitempty"`

	// A 0-based integer value representing the state of an action with multiple states.
	// This is an optional parameter.
	// If not specified, the title is set to all states.
	// Used on events: setTitle, setImage, setState
	State uint8 `json:"state,omitempty"`

	// The image to display encoded in base64 with the image format declared in the mime type (PNG, JPEG, BMP, ...).
	// svg is also supported.
	// If no image is passed, the image is reset to the default image from the manifest.
	// Used on events: setImage
	Image string `json:"image,omitempty"`

	// The name of the profile to switch to. The name should be identical to the name provided in the manifest.json file.
	// Used on events: switchToProfile
	Profile string `json:"profile,omitempty"`
}

// SendEventSetTriggerDescriptionPayload describes a payload for setTriggerDescription event to send to StreamDeck SDK.
type SendEventSetTriggerDescriptionPayload struct {
	// Optional value that describes the long-touch interaction with the touch display.
	// When undefined the description will be hidden.
	LongTouch string `json:"longTouch,omitempty"`

	// Optional value that describes the push interaction with the dial.
	// When undefined the description will be hidden.
	Push string `json:"push,omitempty"`

	// Optional value that describes the rotate interaction with the dial.
	// When undefined the description will be hidden.
	Rotate string `json:"rotate,omitempty"`

	// Optional value that describes the touch interaction with the touch display.
	// When undefined the description will be hidden.
	Touch string `json:"touch,omitempty"`
}

// SendEventSetFeedbackPayload describes a payload for setFeedback event to send to StreamDeck SDK.
type SendEventSetFeedbackPayload map[string]interface{}

// SendEventSetFeedbackLayoutPayload describes a payload for setFeedbackLayout event to send to StreamDeck SDK.
type SendEventSetFeedbackLayoutPayload struct {
	// The layout to display on the Stream Deck + touch display.
	// See [Layouts](https://docs.elgato.com/sdk/plugins/layouts-sd+)
	Layout string `json:"layout"`
}
