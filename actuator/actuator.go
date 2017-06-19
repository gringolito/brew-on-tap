package actuator

import "github.com/gringolito/brew-on-tap/gpio"

func New(a Actuator, g gpio.Gpio) Actuator {
	_ = a.init(g)
	return a
}

type Actuator interface {
	TurnOn()
	TurnOff()
	Toggle()
	Status() bool
	init(gpio.Gpio) error
}
