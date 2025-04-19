package joystick

type Button struct {
	Name          string `json:"name"`
	Value         uint8  `json:"value"`
	KeyDown       string `json:"keyDown"`
	KeyUp         string `json:"keyUp"`
	MouseFunction bool   `json:"mouseFunction"`
}
