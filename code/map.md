map
可以对未初始化的map进行取值，但取出来的东西是空：

var m1 map[string]string
fmt.Println(m1["1"])
不能对未初始化的map进行赋值，这样将会抛出一个异常：panic: assignment to entry in nil map

var m1 map[string]string
m1["1"] = "1"
通过fmt打印map时，空map和nil map结果是一样的，都为map[]。所以，这个时候别断定map是空还是nil，而应该通过map == nil来判断。

不能对map取址。