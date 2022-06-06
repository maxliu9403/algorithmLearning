package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRecev1]\n\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{ syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRecev2]\n\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("received a signal from sigRecv1:%s\n", sig)
		}
		fmt.Printf("end [sigRecv1]\n")
		wg.Done()
	}()

	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("received a signal from sigRecv2:%s\n", sig)
		}
		fmt.Printf("end [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("wait for 2 seconds...")
	time.Sleep(2*time.Second)
	fmt.Printf("stop notification")
	signal.Stop(sigRecv1)
	close(sigRecv1)
	fmt.Println("down")

}
