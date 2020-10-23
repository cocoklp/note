```
func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
```
runtime.Caller
	func Caller(skip int) (pc uintptr, file string, line int, ok bool)
	Caller可以返回函数调用栈的某一层的程序计数器、文件信息、行号。
	0 代表当前函数，也是调用runtime.Caller的函数。1 代表上一层调用者，以此类推。

runtime.Callers
	func Callers(skip int, pc []uintptr) int
	Callers用来返回调用站的程序计数器, 放到一个uintptr中。
	0 代表 Callers 本身，这和上面的Caller的参数的意义不一样，历史原因造成的。 1 才对应这上面的 0。

	比如下面代码可以打印出整个的调用栈。
	```
	func trace() {
		pc := make([]uintptr, 10) // at least 1 entry needed
		n := runtime.Callers(0, pc)
		for i := 0; i < n; i++ {
			f := runtime.FuncForPC(pc[i])
			file, line := f.FileLine(pc[i])
			fmt.Printf("%s:%d %s\n", file, line, f.Name())
		}
	}
	```

runtime.CallersFrames
```
	func CallersFrames(callers []uintptr) *Frames
```
	Callers只是获取栈的程序计数器，如果想获得整个栈的信息，可以使用CallersFrames函数，省去遍历调用FuncForPC。
	```
	func trace2() {
		pc := make([]uintptr, 10) // at least 1 entry needed
		n := runtime.Callers(0, pc)
		frames := runtime.CallersFrames(pc[:n])
		for {
			frame, more := frames.Next()
			fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
			if !more {
				break
			}
		}
	}
	```








