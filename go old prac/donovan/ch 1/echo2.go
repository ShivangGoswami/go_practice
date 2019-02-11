package main

import (
	"fmt"
	"os"
)

func ex2() {
	for ind, arg := range os.Args[1:] {
		fmt.Println(ind, arg)
	}
}
func main3() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// func main() {
// 	ex2()
// }
