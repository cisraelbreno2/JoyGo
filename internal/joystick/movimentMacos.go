//go:build darwin
// +build darwin

package joystick

/*
#cgo darwin CFLAGS: -mmacosx-version-min=10.9
#cgo darwin LDFLAGS: -framework ApplicationServices
#include "mouse_move_macos.c"
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
