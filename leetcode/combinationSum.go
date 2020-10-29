package main

import (
	"fmt"
)

var res [][]int
var res1 []int

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
func combinationSum(candidates []int, target int) [][]int {
	submission(candidates, target, len(candidates)-1, res1)
	return res
}

func submission(candidates []int, target int, index int, res1 []int) {
	if target == 0 {
		res = append(res, res1)
		return
	}
	for index >= 0 {
		if candidates[index] <= target {
			res1 = append(res1, candidates[index])
			submission(candidates, target-candidates[index], index, res1)
			res1 = res1[:len(res1)-1]
		}
		index--
	}
	return
}
