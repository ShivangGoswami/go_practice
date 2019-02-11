package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//Exercise 1.4
func checkfiles(line string) string {
	ret := ""
	for _, file := range os.Args[1:] {
		data, _ := ioutil.ReadFile(file)
		if strings.Contains(string(data), line) {
			ret += filepath.Base(file) + " "
		}
	}
	return ret
}

func main7() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println("Exists in files", checkfiles(line))
		}
	}
}
