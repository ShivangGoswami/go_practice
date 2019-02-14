package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

func main() {
	hash := flag.String("hash", "sha256", "sha256 or sha384 or sha512")
	flag.Parse()
	if strings.EqualFold(*hash, "sha256") {
		fmt.Printf("%x", sha256.Sum256([]byte(strings.Join(flag.Args(), ""))))
	} else if strings.EqualFold(*hash, "sha384") {
		fmt.Printf("%x", sha512.Sum384([]byte(strings.Join(flag.Args(), ""))))
	} else if strings.EqualFold(*hash, "sha512") {
		fmt.Printf("%x", sha512.Sum512([]byte(strings.Join(flag.Args(), ""))))
	}
}
