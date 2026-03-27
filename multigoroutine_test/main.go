package main

import (
	"fmt"
	// "time"

	// "github.com/kapetan-io/tackle/wait"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// ch := make(chan struct{},5)
	// go func() {
	// 	time.Sleep(time.Second*2)
	// 	for i := 1; i <= 4; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	ch <- struct{}{}
	// 	// wg.Done()
	// }()
	// // fmt.Println("HERE")

	// // fmt.Println("THERE")

	// go func() {
	// 	for i := 5; i <= 10; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	ch <- struct{}{}
	// 	// wg.Done()
	// }()
	// <- ch
	// <- ch
	// wg.Wait()


	ch1 := make(chan int, 2)
	go func(){
		for {
			switch <-ch1{
			case 1:
				go func() {
					for i := 1; i <= 4; i++ {
						fmt.Println(i)
					}
				}()
			case 2:
				go func() {
					for i := 5; i <= 10; i++ {
						fmt.Println(i)
					}
				}()
			}
		}
	}()
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	wg.Wait()
}