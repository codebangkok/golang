package main

import (
	"compress/zlib"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("blob")
	if err != nil {
		log.Fatal(err)
	}
	r, err := zlib.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, r)
	r.Close()
}
