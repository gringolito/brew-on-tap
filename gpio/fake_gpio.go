package gpio

import "log"

type fakeGpio struct {
	pin   Pin
	mode  Mode
	value bool
}

func FakeGpio() Gpio {
	f := fakeGpio{}
	log.SetPrefix("[FakeGpio] ")
	return &f
}

func (f *fakeGpio) open(pin Pin, mode Mode) error {
	log.Printf("New FakeGpio %p in Pin %d Mode %s\n", f, pin, mode)
	f.pin = pin
	f.SetMode(mode)
	return nil
}

func (f *fakeGpio) Close() {
	log.Printf("Closing Pin %d", f.pin)
}

func (f *fakeGpio) Set() {
	log.Printf("Setting Pin %d", f.pin)
	f.value = true
}

func (f *fakeGpio) Clear() {
	log.Printf("Clearing Pin %d", f.pin)
	f.value = false
}

func (f *fakeGpio) Read() bool {
	log.Printf("Pin %d is %t", f.pin, f.value)
	return f.value
}

func (f *fakeGpio) Mode() Mode {
	log.Printf("Pin %d mode %s", f.pin, f.mode)
	return f.mode
}

func (f *fakeGpio) SetMode(mode Mode) {
	log.Printf("Setting Pin %d mode %s", f.pin, mode)
	f.mode = mode
}
