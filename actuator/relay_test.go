package actuator

import (
	"testing"

	"github.com/gringolito/brew-on-tap/gpio"
)

func TestRelay(t *testing.T) {
	g := gpio.New(gpio.FakeGpio(), gpio.GPIO0, gpio.Output)
	var a Actuator = New(Relay(), g)
	if a == nil {
		t.Fatal("Relay does not implement Actuator")
	}
}

func TestTurnOnRelay(t *testing.T) {
	a := New(Relay(), gpio.New(gpio.FakeGpio(), gpio.GPIO0, gpio.Output))
	a.TurnOn()
	if !a.Status() {
		t.Fatal("Relay TurnOn() Failed!")
	}
}

func TestTurnOffRelay(t *testing.T) {
	a := New(Relay(), gpio.New(gpio.FakeGpio(), gpio.GPIO0, gpio.Output))
	a.TurnOff()
	if a.Status() {
		t.Fatal("Relay TurnOff() Failed!")
	}
}

func TestToggleRelay(t *testing.T) {
	a := New(Relay(), gpio.New(gpio.FakeGpio(), gpio.GPIO0, gpio.Output))
	// Starts relay in a known condition
	a.TurnOn()

	// Toggle off
	a.Toggle()
	if a.Status() {
		t.Fatal("Relay Toggle() failed to toggle off!")
	}

	// Toggle on
	a.Toggle()
	if !a.Status() {
		t.Fatal("Relay Toggle() failed to toogle on!")
	}
}
