package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./unitlib"
)

func checkerr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exer:%v\n", err)
		os.Exit(1)
	}
}

func main5() {
	var inp []string
	if len(os.Args) == 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			inp = append(inp, input.Text())
		}
		checkerr(input.Err())
	} else {
		inp = os.Args[1:]
	}
	for _, arg := range inp {
		t, err := strconv.ParseFloat(arg, 64)
		checkerr(err)
		f := unitlib.Fahrenheit(t)
		c := unitlib.Celsius(t)
		ft := unitlib.Feet(t)
		mt := unitlib.Meters(t)
		p := unitlib.Pound(t)
		kg := unitlib.Kilograms(t)
		fmt.Printf("%s=%s,%s=%s\n", f, unitlib.FToC(f), c, unitlib.CToF(c))
		fmt.Printf("%s=%s,%s=%s\n", ft, unitlib.FToM(ft), mt, unitlib.MToF(mt))
		fmt.Printf("%s=%s,%s=%s\n", p, unitlib.PoundToKg(p), kg, unitlib.KgToPound(kg))
	}
}
