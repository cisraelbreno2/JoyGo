//go:build windows
// +build windows

package input

/*
#cgo windows LDFLAGS: -luser32
#include "mouse_move_windows.c"
*/
import "C"
import (
	"github.com/go-vgo/robotgo"
)

type RobotWindowsExecutor struct{}

var DefaultExecutor Executor = RobotWindowsExecutor{}

func (r RobotWindowsExecutor) Click(button string) {
	robotgo.Click(button)
}

func (r RobotWindowsExecutor) KeyDown(key string) {
	robotgo.KeyDown(key)
}

func (r RobotWindowsExecutor) KeyUp(key string) {
	robotgo.KeyUp(key)
}

func (r RobotWindowsExecutor) Scroll(x, y int16) {
	robotgo.Scroll(int(x), int(y))
}

func (r RobotWindowsExecutor) MoveMouseAnalog(x, y int16) {
	C.move_mouse_analog(C.int(x), C.int(y))
}
