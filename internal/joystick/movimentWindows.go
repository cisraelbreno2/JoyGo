//go:build windows
// +build windows

package joystick

/*
#cgo windows LDFLAGS: -luser32
#include "mouse_move_windows.c"
*/
import "C"
import (
	"github.com/go-vgo/robotgo"
)

func (j JoystickController) MouseScroller(axisX, axisY int16) {
	robotgo.Scroll(int(axisX/1000)*-1, int(axisY/1000)*-1)
}

func (j JoystickController) MouseMovement(x, y int) {
	C.move_mouse_analog(C.int(x), C.int(y))
}
