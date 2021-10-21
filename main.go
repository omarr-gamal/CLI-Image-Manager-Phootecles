package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: imgrecognition <img_url>")
	}
	fmt.Printf("url: %s", os.Args[1])
}
