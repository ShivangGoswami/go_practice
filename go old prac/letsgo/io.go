package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func write() {
	str := "i have a sample text"
	err := ioutil.WriteFile("D:\\sam.txt", []byte(str), os.FileMode(777))
	fmt.Println(err)
}

func read() {
	bs, err := ioutil.ReadFile("sam.txt")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(2)
	}
	fmt.Println(string(bs))
}

func shuffle(a []int) {
	for i := range a {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		np := r.Intn(len(a) - 1)
		a[i], a[np] = a[np], a[i]
	}
}
func main() {
	write()
	//read()

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// shuffle(a)
	// fmt.Println(a)
}
