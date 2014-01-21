package main

import (
	"fmt"
	"github.com/GeertJohan/go.qrt"
	"log"
)

func main() {
	str, err := qrt.Generate("some example text")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(str)
}
