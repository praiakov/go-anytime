package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var incremento int64
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incremento, 1)
			fmt.Println(atomic.LoadInt64(&incremento))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("O valor final de incremento ", incremento)
}
