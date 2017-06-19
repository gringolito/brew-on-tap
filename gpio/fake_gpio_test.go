package gpio

import (
	"testing"
)

func TestFakeGpio(t *testing.T) {
	var g Gpio = New(FakeGpio(), GPIO0, Output)
	if g == nil {
		t.Fatal("FakeGpio does not implement Gpio!")
	}
	g.Close()
}

func TestSetFakeGpio(t *testing.T) {
	g := New(FakeGpio(), GPIO0, Output)
	g.Set()
	if !g.Read() {
		t.Fatal("FakeGpio Set() Failed!")
	}
}

func TestClearFakeGpio(t *testing.T) {
	g := New(FakeGpio(), GPIO0, Output)
	g.Clear()
	if g.Read() {
		t.Fatal("FakeGpio Clear() Failed!")
	}
}

func TestSetModeFakeGpio(t *testing.T) {
	g := New(FakeGpio(), GPIO0, Output)
	if g.Mode() != Output {
		t.Fatal("FakeGpio does not created in Output!")
	}

	g.SetMode(Input)
	if g.Mode() != Input {
		t.Fatal("FakeGpio SetMode(Input) Failed!")
	}

	g.SetMode(PWM)
	if g.Mode() != PWM {
		t.Fatal("FakeGpio SetMode(PWM) Failed!")
	}

	g.SetMode(Output)
	if g.Mode() != Output {
		t.Fatal("FakeGpio SetMode(Output) Failed!")
	}
}
