redigo第三方包

1. 多线程不安全
https://juejin.im/post/5d07b8e9f265da1b7f297f77?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
Connections support one concurrent caller to the Receive method and one concurrent caller to the Send and Flush methods. No other concurrency is supported including concurrent calls to the Do and Close methods.
连接支持同时运行单个执行体调用 Receive 和 单个执行体调用 Send 和 Flush 方法。不支持并发调用 Do 和 Close 方法。

Do 顺序执行数据的发送和响应的接收，在底层的writeCommand方法上，通过for range 将redis命令发送到redis-server，没有加锁，并发时很可能会导致数据交叉。最终调用到writeString写数据，会写入到bw中，bw是conn的一个net.Conn的writter，所以并发执行Do时，都会往一个net.Conn的writter里写数据，导致数据交叉。写完后，Do会调用bw的Flush方法，将缓存中的数据都发送出去，会判断已写长度是否和缓冲区的长度相等，并发时，会出现一个执行体flush过程中，有其他执行体向bw中写数据，导致flush完成后，一些数据长度小于缓冲区数据长度的现象，报错short write

并发时应每个线程使用一个独立的连接池来保证安全

