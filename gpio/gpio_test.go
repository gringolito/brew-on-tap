package gpio

import (
	"testing"
)

func TestPinString(t *testing.T) {
	m := map[Pin]string{
		GPIO0: "GPIO0",
		GPIO1: "GPIO1",
		GPIO2: "GPIO2",
		GPIO3: "GPIO3",
		GPIO4: "GPIO4",
		GPIO5: "GPIO5",
		GPIO6: "GPIO6",
		GPIO7: "GPIO7",
	}

	for key, val := range m {
		s := key.String()
		if val != s {
			t.Fatalf("%s bad stringfied to %s", val, s)
		}
	}
}

func TestModeString(t *testing.T) {
	m := map[Mode]string{
		Input:  "Input",
		Output: "Output",
		PWM:    "PWM",
	}

	for key, val := range m {
		s := key.String()
		if val != s {
			t.Fatalf("%s bad stringfied to %s", val, s)
		}
	}
}
