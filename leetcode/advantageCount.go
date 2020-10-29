package main

import (
	"fmt"
	"sort"
)

//
// sort A
// indexB sort
// B最小的
// 升序
// A用完，剩下的随便放
type valindex struct {
	value int
	index int
}

type valindexs []valindex

func (v valindexs) Len() int {
	return len(v)
}

func (v valindexs) Less() int {
	return len(v)
}

func advantageCount(A []int, B []int) []int {
	bValInd := make(valindexs, 0, len(B))
	for k, v := range B {
		bValInd = append(bValInd, valindex{value: v, index: k})
	}
	fmt.Println(A)
	sort.Ints(A)
	fmt.Println(A)
	sort.Sort(bValInd)
	return nil

}
func main() {
	A := []int{3, 5, 1, 5}
	advantageCount(A, nil)
}
