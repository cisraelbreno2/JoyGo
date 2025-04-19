package joystick

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	TranslateEnable = true
	buttonRPress    = false
	buttonBackPress = false

	buttonCommand = map[string]uint8{
		"R":    8,
		"BACK": 4,
	}
)

func HandleConnect(controller *JoystickController) {
	var axisLeftX, axisLeftY, axisRightX, axisRightY int16

	fmt.Println("Controle habilitado")
	for TranslateEnable {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.ControllerButtonEvent:
				switch e.State {
				case sdl.PRESSED:
					controller.ExecAction(e.Button, true)
				case sdl.RELEASED:
					controller.ExecAction(e.Button, false)
				}
			case *sdl.ControllerAxisEvent:
				switch e.Axis {
				case 0:
					axisLeftX = e.Value
				case 1:
					axisLeftY = e.Value
				case 2:
					axisRightX = e.Value
				case 3:
					axisRightY = e.Value
				}
			}
		}
		controller.MouseScroller(axisRightX, axisRightY)
		controller.MouseMovement(int(axisLeftX), int(axisLeftY))
	}
}

func HandleDisconnect(controller *JoystickController) {
	fmt.Println("Controle desabilitado")

	for !TranslateEnable {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.ControllerButtonEvent:
				switch e.State {
				case sdl.PRESSED:
					controller.ToggleTranslate(e.Button, true)
				case sdl.RELEASED:
					controller.ToggleTranslate(e.Button, false)
				}
			}
		}
	}
}

func (j JoystickController) ToggleTranslate(value uint8, pressed bool) {
	if buttonBackPress && buttonRPress && !pressed {
		TranslateEnable = !TranslateEnable
	}

	if value == buttonCommand["BACK"] {
		buttonBackPress = pressed
	} else if value == buttonCommand["R"] {
		buttonRPress = pressed
	}
}
