package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type happy []string

func new() happy {
	a := happy{"q", "w", "e", "r", "t", "y"}
	return a
}
func (a happy) toString() string {
	return strings.Join([]string(a), ",")
}
func (a happy) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range a {
		np := r.Intn(len(a) - 1)
		a[np], a[i] = a[i], a[np]
	}
}
func main() {
	a := new()
	// ioutil.WriteFile("sam.txt", []byte(a.toString()), os.FileMode(777))
	// byteSlice, err := ioutil.ReadFile("sam.txt")
	// if err != nil {
	// 	os.Exit(2)
	// }
	// x := strings.Split(string(byteSlice), ",")
	// y := happy(x)
	// fmt.Println(y)
	a.shuffle()
	fmt.Println(a)
}
