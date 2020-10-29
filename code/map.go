package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)

	go func() { // 接收操作
		for {
			if elem, ok := <-mapChan; ok { // 从通道中获取值
				elem["count"]++
				fmt.Printf("[receiver] The elem is: %v. Elem address is: %p \n", elem, &elem)
			} else {
				break
			}
		}
		fmt.Println("[receiver] Stopped.")
		syncChan <- struct{}{} // 发送同步通知
	}()

	go func() { // 发送操作
		countMap := make(map[string]int) // map为引用类型，发送和接收的元素为同一个值
		for i := 0; i < 5; i++ {
			countMap["count"]++
			mapChan <- countMap // 向通道中输入值
			time.Sleep(time.Millisecond)
			fmt.Printf("[sender] The count map: %v. Count map address: %p \n", countMap, &countMap)
		}
		close(mapChan)         // 关闭通道
		syncChan <- struct{}{} // 发送同步通知
	}()

	<-syncChan //接收同步通知
	<-syncChan //接收同步通知
	fmt.Println(" map为引用类型，发送和接收的元素指向同一个值.")

}
