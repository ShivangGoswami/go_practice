package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create new type
type mytp []string

//receiver function
func (m mytp) printSlice() {
	fmt.Println(m)
}

//new function
func newtype() mytp {
	var a mytp
	b := []string{"a", "b", "c", "d"}
	c := []string{"e", "d", "f", "g"}
	for _, d := range b {
		for _, e := range c {
			a = append(a, d+e)
		}
	}
	return a
}

func (m mytp) splitter() (mytp, mytp) {
	return m[0:2], m[2:len(m)]
	//return m[:2],m[2:]
}

func (m mytp) toString() string {
	return strings.Join([]string(m), ",")
}

func (m mytp) write() {
	error := ioutil.WriteFile("sam.txt", []byte(m.toString()), os.FileMode(777))
	fmt.Println(error)
}

func read() {
	byteSlice, error := ioutil.ReadFile("sam.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(2)
	}
	//fmt.Println(byteSlice)
	x := strings.Split(string(byteSlice), ",")
	y := mytp(x)
	fmt.Println(y)
}

func (m mytp) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range m {
		np := r.Intn(len(m) - 1)
		m[i], m[np] = m[np], m[i]
	}
}
func main() {
	//var a mytp
	//a = append(a, "q", "w")
	a := newtype()
	//a := mytp{"q", "w"}
	//a.printSlice()
	//fmt.Println(a.splitter())
	//read()
	a.shuffle()
	a.printSlice()
}
