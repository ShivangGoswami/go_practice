package main

import (
	"fmt"
	"os"
	"strings"
)

func main4() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
