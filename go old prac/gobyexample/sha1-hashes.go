package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()
	md := md5.New()
	h.Write([]byte(s))
	md.Write([]byte(s))

	bs := h.Sum(nil)
	bs1 := md.Sum(nil)

	fmt.Println(s)
	fmt.Printf("SHA1 hash:%x\n", bs)
	fmt.Printf("MD5 Digest:%x\n", bs1)
}
