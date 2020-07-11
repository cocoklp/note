/*
具体类型：
	可以直接操作或通过内置方法操作
接口类型：
	抽象类型，不会暴露所代表的对象的内部值的结构和这个对象支持的基础操作的集合，只会展示出自己的方法
	不知道是什么，只知道可以用来做什么

go接口时一组方法的集合，是一种抽象的类型，任何类型只要实现了接口中的方法集，就属于这个类型
空接口interface{}，没有定义任何方法，所以所有的类型都能满足它，所以当函数参数类型为interface{}时，可以给它传任意类型的参数
*/

package main

import (
	"fmt"
)

/*定义一个接口，有两个方法*/
type Duck interface {
	Quack()
	DuckGo()
}

/*定义一个具体类型，实现方法*/
type Chicken struct {
}

func (c Chicken) Quack() {
	fmt.Println("gaga")
}

func (c Chicken) DauckGo() {
	fmt.Println("gogo")
}

func DoDuck(d Duck) {
	d.DuckGo()
	d.Quack()
}

func main() {
	fmt.Println("hello")
	c := Chicken{}
	DoDuck(c) // 正常，因为 Chicken 实现了 Duck的所有方法
}
