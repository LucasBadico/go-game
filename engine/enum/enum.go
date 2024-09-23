package enum

// Step 1: Define a custom type for the `type` field
type KeyEvent string

// Step 2: Define constants for the allowed values
const (
	KeyEvent_NoKey      KeyEvent = "KeyEvent.ZERO"
	KeyEvent_ArrowUp    KeyEvent = "KeyEvent.ARROW.UP"
	KeyEvent_ArrowDown  KeyEvent = "KeyEvent.ARROW.DOW"
	KeyEvent_ArrowLeft  KeyEvent = "KeyEvent.ARROW.LEFT"
	KeyEvent_ArrowRight KeyEvent = "KeyEvent.ARROW.RIGHT"
	KeyEvent_Q          KeyEvent = "KeyEvent.Q"
	KeyEvent_SPACE      KeyEvent = "KeyEvent.SPACE"
	KeyEvent_ENTER      KeyEvent = "KeyEvent.ENTER"
)

type EventSource string
const (
	EventSource_KeyPress   EventSource = "EventSource.KEY.PRESS"
	EventSource_KeyDown  EventSource = "EventSource.KEY.DOW"
	EventSource_KeyUp EventSource = "EventSource.KEY.UP"
	// EventSource_ArrowRight EventSource = "EventSource.ARROW.RIGHT"
)
