package main

import (
	"fmt"
	"os"
)

func ex1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Print(s)
}
func main2() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Print(s)
}

// func main() {
// 	ex1()
// }
