package main

import (
	"flag"
	"fmt"
)

var str string

func main() {
	flag.StringVar(&str, "string", "abc", "string")
	flag.Parse()
	fmt.Println(lengthOfLongestSubstring(str))
}

/*
	abcdcdaad
	使用一个map记录各个元素位置，当出现重复时，start改为重复元素的下一个
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	charMap := make(map[string]int)
	start, end := 0, 0
	maxLen := 0

	for k, vr := range s {
		v := string(vr)
		index, f := charMap[v]
		end = k
		charMap[v] = k

		if f && index >= start {
			if end-start > maxLen {
				maxLen = end - start
			}
			start = index + 1
		}
	}

	if end-start+1 > maxLen {
		maxLen = end - start + 1
	}

	return maxLen
}
