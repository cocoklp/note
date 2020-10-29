package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))
	findMap := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			findMap[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	rst := make([]int, len(nums1))
	for k, v := range nums1 {
		if val, ok := findMap[v]; ok {
			rst[k] = val
		} else {
			rst[k] = -1
		}
	}
	fmt.Println(rst)
	return rst
}

func nextGreaterElements1(nums []int) []int {
	fmt.Println(nums)
	rst := make([]int, len(nums))
	for i := 0; i <= len(nums)-1; i++ {
		for l := 1; l <= len(nums)-1; l++ {
			tmp := i + l
			if tmp >= len(nums) {
				tmp = tmp - len(nums)
			}
			rst[i] = -1
			if nums[tmp] > nums[i] {
				rst[i] = nums[tmp]
				break
			}
		}
	}
	fmt.Println(rst)
	return rst
}

func nextGreaterElements(nums []int) []int {
	nums2 := make([]int, 0, 2*len(nums))
	nums2 = append(append(nums2, nums...), nums...)
	fmt.Println(nums, nums2)
	rst := make([]int, len(nums))
	stack := make([]int, 0, len(nums))
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums2[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if i < len(nums) {
			if len(stack) != 0 {
				rst[i] = stack[len(stack)-1]
			} else {
				rst[i] = -1
			}
		}
		stack = append(stack, nums2[i])
	}
	fmt.Println(rst)
	return rst
}

func nextGreaterElement32(n int) int {
	nums := make([]int, 0, 10)
	for n != 0 {
		tmp := n % 10
		n = n / 10
		nums = append([]int{tmp}, nums...)
	}
	fmt.Println(nums)
	found := false
	stack := make([]int, 0, 10)
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			nums[i], nums[stack[len(stack)-1]] = nums[stack[len(stack)-1]], nums[i]
			found = true
			break
		}
		stack = append(stack, i)
	}
	fmt.Println(nums)
	rst := 0
	if found {
		for _, v := range nums {
			rst = 10*rst + v
		}
		return rst
	}
	return -1

}

func main() {
	{
		nums1 := []int{4, 1, 2}
		nums2 := []int{1, 3, 4, 2}
		nextGreaterElement(nums1, nums2)
	}
	{
		nums1 := []int{2, 4}
		nums2 := []int{1, 2, 3, 4}
		nextGreaterElement(nums1, nums2)
	}
	{
		nums1 := []int{1, 3, 5, 2, 4}
		nums2 := []int{6, 5, 4, 3, 2, 1, 7}
		nextGreaterElement(nums1, nums2)
	}
	fmt.Println("=============")
	{
		nextGreaterElements([]int{1, 2, 1})
	}
	{
		nextGreaterElements([]int{41, 12, 3})
	}
	{
		fmt.Println(nextGreaterElement32(121))
	}
}
