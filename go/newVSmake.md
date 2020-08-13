#new

```
func new(Type) *Type
```
The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.
仅分配空间，传给new一个类型，会分配一个指针并返回，指针指向该类型的零值，
```
	test := new(map[string]string)
	(*test)["test"]="abc"
	// nil map panic
```

```
	var p1 *int
	i:=0
	p1 = &i
```
```
	p1 = new(int)
```
二者等价



# make
```
func make(Type, size IntegerType) Type 
```
分配空间+初始化
返回值是一个类型，而不是指针，会初始化
make只能用于 map slice channel

