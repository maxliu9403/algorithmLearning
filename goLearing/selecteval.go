package main

import "fmt"

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}

var nums = []int{1, 2, 3, 4, 5}

func main() {
	select {
	case getChan(0) <- getNums(0):
		fmt.Println("1th case is selected")
	case getChan(1) <- getNums(1):
		fmt.Println("2nd case is selected")
	default:
		fmt.Println("default")
	}
}

func getNums(i int) int {
	fmt.Printf("nums[%d]\n", i)
	return nums[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]

}
