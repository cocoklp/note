package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	index map[int]int
	data  []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		index: make(map[int]int),
		data:  make([]int, 0, 100),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, f := this.index[val]; f {
		return false
	}
	this.data = append(this.data, val)
	this.index[val] = len(this.data) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	k, f := this.index[val]
	if !f {
		return false
	}
	fmt.Println(this.data, this.index, val, k)
	lastVal := this.data[len(this.data)-1]
	this.data[k] = lastVal
	this.data = this.data[:len(this.data)-1]
	this.index[lastVal] = k
	delete(this.index, val)
	fmt.Println(this.data, this.index, val, k)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.data) <= 0 {
		return 0
	}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	k := rd.Intn(len(this.data))
	return this.data[k]
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Remove(3))
	/*
		fmt.Println(obj.Insert(3))
		fmt.Println(obj.Insert(4))
		fmt.Println(obj.Insert(2))
		fmt.Println(obj.Remove(2))
		fmt.Println(obj.Insert(1))
		fmt.Println(obj.Insert(-3))
		fmt.Println(obj.Insert(-2))
		fmt.Println(obj.Remove(-2))
		fmt.Println(obj.Remove(3))
		fmt.Println(obj.Insert(-1))
		fmt.Println(obj.Remove(-3))
		fmt.Println(obj.Insert(1))
		fmt.Println(obj.Insert(-2))
		fmt.Println(obj.Insert(-2))
		fmt.Println(obj.Insert(-2))
		fmt.Println(obj.Insert(1))
		fmt.Println(obj.Insert(-2))
		fmt.Println(obj.Remove(3))
		fmt.Println(obj.Insert(-3))
		fmt.Println(obj.Insert(1))
	*/
	fmt.Println(obj)
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
  * obj := Constructor();
   * param_1 := obj.Insert(val);
    * param_2 := obj.Remove(val);
	 * param_3 := obj.GetRandom();
*/
