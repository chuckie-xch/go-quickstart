package main

type Celsius float64

type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -275.15
	FreeingC      Celsius = 0
	Boiling       Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
