// Package tempconv provides types and functions for conversions between Celsius, Fahrenheit and
// Kelivn.
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// Convenient temperature constants for each unit of measure.
const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
	AbsoluteZeroK Kelvin     = 0
	FreezingK     Kelvin     = 273.15
	BoilingK      Kelvin     = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	if c <= AbsoluteZeroC {
		return AbsoluteZeroF
	}
	return Fahrenheit(c*9/5 + 32)
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin {
	if c <= AbsoluteZeroC {
		return AbsoluteZeroK
	}
	return Kelvin(c - AbsoluteZeroC)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	if f <= AbsoluteZeroF {
		return AbsoluteZeroC
	}
	return Celsius((f - 32) * 5 / 9)
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin {
	if f <= AbsoluteZeroF {
		return AbsoluteZeroK
	}
	return Kelvin((f - AbsoluteZeroF) * 5 / 9)
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	if k <= AbsoluteZeroK {
		return AbsoluteZeroC
	}
	return Celsius(k) + AbsoluteZeroC
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	if k <= AbsoluteZeroK {
		return AbsoluteZeroF
	}
	return Fahrenheit(k*9/5) + AbsoluteZeroF
}

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var value float64
	var unit string
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invlid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name, default value and usage, and returns
// the address of the flag variable. The flag argument must have a quantity and a unit, e.g.,
// "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
