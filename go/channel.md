CSP (communicating sequential processes)
生产者 消费者模型

```
func main() {
	msg := make(chan int, 10)
	wgC := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wgC.Add(1)
		go func() {
			defer wgC.Done()
			consumer(msg)
		}()
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			producer(i, msg)
		}(i)
	}
	wg.Wait()
	close(msg)
	fmt.Println("closed")
	wgC.Wait()
	fmt.Println("finish")
}

func producer(i int, msg chan int) {
	msg <- i
}

func consumer(msg chan int) {
	for id := range msg {
		fmt.Printf("recive %d\n", id)
	}
	fmt.Println("consumer exist")
}
```
chan 结构
```
type hchan struct {
    qcount   uint           // total data in the queue
    dataqsiz uint           // size of the circular queue
    buf      unsafe.Pointer // points to an array of dataqsiz elements
    elemsize uint16
    closed   uint32
    elemtype *_type // element type
    sendx    uint   // send index
    recvx    uint   // receive index
    recvq    waitq  // list of recv waiters
    sendq    waitq  // list of send waiters

    // lock protects all fields in hchan, as well as several
    // fields in sudogs blocked on this channel.
    //
    // Do not change another G's status while holding this lock
    // (in particular, do not ready a G), as this can deadlock
    // with stack shrinking.
    lock mutex
}
```
buf 环形队列，存储发送的数据，如果有缓存的话
G1向chan发送数据时，会先对buf加锁，然后将要发送的数据copy到buf里，并增加sendx的值，最后释放buf的锁
G2读取时，先对buf加锁，然后将数据copy到变量对应的内存里，增加recvx，最后释放锁。
写阻塞：
	G1向buf已满的ch发送数据时，当runtime检测到buf满了会通知调度器，
	将G1置为waiting状态，从p的运行队列移除，然后从p的运行队列中选择一个新的继续运行，这个过程不涉及到线程切换。
	G1放到sendq里，保存待发送数据变量的地址
	G2读出一个数据，会通知调度器把G1的状态置为runnable，然后将p加到runqueue中，等待调度。

读阻塞：
	G2读取buf为空的ch时，当runtime检测到buf满了会通知调度器，
	将G2置为waiting状态，从p的运行队列移除，然后从p的运行队列中选择一个新的继续运行，这个过程不涉及到线程切换。
	G2放到recvq里，保存待发送数据变量的地址
	G2读出一个数据，会通知调度器把G2的状态置为runnable，然后将p加到runqueue中，等待调度。
	
golang的channel是个结构体，里面大概包含了三大部分：
a. 指向内容的环形缓存区，及其相关游标
b. 读取和写入的排队goroutine链表
c. 锁
任何操作前都需要获得锁， 当写满或者读空的时候，就将当前goroutine加入到recvq或者sendq中， 并出让cpu(gopark)。

