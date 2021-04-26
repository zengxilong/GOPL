package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func isUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagMulticast|FlagBroadcast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, isUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, isUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, isUp(v))
	fmt.Printf("%b %v\n", v, IsCast(v))
}
