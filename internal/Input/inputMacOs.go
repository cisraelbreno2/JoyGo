//go:build darwin
// +build darwin

package input

/*
#cgo darwin CFLAGS: -mmacosx-version-min=10.9
#cgo darwin LDFLAGS: -framework ApplicationServices
#include "mouse_move_macos.c"
*/
import "C"
import "github.com/go-vgo/robotgo"

type RobotMacOsExecutor struct{}

var DefaultExecutor Executor = RobotMacOsExecutor{}

func (r RobotMacOsExecutor) Click(button string) {
	robotgo.Click(button)
}

func (r RobotMacOsExecutor) KeyDown(key string) {
	robotgo.KeyDown(key)
}

func (r RobotMacOsExecutor) KeyUp(key string) {
	robotgo.KeyUp(key)
}

func (r RobotMacOsExecutor) Scroll(x, y int16) {
	robotgo.Scroll(int(x), int(y))
}

func (r RobotMacOsExecutor) MoveMouseAnalog(x, y int16) {
	C.move_mouse_analog(C.int(x), C.int(y))
}
