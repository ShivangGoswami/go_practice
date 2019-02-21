package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq() {
	words := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		words[in.Text()]++
	}
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
	for key, val := range words {
		fmt.Println(key, ":", val)
	}
}

func main() {
	wordfreq()
}
