package main

import "fmt"

func main4() {
	// o := 0666
	// fmt.Printf("%d %[1]o %#[1]o\n", o)
	// x := int64(0xdeadbeef)
	// fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Print(string(1234567))
}
