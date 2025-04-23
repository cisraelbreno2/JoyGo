package joystick

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	interpretationEnable = true
	buttonRPress         = false
	buttonBackPress      = false
	buttonCommand        = map[string]uint8{
		"R":    8,
		"BACK": 4,
	}
)

func HandleConnect(controller *JoystickController) {
	var axisLeftX, axisLeftY, axisRightX, axisRightY int16

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			if interpretationEnable {
				switch e := event.(type) {
				case *sdl.ControllerButtonEvent:
					switch e.State {
					case sdl.PRESSED:
						disableOrEnableInterpreter(e.Button, true)
						controller.ExecAction(e.Button, true)
					case sdl.RELEASED:
						disableOrEnableInterpreter(e.Button, false)
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

			} else {
				switch e := event.(type) {
				case *sdl.ControllerButtonEvent:
					switch e.State {
					case sdl.PRESSED:
						disableOrEnableInterpreter(e.Button, true)
					case sdl.RELEASED:
						disableOrEnableInterpreter(e.Button, false)
					}
				}
			}

		}
		controller.Executor.Scroll(axisRightX, axisRightY)
		controller.Executor.MoveMouseAnalog(axisLeftX, axisLeftY)
	}
}

func disableOrEnableInterpreter(value uint8, pressed bool) {
	if buttonBackPress && buttonRPress && !pressed {

		fmt.Printf("Interpreter active: %s\n", !interpretationEnable)

		interpretationEnable = !interpretationEnable
	}

	if value == buttonCommand["BACK"] {
		buttonBackPress = pressed
	} else if value == buttonCommand["R"] {
		buttonRPress = pressed
	}
}
