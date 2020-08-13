package main

import "fmt"

func bubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		needBreak := true
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				needBreak = false
			}
		}
		if needBreak {
			break
		}
	}
}

func selectSort(data []int) {
	if len(data) <= 1 {
		return
	}
	min := data[0]
	index := 0
	for i, v := range data {
		if min > v {
			min = v
			index = i
		}
	}
	data[0], data[index] = data[index], data[0]
	selectSort(data[1:])
}

func insertSort(data []int) {
	for i := 1; i < len(data); i++ {
		preIndex := i - 1
		cur := data[i]
		for preIndex >= 0 && data[preIndex] > cur {
			data[preIndex+1] = data[preIndex]
			preIndex--
		}
		data[preIndex+1] = cur
	}
}

func shellSort(data []int) {
	for gap := len(data) / 2; gap > 0; gap = gap >> 1 {
		for i := 0; i < len(data)-gap; i++ {
			if data[i] > data[i+gap] {
				data[i], data[i+gap] = data[i+gap], data[i]
			}
		}
	}
}

// 递归
func mergeSort1(data []int) {
	fmt.Println(data)
	if len(data) < 2 {
		return
	}
	length := len(data)
	left := data[:length/2]
	right := data[length/2:]
	mergeSort1(left)
	mergeSort1(right)
	data = merge1(left, right)

}
func merge1(left, right []int) []int {
	fmt.Println(left, right)
	lindex, rindex := 0, 0
	rst := make([]int, 0, len(left)+len(right))
	for lindex < len(left) && rindex < len(right) {
		if left[lindex] < right[rindex] {
			rst = append(rst, left[lindex])
			lindex++
		} else {
			rst = append(rst, right[rindex])
			rindex++
		}
	}
	rst = append(rst, left[:lindex]...)
	rst = append(rst, right[:rindex]...)
	return rst

}

func mergeSort(data []int, first, end int) {
	if first < end {
		mid := (first + end) / 2
		mergeSort(data, first, mid)
		mergeSort(data, mid+1, end)
		merge(data, first, mid, end)
	}
}

func merge(data []int, first, mid, end int) {
	l1 := make([]int, mid-first+1)
	l2 := make([]int, end-mid)
	copy(l1, data[first:mid+1])
	copy(l2, data[mid+1:end+1])
	fmt.Println(l1, l2, data, first, mid, end)
	i, j, k := 0, 0, first
	for i < mid-first+1 && j < end-mid {
		if l1[i] < l2[j] {
			data[k] = l1[i]
			i++
		} else {
			data[k] = l2[j]
			j++
		}
		k++
	}
	if i < len(l1) {
		for ij := i; ij < len(l1); ij++ {
			data[k] = l1[ij]
			k++
		}
	}
	if j < len(l2) {
		for ij := j; ij < len(l2); ij++ {
			data[k] = l2[ij]
			k++
		}
	}
}

func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	privot := data[0]
	less := make([]int, 0, len(data))
	greater := make([]int, 0, len(data))
	for i := 1; i < len(data); i++ {
		if data[i] < privot {
			less = append(less, data[i])
		} else {
			greater = append(greater, data[i])
		}
	}
	left := quickSort(less)
	right := quickSort(greater)
	res := append(left, privot)
	res = append(res, right...)
	return res
}

func main() {
	fmt.Println("vim-go")
	data := []int{2, 1, 3, 9, 3, 5, 8}
	input := make([]int, len(data))
	copy(input, data)
	bubbleSort(input)
	fmt.Println(input, data)
	copy(input, data)
	selectSort(input)
	fmt.Println(input, data)
	copy(input, data)
	insertSort(input)
	fmt.Println(input, data)
	copy(input, data)
	shellSort(input)
	fmt.Println(input, data)
	copy(input, data)
	mergeSort(input, 0, len(input)-1)
	fmt.Println(input, data)
	copy(input, data)
	mergeSort1(input)
	fmt.Println(input, data)
	copy(input, data)
	quickSort(input)
	fmt.Println(input, data)
}
