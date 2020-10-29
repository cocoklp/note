package main

import (
	"fmt"
)

var regular = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var special = map[string]map[string]int{
	"I": map[string]int{
		"IV": 4,
		"IX": 9,
	},
	"X": map[string]int{
		"XL": 50,
		"XC": 90,
	},
	"C": map[string]int{
		"CD": 400,
		"CM": 900,
	},
}

func romanToInt(s string) int {
	var result int
	for i := 0; i < len(s)-1; i++ {
		v := string[s[i]]
		if _, f := special[v]; f {
			if i < len(s)-1 {
				if tmp, f := special[v][string(s[i:i+2])]; f {
					result += tmp
					i++
					continue
				}
			}
		}
		result += regular[v]
	}
	return result
}
