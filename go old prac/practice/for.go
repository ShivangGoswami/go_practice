package main

import (
	"fmt"
)

func hi() {
	i := 0
	//while loop example
	for i < 2 {
		fmt.Println(i)
		i++
	}
	//primitive for loop
	for i = 0; i < 2; i++ {
		fmt.Println(i)
	}
	//do while loop in go
	i = 0
	for {
		fmt.Println(i)
		if i == 2 {
			break
		}
		i++
	}
	// for in range
	arr := []int{0, 1, 2, 3, 4}
	for _, q := range arr {
		fmt.Println(q)
	}
}
