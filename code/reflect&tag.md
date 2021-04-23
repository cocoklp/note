# reflect

反射主要与interface类型相关，只有interface才有反射一说。变量包括type和value两部分，在golang中，每个interface都有一个pair记录了实际变量的值和类型。

反射是用来检测存储在接口变量内部pair对的一种机制。通过reflect.ValueOf() 和 reflect.TypeOf() 访问接口变量的内容。

## ValueOf() TypeOf()

获取参数接口中数据的值，如果接口为空返回0

获取参数接口中数据的类型，如果接口为空返回nil

### 已知原有类型

```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

运行结果：
0xc42000e238
1.2345
```



### 未知原有类型—遍历field

```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func main() {

	user := User{1, "Allen.Wu", 25}

	DoFiledAndMethod(user)

}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

运行结果：
get Type is : User
get all Fields is: {1 Allen.Wu 25}
Id: int = 1
Name: string = Allen.Wu
Age: int = 25
ReflectCallFunc: func(main.User)
```



## 设置变量的值

```
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)

	////////////////////
	// 如果reflect.ValueOf的参数不是指针，会如何？
	pointer = reflect.ValueOf(num)
	//newValue = pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}

运行结果：
old value of pointer: 1.2345
type of pointer: float64
settability of pointer: true
new value of pointer: 77
```



## 函数调用

```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

// 如何通过反射来进行方法的调用？
// 本来可以用u.ReflectCallFuncXXX直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调动mv.Call

func main() {
	user := User{1, "Allen.Wu", 25}
	
	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
	getValue := reflect.ValueOf(user)

	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法
	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
}
```

## 性能

反射涉及到GC，并且实现里有大量的枚举，性能不好。

https://studygolang.com/articles/12348?fr=sidebar

https://www.cnblogs.com/itbsl/p/10551880.html

# tag

tag由一个或多个键值对组成，键值使用冒号分隔，值用双引号括起来。键值对之间使用空格分隔

在结构体字段后书写格式为

`key1:"value1" key2:"value2"`

```
package main

import (
	"fmt"
	"reflect"
)

type ImageQuery struct {
	ProjectId string `harbor:"project_id"`
	PageNo    string `harbor:"page"`
	PageSize  string `harbor:"page_size"`
	NameQuery string `harbor:"q"`
	Sort      string //`harbor:"sort"`
}

func struct2Map(obj interface{}) (rst map[string]interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fmt.Println("type", t.Name())
	fmt.Println("val", v)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("harbor")
		fmt.Println(tag)
		if tag != "" {
			data[tag] = v.Field(i).Interface()
		}
	}
	return data
}

func main() {
	image := ImageQuery{
		ProjectId: "1",
	}
	struct2Map(image)
}

输出 
type ImageQuery
val {1    }
project_id
page
page_size
q


```

