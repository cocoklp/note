package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}

func min(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	width := right - left
	area := width * min(height[left], height[right])
	i := 1
	for left < right {
		if i > 100 {
			return area
		}
		fmt.Println(left, right)
		if height[left] < height[right] {
			temp := left + 1
			for height[temp] <= height[left] && temp < right {
				temp++
			}
			if temp >= right {
				break
			}
			width := right - temp
			tmpArea := width * min(height[temp], height[right])
			if tmpArea > area {
				area = tmpArea
				left = temp
			}
		} else {
			temp := right - 1
			for height[temp] <= height[left] {
				temp--
			}
			if temp <= left {
				break
			}
			width := temp - left
			tmpArea := width * min(height[temp], height[left])
			if tmpArea > area {
				area = tmpArea
				right = temp
			}

		}
	}
	return area
}
