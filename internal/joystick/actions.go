package joystick

import "github.com/go-vgo/robotgo"

func (j JoystickController) ExecAction(value uint8, pressed bool) {
	j.ToggleTranslate(value, pressed)

	for _, button := range j.Buttons {
		if button.Value == value {
			if button.MouseFunction {
				mouseAction(pressed, button.KeyDown, button.KeyUp)
			} else {
				keyboardAction(pressed, button.KeyDown, button.KeyUp)
			}
		}
	}
}

func mouseAction(pressed bool, keyDown, keyUp string) {
	if pressed {
		robotgo.Click(keyDown)
	} else {
		robotgo.Click(keyUp)
	}
}

func keyboardAction(pressed bool, keyDown, keyUp string) {
	if pressed {
		robotgo.KeyDown(keyDown)
	} else {
		robotgo.KeyUp(keyUp)
	}
}
