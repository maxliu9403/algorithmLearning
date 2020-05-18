package main

import (
	"fmt"
	"time"
)

// 单向通道的使用

var strChanss = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChanss, syncChan1, syncChan2) // 用于演示接收
	go send(strChanss, syncChan1, syncChan2)    // 用于演示发送

	<-syncChan2
	<-syncChan2

}

// syncChan1 <- chan struct{} 单向接收通道
// syncChan2 chan <- struct{} 单向发送通道
func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {

	// 等待接收信号
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second...[receiver]")
	time.Sleep(time.Second)

	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("received", elem)
		} else {
			break
		}
	}
	fmt.Println("stop received")
	syncChan2 <- struct{}{}

}

func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {

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
