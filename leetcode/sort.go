package main

import (
	"fmt"
	"sort"
)

func minHeap(root int, end int, c []int) {
	for {
		var child = 2*root + 1
		//判断是否存在child节点
		if child > end {
			break
		}
		//判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}

func minHeap1(data []int) {
	if len(data) < 2 {
		return
	}
	//root := (len(data) - 1) / 2
	root := 0
	for {
		child := root*2 + 1
		if child >= len(data) {
			break
		}
		if child+1 < len(data) && data[child] > data[child+1] {
			child++
		}
		if data[child] < data[root] {
			data[child], data[root] = data[root], data[child]
			root = child
		} else {
			break
		}

	}
	return
}

//降序排序
func HeapSort(c []int) {
	var n = len(c) - 1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, c)
	}
	//fmt.Println("堆构建完成", c)
	for end := n; end >= 0; end-- {
		if c[0] < c[end] {
			c[0], c[end] = c[end], c[0]
			minHeap(0, end-1, c)
		}
	}
}

func HeapSort1(data []int) []int {
	n := len(data) - 1
	for i := n / 2; i >= 0; i-- {
		minHeap1(data[i:])
	}
	//fmt.Println("my heap", data)
	for end := n; end > 0; end-- {
		if data[0] < data[end] {
			data[0], data[end] = data[end], data[0]
			minHeap1(data[:end])
		}
	}
	return data
}

func CountSort(data []int) {
	if len(data) < 1 {
		return
	}
	max, min := data[0], data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	countArray := make([]int, max-min+1)
	for _, v := range data {
		countArray[v-min]++
	}
	fmt.Println(countArray, min, max)
	index := 0
	for k := 0; k < len(countArray); k++ {
		for countArray[k] != 0 {
			data[index] = k + min
			countArray[k]--
			index++
		}
	}
}

func main() {
	data := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	sort.Ints(data)
	data = []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	HeapSort(data)
	fmt.Println(data)
	data = []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	data = HeapSort1(data)
	fmt.Println(data)

	data = []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	CountSort(data)
	fmt.Println(data)
	data = []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	BucketSort(data)
	fmt.Println(data)
	data = []int{72, 11, 82, 32, 44, 13, 17, 95, 54, 28, 79, 56}
	RadixSort(data)
	fmt.Println(data)
}

func BucketSort(arr []int) {
	var bucket [][]int
	for i := 0; i < 20; i++ {
		tmp := make([]int, 1)
		bucket = append(bucket, tmp)
	}
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]] = append(bucket[arr[i]], arr[i])
	}
	fmt.Println(bucket)

	index := 0
	for i := 0; i < 20; i++ {
		if len(bucket[i]) > 1 {
			for j := 1; j < len(bucket[i]); j++ {
				arr[index] = bucket[i][j]
				index++
			}
		}
	}

}

func RadixSort(data []int) []int {
	radix := 10
	for {
		bucket := make([][]int, 10)
		needBreak := true
		for _, v := range data {
			tmp := (v - v/radix*radix) / (radix / 10)
			//fmt.Println(radix, tmp, v)
			if tmp != 0 {
				needBreak = false
			}
			if bucket[tmp] == nil {
				bucket[tmp] = make([]int, 0, len(data))
			}
			bucket[tmp] = append(bucket[tmp], v)
		}
		fmt.Println(bucket)
		if needBreak {
			break
		}
		radix = radix * 10
		index := 0
		for _, vs := range bucket {
			for _, v := range vs {
				data[index] = v
				index++
			}
		}
		fmt.Println(data)
	}
	return data
}
