package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

//Fetch downloads the URL and returns the name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	//Close file,but prefer error from Copy,if any
	defer func(f *os.File) {
		if closeErr := f.Close(); err == nil {
			fmt.Println(4)
			err = closeErr
		}
	}(f)
	return local, n, err
}

func main() {
	fmt.Println(fetch("https://github.com/adonovan/gopl.io/"))
}
