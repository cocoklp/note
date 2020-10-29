package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	last := 0
	for n > 1 {
		tmp := n >> 1
		last = n - tmp<<1
		n = tmp
		if last != 0 {
			return false
		}
	}
	fmt.Println(last)
	return last == 0
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println(isPowerOfTwo(n))
}
