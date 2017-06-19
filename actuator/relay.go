package actuator

import "github.com/gringolito/brew-on-tap/gpio"

type relay struct {
	gpio gpio.Gpio
}

func Relay() Actuator {
	r := relay{}
	return &r
}

func (r *relay) init(gpio gpio.Gpio) error {
	r.gpio = gpio
	return nil
}

func (r *relay) TurnOn() {
	if r.gpio.Read() != true {
		r.gpio.Set()
	}
}

func (r *relay) TurnOff() {
	if r.gpio.Read() != false {
		r.gpio.Clear()
	}
}

func (r *relay) Toggle() {
	if r.gpio.Read() == true {
		r.gpio.Clear()
	} else {
		r.gpio.Set()
	}
}

func (r *relay) Status() bool {
	return r.gpio.Read()
}
