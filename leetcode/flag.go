package main

import (
	"flag"
	"fmt"
)

var one int

func main() {
	flag.IntVar(&one, "o", 1, "arg one")
	flag.Parse()
	tail := flag.Args()
	fmt.Printf("Tail: %+v\n", tail)
}
