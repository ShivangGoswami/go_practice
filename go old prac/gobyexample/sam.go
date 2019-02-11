package main

import (
	"fmt"
	"strconv"
)

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a string, b, c int) string {
	return a + string(b) + strconv.Itoa(c)
}

func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res1 := plusPlus("gaurav", 50, 3)
	fmt.Println("1+2+3 =", res1)
}
