package main

import (
	"fmt"
	"time"
)


// 值类型通过通道传递（通道接收方只是接收到值的副本，不是值的本身，所以对其修改，不会改变值本身）
var mapChan = make(chan map[string]int, 1)

func main() {

	// 信号通道
	syncChan := make(chan struct{}, 2)

	// 接收数据
	go func() {
		for  {

			// 一个一个取出元素
			if elem, ok := <- mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("stop receiver")
		// 发送信号通知已经接收完毕
		syncChan <- struct{}{}
	}()


	// 发送数据
	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i ++ {
			// 发送数据到通道
			mapChan <- countMap
			time.Sleep(time.Microsecond)
			fmt.Printf("the countMap :%v.[sender]\n", countMap)
		}
		// 关闭通道
		close(mapChan)
		// 发送信号告知主goroutine已经完成写
		syncChan <- struct{}{}
	}()
	// 等待子goroutine执行结束
	<- syncChan
	<- syncChan
}

/*
结果值解释：
因为mapChan的元素类型是一个map，属于引用类型，接收方对元素副本进行修改，将会影响到发送方持有的源值
*/
