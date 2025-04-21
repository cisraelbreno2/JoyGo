package joystick

func (j JoystickController) ExecAction(value uint8, pressed bool) {

	for _, button := range j.Buttons {
		if button.Value == value {
			if button.MouseFunction {
				j.mouseAction(pressed, button.KeyDown, button.KeyUp)
			} else {
				j.keyboardAction(pressed, button.KeyDown, button.KeyUp)
			}
		}
	}
}

func (j JoystickController) mouseAction(pressed bool, keyDown, keyUp string) {
	if pressed {
		j.Executor.Click(keyDown)
	} else {
		j.Executor.Click(keyUp)
	}
}

func (j JoystickController) keyboardAction(pressed bool, keyDown, keyUp string) {
	if pressed {
		j.Executor.KeyDown(keyDown)
	} else {
		j.Executor.KeyUp(keyUp)
	}
}
