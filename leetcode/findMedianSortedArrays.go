package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	arg1 string
	arg2 string
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	flag.StringVar(&arg1, "list1", "1", "list1")
	flag.StringVar(&arg2, "list2", "1", "list2")
	flag.Parse()
	array1 := strings.Split(arg1, ",")
	array2 := strings.Split(arg2, ",")
	nums1 := make([]int, 0, len(array1))
	for _, v := range array1 {
		vi, _ := strconv.Atoi(v)
		nums1 = append(nums1, vi)
	}
	nums2 := make([]int, 0, len(array2))
	for _, v := range array2 {
		vi, _ := strconv.Atoi(v)
		nums2 = append(nums2, vi)
	}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				total = append(total, nums1[i])
				i++
			} else {
				total = append(total, nums2[j])
				j++
			}
		} else if i < len(nums1) {
			total = append(total, nums1[i:]...)
			break
		} else if j < len(nums2) {
			total = append(total, nums2[j:]...)
			break
		}
	}
	end := len(total) - 1
	return float64(total[end/2]+total[end/2+end%2]) / 2
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
}

func findKth(nums1, nums2 []int, start1, end1, start2, end2, k int) int {
	len1, len2 := end1-start1+1, end2-start2+1
	if len1 > len2 {
		return findKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		if nums1[start1] < nums2[start2] {
			return nums1[start1]
		} else {
			return nums2[start2]
		}
	}
	i := start1 + k/2 - 1
	if i >= end1 {
		i = end1 - 1
	}
	j := start2 + k/2 - 1
	if j >= end2 {
		j = end2 - 1
	}
	if nums1[i] > nums2[j] {
		return findKth(nums1, nums2, start1, end1, j+1, k-(j-start2+1))
	} else {
		return findKth(nums1, nums2, i+1, end1, start2, k-(i-start1+1))
	}
}
