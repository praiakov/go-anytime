// Pacote tempconv realiza conversões de Celsius e Fahrenheit
package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Metros float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	KelvinC       Kelvin  = 1
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
func (m Metros) String() string     { return fmt.Sprintf("%g M", m) }
