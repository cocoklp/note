package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(squares()) // "1"
	fmt.Println(squares()) // "4"
	fmt.Println(squares()) // "9"
	fmt.Println(squares()) // "16"
	fmt.Println(f())       // "1"
	fmt.Println(f())       // "4"
	fmt.Println(f())       // "9"
	fmt.Println(f())       // "16"
	ff := squares()
	fmt.Println(squares()) // "1"
	fmt.Println(squares()) // "4"
	fmt.Println(squares()) // "9"
	fmt.Println(squares()) // "16"
	fmt.Println(ff())      // "1"
	fmt.Println(ff())      // "4"
	fmt.Println(ff())      // "9"
	fmt.Println(ff())      // "16"
}
