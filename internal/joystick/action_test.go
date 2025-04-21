package joystick

import (
	"testing"
)

type MockExecutor struct {
	clicks           []string
	keyDowns         []string
	keyUps           []string
	scrollX, scrollY int16
	moveX, moveY     int16
}

func (m *MockExecutor) Click(button string) {
	m.clicks = append(m.clicks, button)
}

func (m *MockExecutor) KeyDown(key string) {
	m.keyDowns = append(m.keyDowns, key)
}

func (m *MockExecutor) KeyUp(key string) {
	m.keyUps = append(m.keyUps, key)
}

func (m *MockExecutor) Scroll(x, y int16) {
	m.scrollX = x
	m.scrollY = y
}

func (m *MockExecutor) MoveMouseAnalog(x, y int16) {
	m.moveX = x
	m.moveY = y
}

func TestJoystickController_ExecAction_MouseClick(t *testing.T) {
	mock := &MockExecutor{}

	j := JoystickController{
		Executor: mock,
		Buttons: []Button{
			{Name: "ClickButton", Value: 1, KeyDown: "left", KeyUp: "left", MouseFunction: true},
		},
	}

	j.ExecAction(1, true)
	j.ExecAction(1, false)

	if len(mock.clicks) != 2 {
		t.Errorf("expected 2 clicks, target %d", len(mock.clicks))
	}
}

func TestJoystickController_ExecAction_Keyboard(t *testing.T) {
	mock := &MockExecutor{}

	j := JoystickController{
		Executor: mock,
		Buttons: []Button{
			{Name: "KeyA", Value: 2, KeyDown: "a", KeyUp: "a", MouseFunction: false},
		},
	}

	j.ExecAction(2, true)
	j.ExecAction(2, false)

	if len(mock.keyDowns) != 1 || len(mock.keyUps) != 1 {
		t.Errorf("expected 1 KeyDown e 1 KeyUp, target %d KeyDowns and %d KeyUps", len(mock.keyDowns), len(mock.keyUps))
	}
}
