// Realiza conversoes de Celsius, Fahrenheit e Kelvin.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 0
)

func (c Celsius) String() string {
	return fmt.Sprint("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprint("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprint("%g°K", k)
}
