package joystick

import (
	"github.com/go-vgo/robotgo"
	"math"
)

var prevSpeedX, prevSpeedY float64

func (j JoystickController) MouseScroller(axisX, axisY int16) {
	robotgo.Scroll(int(axisX/1000)*-1, int(axisY/1000)*-1)
}

func (j JoystickController) MouseMovement(x, y int) {
	const (
		deadzone     = 2000
		maxSpeed     = 15.0
		minSpeed     = 0.2
		acceleration = 0.5
		deceleration = 0.8
	)

	deltaX := float64(x)
	deltaY := float64(y)

	if math.Abs(deltaX) < deadzone {
		deltaX = 0
	}
	if math.Abs(deltaY) < deadzone {
		deltaY = 0
	}

	if deltaX == 0 && deltaY == 0 {
		prevSpeedX *= deceleration
		prevSpeedY *= deceleration

		if math.Abs(prevSpeedX) < 0.1 {
			prevSpeedX = 0
		}
		if math.Abs(prevSpeedY) < 0.1 {
			prevSpeedY = 0
		}

		robotgo.MoveSmoothRelative(int(prevSpeedX), int(prevSpeedY))
		return
	}

	length := math.Hypot(deltaX, deltaY)
	normX := deltaX / length
	normY := deltaY / length

	intensity := math.Min(length/32767.0, 1.0)
	targetSpeed := minSpeed + intensity*(maxSpeed-minSpeed)

	newSpeedX := prevSpeedX + (normX*targetSpeed-prevSpeedX)*acceleration
	newSpeedY := prevSpeedY + (normY*targetSpeed-prevSpeedY)*acceleration

	speedLength := math.Hypot(newSpeedX, newSpeedY)
	if speedLength > 0 {
		scale := targetSpeed / speedLength
		newSpeedX *= scale
		newSpeedY *= scale
	}

	prevSpeedX = newSpeedX
	prevSpeedY = newSpeedY

	robotgo.MoveSmoothRelative(int(prevSpeedX), int(prevSpeedY))
}
