package main

import (
	"fmt"
)

type mytp []string

func (m mytp) printSlice() {
	fmt.Println(m)
}

func newtype() mytp {
	var a mytp
	b := []string{"a", "b", "c", "d"}
	c := []string{"e", "f", "g", "h"}
	for _, x := range b {
		for _, y := range c {
			a = append(a, x+y)
		}
	}
	return a
}

func main() {
	a := newtype()
	fmt.Println(a)
	//var a mytp
	//a = append(a, "q", "w")
	//a:=mytp{"q","w"}
	//a.printSlice()
}
