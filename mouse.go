package tea

import (
	"github.com/charmbracelet/x/input"
)

// MouseButton represents the button that was pressed during a mouse message.
type MouseButton = input.MouseButton

// Mouse event buttons
//
// This is based on X11 mouse button codes.
//
//	1 = left button
//	2 = middle button (pressing the scroll wheel)
//	3 = right button
//	4 = turn scroll wheel up
//	5 = turn scroll wheel down
//	6 = push scroll wheel left
//	7 = push scroll wheel right
//	8 = 4th button (aka browser backward button)
//	9 = 5th button (aka browser forward button)
//	10
//	11
//
// Other buttons are not supported.
const (
	MouseNone       = input.MouseNone
	MouseLeft       = input.MouseLeft
	MouseMiddle     = input.MouseMiddle
	MouseRight      = input.MouseRight
	MouseWheelUp    = input.MouseWheelUp
	MouseWheelDown  = input.MouseWheelDown
	MouseWheelLeft  = input.MouseWheelLeft
	MouseWheelRight = input.MouseWheelRight
	MouseBackward   = input.MouseBackward
	MouseForward    = input.MouseForward
	MouseButton10   = input.MouseButton10
	MouseButton11   = input.MouseButton11
)

// MouseMsg represents a mouse message. This is a generic mouse message that
// can represent any kind of mouse event.
type MouseMsg interface {
	input.MouseEvent
}

// MouseClickMsg represents a mouse button click message.
type MouseClickMsg struct {
	input.MouseClickEvent
}

// MouseReleaseMsg represents a mouse button release message.
type MouseReleaseMsg struct {
	input.MouseReleaseEvent
}

// MouseWheelMsg represents a mouse wheel message event.
type MouseWheelMsg struct {
	input.MouseWheelEvent
}

// MouseMotionMsg represents a mouse motion message.
type MouseMotionMsg struct {
	input.MouseMotionEvent
}
