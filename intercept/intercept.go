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
	var wg sync.WaitGroup
	wg.Add(1)

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	stopChan := make(chan struct{})

	go goWithTimer(10, &wg, stopChan)

	go func() {
		s := <-sigChan
		fmt.Println("\nReceived signal:", s)
		close(stopChan)
	}()

	wg.Wait()
	fmt.Println("All go routine completed")
}

func goWithTimer(tm int, wg *sync.WaitGroup, stopChan <-chan struct{}) {
	defer wg.Done()
	round := 1
	for {
		for i := 0; i < tm; i++ {
			fmt.Printf("From %d go -> %d\n", round, i)
			time.Sleep(time.Second)
		}
		select {
		case <-stopChan:
			fmt.Println("Stopping goroutine at round:", round)
			return
		default:
		}
		round++
	}
}
