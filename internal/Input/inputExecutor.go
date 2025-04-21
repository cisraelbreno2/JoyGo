package input

type Executor interface {
	Click(button string)
	KeyDown(key string)
	KeyUp(key string)
	Scroll(x, y int16)
	MoveMouseAnalog(x, y int16)
}
