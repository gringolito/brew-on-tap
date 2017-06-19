package gpio

import (
	"testing"
)

func TestDigitalPin(t *testing.T) {
	var g Gpio = DigitalPin()
	if g == nil {
		t.Fatal("Digital Pin does not implement Gpio")
	}
}
