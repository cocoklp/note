package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))

}

var letMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	} else if len(digits) == 1 {
		return letMap[digits]
	}

	befStr := letterCombinations(string(digits[1:]))
	curLets := letMap[string(digits[0])]
	res := make([]string, 0, len(curLets)*len(befStr))
	for _, v := range curLets {
		for _, b := range befStr {
			res = append(res, v+b)
		}
	}
	return res
}
