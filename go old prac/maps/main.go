package main

import (
	"fmt"
)

func main() {
	//type3
	//colors := make(map[int]string)

	//type2
	//var colors map[string]string

	//type1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//colors[10] = "#ffffff"
	//	fmt.Println(colors[11]) no output
	//delete(colors, 10)
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
