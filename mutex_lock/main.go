package main

import (
    "fmt"
    "sync"
)

func main() {
    var counter int
    // var mu sync.Mutex
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()

            // mu.Lock()
            counter++
			fmt.Println(counter)
            // mu.Unlock()
        }()
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
