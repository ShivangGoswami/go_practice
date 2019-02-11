package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

func commabuffer(s string) string {
	var buf bytes.Buffer
	n := 0
	for n < len(s)-3 {
		buf.WriteString(s[n:n+3] + ",")
		n = n + 3
	}
	buf.WriteString(s[n:])
	return buf.String()
}

func commabase(s string) string {
	sign := ""
	if s[:1] == "+" || s[:1] == "-" {
		sign = s[:1]
		s = s[1:]
	}
	if a := strings.LastIndex(s, "."); a != -1 {
		if strings.Count(s, ".") > 1 {
			log.Println("INVALID NUMBER:", s)
		} else if a == len(s)-1 {
			s = s + "0"
		}
		return sign + commabuffer(s[:a]) + s[a:]
	}
	return sign + commabuffer(s)
}

func analgram(s1 string, s2 string) bool {
	sachi := false
	if len(s1) != len(s2) {
		return false
	}
	for _, ele := range s1 {
		if strings.ContainsRune(s2, ele) && strings.Count(s1, string(ele)) == strings.Count(s2, string(ele)) {
			sachi = true
		} else {
			sachi = false
			break
		}
	}
	return sachi
}

func main6() {
	// fmt.Println(analgram("sachi", "achis"))
	// fmt.Println(analgram("chi", "achis"))
	// fmt.Println(analgram("chiquee", "chi"))
	// fmt.Println(commabuffer("12345678901"))
	// fmt.Println(commabuffer("1234567890"))
	// fmt.Println(commabuffer("123456789"))
	// fmt.Println(commabuffer("123456789012"))
	fmt.Println(commabase("-123456.789"))
	fmt.Println(commabase("-12345.789"))
	fmt.Println(commabase("-1234.789"))
	fmt.Println(commabase("-123456.78"))
	fmt.Println(commabase("-123456.7"))
	fmt.Println(commabase("+123456."))
	fmt.Println(commabase("-123456.789."))
}
