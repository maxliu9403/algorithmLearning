package main

import (
	"fmt"
	"sync"
	"time"
)

var syncMap sync.Map

func main() {
	for i := 0; i < 10; i++ {
		go writeMap(i, i)
		go readMap(i)
	}

	time.Sleep(time.Second)
	outPut(syncMap)
}

func readMap(key int) interface{} {
	value, _ := syncMap.Load(key)
	return value
}

func writeMap(key int, value int) {
	syncMap.Store(key, value)

}

func outPut(syncMap sync.Map) {
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
