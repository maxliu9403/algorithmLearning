package main

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	var mutex sync.Mutex
//	fmt.Println("lock teh lock.(main)")
//	mutex.Lock()
//
//	fmt.Println("the lock is locked")
//	for i := 1; i <= 3; i++ {
//		go func(i int) {
//			fmt.Printf("lock the lock (g%d)\n", i)
//			mutex.Lock()
//			fmt.Printf("the lock is locked .(g%d)\n", i)
//
//		}(i)
//	}
//
//	time.Sleep(time.Second)
//	fmt.Println("unlock the lock. (main)")
//	mutex.Unlock()
//	fmt.Println("the lock is unlocked. (main)")
//	time.Sleep(time.Second)
//
//}
//
//func main() {
//	defer func() {
//		fmt.Println("Try to recover the panic.")
//		if p := recover(); p != nil {
//			fmt.Printf("Recovered the panic(%#v).\n", p)
//		}
//	}()
//	var mutex sync.Mutex
//	fmt.Println("Lock the lock.")
//	mutex.Lock()
//	fmt.Println("The lock is locked.")
//	fmt.Println("Unlock the lock.")
//	mutex.Unlock()
//	fmt.Println("The lock is unlocked.")
//	fmt.Println("Unlock the lock again.")
//	mutex.Unlock()
//}

func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try to lock for reading... [%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked for reading. [%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try to unlock for reading... [%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlocked for reading. [%d]\n", i)
		}(i)
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try to lock for writing...")
	rwm.Lock()
	fmt.Println("Locked for writing.")
}

