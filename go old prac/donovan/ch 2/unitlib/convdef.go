package unitlib

import "fmt"

//Celsius type
type Celsius float64

//Fahrenheit type
type Fahrenheit float64

//Feet type
type Feet float64

//Meters type
type Meters float64

//Pound type
type Pound float64

//Kilograms type
type Kilograms float64

//constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%g feet", f) }
func (m Meters) String() string     { return fmt.Sprintf("%g metres", m) }
func (p Pound) String() string      { return fmt.Sprintf("%g lbs", p) }
func (kg Kilograms) String() string { return fmt.Sprintf("%g kg", kg) }
