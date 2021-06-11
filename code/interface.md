package main

import "fmt"

type upf interface {
	test()
}

var upfvar upf

type cupf struct {
}

type wupf struct {
}

func main() {
	{
		upfvar := cupf{}
		upfvar.test()
	}
	{
		upfvar := wupf{}
		upfvar.test()
	}
}

func (h *cupf) test() {
	fmt.Println("cupf test")
}

func (h *wupf) test() {
	fmt.Println("wupf test")
}
