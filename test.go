package main

import (
	"fmt"
	"sync"
	atomic2 "sync/atomic"
)

//var atomics atomic2.Value

type readOnly struct {
	m       map[int]int
	amended bool
}


func main() {
	//mapTest := make(map[int]int)
	//mapTest[1] = 1
	//testReadOnly := readOnly{mapTest, false}
	//atomics.Store(testReadOnly)
	////read, _ := atomics.Load().(readOnly)
	//mapTest[1] = 2
	//mapTest[2] = 3
	//atomics.Store(readOnly{m: mapTest})
	//read, _ := atomics.Load().(readOnly)
	//
	//
	////a := atomics.CompareAndSwapPointer()
	////e, ok := read.m[1]
	////fmt.Println(e, ok)
	//fmt.Println(read)
	var sum uint32 = 100
	var wg sync.WaitGroup
	for i := uint32(0); i < 100; i++ {
		wg.Add(1)
		go func(t uint32) {
			defer wg.Done()
			atomic2.CompareAndSwapUint32(&sum, 100, sum+1)
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
}

