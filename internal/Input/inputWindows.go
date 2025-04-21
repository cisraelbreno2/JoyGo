//go:build windows
// +build windows

package input

/*
#cgo windows LDFLAGS: -luser32
#include "mouse_move_windows.c"
#include "keyboard_windows.c"
*/
import "C"

type RobotWindowsExecutor struct{}

var DefaultExecutor Executor = RobotWindowsExecutor{}

func (r RobotWindowsExecutor) Click(button string) {
	C.mouseClick(C.CString(button))
}

func (r RobotWindowsExecutor) KeyDown(key string) {
	C.keyDown(C.CString(key))
}

func (r RobotWindowsExecutor) KeyUp(key string) {
	C.keyUp(C.CString(key))
}

func (r RobotWindowsExecutor) Scroll(x, y int16) {
	C.scrollMouse(C.int(x), C.int(y))
}

func (r RobotWindowsExecutor) MoveMouseAnalog(x, y int16) {
	C.move_mouse_analog(C.int(x), C.int(y))
}
