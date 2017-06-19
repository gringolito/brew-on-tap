// gpio package defines how to handle General Porpuse IO interfaces
package gpio

type Pin uint8

const (
	GPIO0 Pin = iota
	GPIO1
	GPIO2
	GPIO3
	GPIO4
	GPIO5
	GPIO6
	GPIO7
)

func (p Pin) String() string {
	switch p {
	case GPIO0:
		return "GPIO0"
	case GPIO1:
		return "GPIO1"
	case GPIO2:
		return "GPIO2"
	case GPIO3:
		return "GPIO3"
	case GPIO4:
		return "GPIO4"
	case GPIO5:
		return "GPIO5"
	case GPIO6:
		return "GPIO6"
	case GPIO7:
		return "GPIO7"
	}
	return "Invalid GPIO"
}

type Mode uint8

const (
	Input Mode = iota
	Output
	PWM
)

func (m Mode) String() string {
	switch m {
	case Input:
		return "Input"
	case Output:
		return "Output"
	case PWM:
		return "PWM"
	}
	return "Invalid Mode"
}

func New(g Gpio, p Pin, m Mode) Gpio {
	_ = g.open(p, m)
	return g
}

type Gpio interface {
	Set()
	Clear()
	Read() bool
	Mode() Mode
	SetMode(Mode)
	Close()
	open(Pin, Mode) error
}
