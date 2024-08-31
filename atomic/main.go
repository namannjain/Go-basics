package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var counter int32
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// counter += 1
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	timeTaken := time.Since(start)
	fmt.Println("counter: ", counter)
	fmt.Println("Load counter: ", atomic.LoadInt32(&counter))
	fmt.Println("Time taken: ", timeTaken)
}
