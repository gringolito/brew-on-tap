package gpio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gringolito/brew-on-tap/utils/fileutils"
)

const (
	basePath     = "/sys/class/gpio"
	exportPath   = "/sys/class/gpio/export"
	unexportPath = "/sys/class/gpio/unexport"
)

var (
	set   = []byte{'1'}
	clear = []byte{'0'}
)

// GPIO to Pin number map for C.H.I.P. kernel 4.4.13. left empty when setUp() are implemented
var pinMap = map[Pin]int{
	GPIO0: 1013,
	GPIO1: 1014,
	GPIO2: 1015,
	GPIO3: 1016,
	GPIO4: 1017,
	GPIO5: 1018,
	GPIO6: 1019,
	GPIO7: 1020,
}

var mode2string = map[Mode]string{
	Input:  "in",
	Output: "out",
	PWM:    "pwm",
}

var string2mode = map[string]Mode{
	"in":  Input,
	"out": Output,
	"pwm": PWM,
}

var initialized = false

type digitalPin struct {
	number    int
	pinPath   string
	modePath  string
	valueFile *os.File
}

func setUp() error {
	/* TODO:
	 * Read pin values from GPIO eXpander base at address
	 * /sys/bus/i2c/drivers/pcf857x/2-0038/gpio/gpiochip1013/base
	 * and fill pinMap with correct pin values
	 */
	initialized = true
	return nil
}

func DigitalPin() Gpio {
	p := digitalPin{}
	return &p
}

func (p *digitalPin) open(pin Pin, mode Mode) (err error) {
	err = nil

	if !initialized {
		if err = setUp(); err != nil {
			return
		}
	}

	var ok bool
	p.number, ok = pinMap[pin]
	if !ok {
		err = errors.New(fmt.Sprintf("Cannot map pin %s", pin))
		return
	}

	p.pinPath = filepath.Join(basePath, fmt.Sprintf("gpio%d", p.number))

	if err = p.exportPin(); err != nil {
		return
	}

	if err = p.openValueFile(); err != nil {
		return
	}

	p.SetMode(mode)
	return
}

func (p *digitalPin) Close() {
	err := p.valueFile.Close()
	if err != nil {
		// TODO: Handle error
		return
	}

	if _, err = fileutils.WriteFile(unexportPath, "%d", p.number); err != nil {
		// TODO: Handle error
	}
	return
}

func (p *digitalPin) exportPin() (err error) {
	err = nil

	if _, stat := os.Stat(p.pinPath); os.IsNotExist(stat) {
		_, err = fileutils.WriteFile(exportPath, "%d", p.number)
	}
	return
}

func (p *digitalPin) openValueFile() (err error) {
	p.valueFile, err = os.OpenFile(filepath.Join(p.pinPath, "value"), os.O_RDWR, 0600)
	return
}

func (p *digitalPin) Set() {
	if _, err := p.valueFile.Write(set); err != nil {
		// TODO: handle error
	}
}

func (p *digitalPin) Clear() {
	if _, err := p.valueFile.Write(clear); err != nil {
		// TODO: handle error
	}
}

func (p *digitalPin) Read() bool {
	value := make([]byte, 1)
	if _, err := p.valueFile.ReadAt(value, 0); err != nil {
		// TODO: handle error
	}
	return value[0] == set[0]
}

func (p *digitalPin) Mode() Mode {
	bytes, _ := ioutil.ReadFile(p.modePath)
	mode, ok := string2mode[string(bytes)]
	if !ok {
		// TODO: handle error
	}
	return mode
}

func (p *digitalPin) SetMode(mode Mode) {
	modeStr, ok := mode2string[mode]
	if !ok {
		// TODO: handle error
		return
	}

	if _, err := fileutils.WriteFile(p.modePath, modeStr); err != nil {
		// TODO: handle error
	}
}
