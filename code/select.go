select
 	监听channel有关的IO操作。 
 	select case must be receive, send or assign recv
 	//select基本用法
	select {
	case <- chan1:
	// 如果chan1成功读到数据，则进行该case处理语句
	case chan2 <- 1:
	// 如果成功向chan2写入数据，则进行该case处理语句
	default:
	// 如果上面都没有成功，则进入default处理流程
	}
如果有多个IO操作，会随机选择一个，否则执行default，如果没有default，则一直阻塞，直到至少有一个io可执行。
case 里的所有表达式都会被求值，自左向右，自上而下。
break 退出select
