package main

import (
	"fmt"
	"time"
)

// 创建一个带有3个缓冲元素的channel
var strChan = make(chan string, 3)

func main() {
	// 初始化2个信号channel
	synChan1 := make(chan struct{}, 1)
	synChan2 := make(chan struct{}, 2)

	// 演示接收操作
	go func() {
		// 等待发送方信号，没有数据一直阻塞未执行
		fmt.Println("start")
		<- synChan1
		fmt.Println("Received a sync signal and wait a second...[receiver]")
		time.Sleep(time.Second)

		for  {
			if elem, ok := <- strChan; ok {
				fmt.Println("received:", elem, "[received]")
			} else {
				// 取出所有值中断
				break
			}
		}

		fmt.Println("stopped.[received]")

		synChan2 <- struct{}{}

	}()

	// 用于演示发送通道
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("send:", elem, "[send]")

			// 缓冲通道只有3个元素，如果有三个元素了，则需要发出信号
			if elem == "c" {
				synChan1 <- struct{}{}
				fmt.Println("send a sync signal. [send]")
			}
		}
		// 等到接收方接收完所有的元素
		fmt.Println("wait 2s, [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		synChan2 <- struct{}{}

	}()

	// 等待两个goroutine执行结束，让主goroutine等待
	<- synChan2
	<- synChan2

}
