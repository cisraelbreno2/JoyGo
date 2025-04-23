package main

import (
	"JoyGo/internal/joystick"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	controller := joystick.NewJoystickController()
	defer sdl.Quit()

	joystick.HandleConnect(controller)

}
