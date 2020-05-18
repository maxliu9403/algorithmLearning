package main

import "fmt"

// 通道断言
// 单向通道是由双向通道转换而来（其实就是利用语法级别的约束）
// 但是单向通道是不可能转换为双向通道的
// 因为通道允许的数据传递方向是其类型的一部分
// 数据传递方向的同步就意味着他们类型不同

func main()  {
	var ok bool
	// 做类型的转换
	ch := make(chan int, 1)  // 双向通道


	// 转换为发送通道
	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <- chan int:", ok)

	// 转换为接收通道
	_, ok = interface{}(ch).(chan <- int)
	fmt.Println("chan int => chan<- int:", ok)

	// 发送通道
	sch := make(chan <- int, 1)
	// 转换为双向通道
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan int => chan int:", ok)

	// 接收通道
	rch := make(<- chan int, 1)
	// 转换为双向通道
	_, ok = interface{}(rch).(chan int)
	fmt.Println("chan int => chan<- int:", ok)





}

