package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

// 值传递和引用传递问题

var mapChans = make(chan map[string]*Counter, 1)
//var mapChans = make(chan map[string]Counter, 1)

func (counter * Counter) String() string  {
	return fmt.Sprintf("{count: %d}", counter.count)

}

func main() {
	// 信号通道
	syncChan := make(chan struct{}, 2)

	// 开启接收
	go func() {
		for {
			if elem, ok := <-mapChans; ok {
				counter := elem["count"]
				counter.count++
				value := counter.String()
				fmt.Print(value)
			} else {
				break
			}
		}
		fmt.Println("stop receiver")
		// 发送信号通知已经接收完毕
		syncChan <- struct{}{}

	}()

	// 开启发送通道
	go func() {
		countMaps := map[string]*Counter{
			"count": &Counter{},
		}

		//countMaps := map[string]Counter{
		//	"count": Counter{},
		//}

		for i:=0; i<5; i++ {
			mapChans <- countMaps
			time.Sleep(time.Microsecond)
			fmt.Printf("the countMap :%v.[sender]\n", countMaps)
		}
		close(mapChans)
		syncChan <- struct{}{}
	}()

	<- syncChan
	<- syncChan
}


/*
结果解释：
1.值传递和引用传递
发送方向通道发送的值会被复制，接收方接收到的总是该值的副本，而不是值本身。
因此，当接收方从通道中接收到是一个值类型的值时，对该值的修改就不会影响到发送方持有的数据。
但是对于引用类型的值来说，这种修改会同时影响到收发双发的持有的值

*/