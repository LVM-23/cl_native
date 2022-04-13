package main

import (
	"fmt"
	"time"
)
/**

编写一个简单的单线程生产者消费者模型队列长度 10 ，
队列元素类型为 int ，生产者：每 1 秒往队列中放入一个类型为 int 的元素，
消费者：队列满时生产者可以阻塞每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

*/

func main() {
	ints := make(chan int, 10)
	timer := time.NewTicker(1 * time.Second)
	times := timer.C
	i := 0

	// 生产者
	go func() {
		for {
			if times != nil {
				time.Sleep( 1 *time.Second)
				ints <- i
				i++
			}
		}

	}()

	//消费者
	for ni := range ints {
		fmt.Println(ni)
	}

}
