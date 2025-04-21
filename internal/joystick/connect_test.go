package joystick

import (
	"testing"
)

func TestDisableOrEnableInterpreterToggle(t *testing.T) {

	disableOrEnableInterpreter(4, true)
	disableOrEnableInterpreter(8, true)
	disableOrEnableInterpreter(8, false)

	if interpretationEnable != false {
		t.Errorf("expected interpretationEnable = false, target true")
	}

	disableOrEnableInterpreter(4, true)
	disableOrEnableInterpreter(8, true)
	disableOrEnableInterpreter(4, false)

	if interpretationEnable != true {
		t.Errorf("expected interpretationEnable = true, target false")
	}
}
