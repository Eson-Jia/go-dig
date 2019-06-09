package main

import (
	"errors"
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type celsiusFlag struct {
	Celsius
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KToC(k Kelvin) Celsius { return Celsius(k + -273.15) }

// Set 要想实现flag.Value接口必须实现 Set 和 String 方法,因为 celsiusFlag 内嵌了(embed) Celsius,而 Celsius 已经有 String 方法，
//所以 Celsius 有String方法。因为 Celsius有String方法，所以 *Celsius 有String方法。那么 *Celsius只需要再实现
// Set 方法就能实现 flag.Value 接口。
func (c *celsiusFlag) Set(s string) error {
	var value float64
	var flag string
	fmt.Sscanf(s, "%f%s", &value, &flag)
	switch flag {
	case "C", "°C":
		c.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		c.Celsius = KToC(Kelvin(value))
		return nil
	default:
		return errors.New("wrong unit")
	}
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	cf := &celsiusFlag{value}
	flag.CommandLine.Var(cf, name, usage)
	return &cf.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
