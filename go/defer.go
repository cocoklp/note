/************
defer func
	defer 后面的函数会被延迟执行，直到包含 改defer语句的函数执行完毕时，defer后的函数才会被执行。
	无论是正常return还是panic，都会执行defer
	多条refer语句，执行顺序和声明顺序相反。
函数返回过程
	return 是非原子操作
	先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中

os.Exist() 不执行defer
*************/

/*记录进入函数和退出函数的时间*/
func bigSlowOperation() {
    defer trace("bigSlowOperation")() // don't forget the extra parentheses不要忘记defer语句后的圆括号，否则本该在进入时执行的操作会在退出时执行，而本该在退出时执行的，永远不会被执行
    // ...lots of work…
    time.Sleep(10 * time.Second) // simulate slow
    operation by sleeping
}
func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
    return func() { 
        log.Printf("exit %s (%s)", msg,time.Since(start)) 
    }
}

/*
	defer被声明时，参数就会被实时解析
*/
func a() {
	i := 0
	defer fmt.Println(i)  // 相当于defer fmt.Println(0)
	i++
	return
}

type SliceNum []int

func NewSlice() SliceNum {
    return make(SliceNum, 0)

}

func (s *SliceNum) Add(elem int) *SliceNum {
    *s = append(*s, elem)
    fmt.Println("add", elem)
    fmt.Println("add SliceNum end", s)
    return s
}
func (s *SliceNum) test(elem int) *SliceNum {
    fmt.Println("test", elem)
    fmt.Println("test", s)
    return s
}

/*
	返回 1 11 12 13 4
	defer 延迟执行的是最后一个
*/
func deferFunc() {
    s := NewSlice()
    defer s.Add(1).Add(4)
    s.Add(11)
    s.Add(12)
    s.Add(13)
}

/*
	11 12 13 1 4
*/
func deferFunc() {
    s := NewSlice()
    defer func(){
        s.Add(1).Add(4)
    }()
    s.Add(11)
    s.Add(12)
    s.Add(13)
}

/*
	返回 2
	result = i
	result ++
	return result
*/
func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}

/*
	返回1
	直接把1入栈作为返回值，延迟函数无法操作改返回值
*/
func deferFuncReturn() int {
	var i int
	defer func() {
		i++
	}()
	return 1
}

/*
	返回0
	新变量存储 newret = i
	i++
	返回newret
	defer可以操作局部变量i，但是不能影响返回值
*/
func deferFuncReturn() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

/*
	返回1
*/
func deferFuncReturn() (result int) {
	defer func() {
		result++
	}()
	return 0
}
