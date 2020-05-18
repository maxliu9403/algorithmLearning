package main

import (
	"fmt"
	"time"
)

var strChan1 = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 3)
	go receiver(strChan1, syncChan1, syncChan2)
	go sender(strChan1, syncChan1, syncChan2)

	<- syncChan2
	<- syncChan2
}

//参数：接收strChan的元素 接收开始接收信号的单向通道 发送接收完毕的单向通道
func receiver(strChan <- chan string,
	syncChan1 <- chan struct{},
	syncChan2 chan <- struct{})  {

	// 接收信号通道
	<- syncChan1
	fmt.Println("start receive")

	time.Sleep(time.Second)

	for elem := range strChan {
		fmt.Println(elem)
	}

	fmt.Println("stop receive")
	// 接收完毕发送信号
	// 告知主程序可以结束
	syncChan2 <- struct{}{}

}

func sender(strChan chan<- string,
	syncChan1 chan<- struct{},
	syncChan2 chan<- struct{}) {

	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("send, [sender]", elem)

		if elem == "c" {
			syncChan1 <- struct{}{}

			fmt.Println("send a sync signal, [sender]")

		}

	}
	fmt.Println("wait 2s")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}