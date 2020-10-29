package main

import (
	"fmt"
	"sort"
	"strconv"
)

type myInt []int

func (m myInt) Len() int {
	return len(m)
}
func (m myInt) Less(i, j int) bool {
	datai := int(m[i])
	dataj := int(m[j])
	is := []rune(strconv.Itoa(datai))
	js := []rune(strconv.Itoa(dataj))
	di := is[0]
	dj := js[0]
	for k := 0; k < len(is) || k < len(js); k++ {
		if k < len(is) {
			di = is[k]
		}
		if k < len(js) {
			dj = js[k]
		}
		if di < dj {
			return true
		} else if di > dj {
			return false
		}
	}
	return false
}

func (m myInt) Swap(i, j int) {
	fmt.Println(i, j)
	m[i], m[j] = m[j], m[i]
}

func combine(nums []int) int {
	data := make(myInt, 0, len(nums))
	for _, n := range nums {
		data = append(data, n)
	}
	fmt.Println(data)
	sort.Sort(data)
	fmt.Println(data)
	str := ""
	for i := len(data) - 1; i >= 0; i-- {
		str += strconv.Itoa(int(data[i]))
	}
	v, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return v
}
func test(nums [3]int) {
	fmt.Printf("%p %p\n", nums, &nums)
	nums[0] = 1111
	//numss[0], numss[1] = numss[1], numss[0]
}
func test2(nums []int) {
	fmt.Printf("%p %p\n", nums, &nums)
	nums[0] = 1111
	//numss[0], numss[1] = numss[1], numss[0]
}
func main() {
	nums := []int{34, 3, 4}
	fmt.Println(combine(nums))
	nums = []int{456, 12, 342, 78}
	fmt.Println(combine(nums))

	/*
		for i := 0; i < len(nums); i++ {
			numss = append(numss, nums[i])
		}
	*/
	array := [3]int{1, 2, 3}
	fmt.Printf("%v %p \n", array, &array)
	test(array)
	fmt.Printf("%v %p \n", array, &array)
	slice := []int{1, 2, 3}
	fmt.Printf("%v %p \n", slice, &slice)
	test2(slice)
	fmt.Printf("%v %p \n", slice, &slice)

	arrayA := [...]int{1, 2, 3, 4, 5, 5, 6}
	fmt.Println(len(arrayA))

	{
		darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
		dslice := darr[2:5]
		fmt.Println("array before", darr)
		for i := range dslice {
			dslice[i]++
		}
		dslice = append(dslice, 666)
		fmt.Println("array after", darr, dslice)
	}

}
