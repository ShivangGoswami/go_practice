package main

import "fmt"

//Flags type
type Flags uint

const (
	//FlagUp var
	FlagUp Flags = 1 << iota
	//FlagBroadcast var
	FlagBroadcast
	//FlagLoopback var
	FlagLoopback
	//FlagPointToPoint var
	FlagPointToPoint
	//FlagMulticast var
	FlagMulticast
)

//IsUp comment
func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

//TurnDown comment
func TurnDown(v *Flags) { *v &^= FlagUp }

//SetBroadcast comment
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }

//IsCast comment
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main7() {
	//var v Flags = FlagMulticast | FlagUp
	var v = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
}
