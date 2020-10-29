package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 4, 5}))
	TestArr1()
}

func TestArr1() {
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	var b = make([]int, 0, 10)
	copy(b, a)
	fmt.Println(b, a)
}

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	before := permute(nums[:len(nums)-1])
	fmt.Println(before)
	new := make([][]int, 0, len(before))
	cur := nums[len(nums)-1]

	for _, v := range before {
		fmt.Println(v)
		for k := 0; k < len(v); k++ {
			one := make([]int, 0, len(v)+1)
			tmp := make([]int, len(v[0:k]))
			copy(tmp, v[0:k])
			fmt.Println("tmp", tmp, v[:k])
			fmt.Println("one", one, tmp, v)
			one = append(one, tmp...)
			fmt.Println("one", one, tmp, v)
			one = append(one, cur)
			fmt.Println("one", one, tmp, v)
			one = append(one, v[k:]...)
			fmt.Println("one", one, tmp, v)
			new = append(new, one)
		}
		new = append(new, append(v, cur))
	}
	return new
}
