package main

import "fmt"

type celsius float64
type fahrenheit float64

const (
	AbsoluteZeroC celsius = -273.15
	FreezingC     celsius = 0
	BoilingC      celsius = 100
)

func main() {
	fmt.Println(cToF(AbsoluteZeroC))
	fmt.Println(fToC(fahrenheit(0)))
}

func fToC(f fahrenheit) celsius { return celsius((f - 32) * 5 / 9) }
func cToF(c celsius) fahrenheit { return fahrenheit(c*9/5 + 32) }
