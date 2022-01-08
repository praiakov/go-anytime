package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Números de CPUS inicio: %v\n", runtime.NumCPU())
	fmt.Printf("Números de Goroutinas inciio : %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("goroutine 1")

		wg.Done()
	}()

	go func() {
		fmt.Println("goroutine 2")

		wg.Done()
	}()

	fmt.Printf("Números de CPUS meio: %v\n", runtime.NumCPU())
	fmt.Printf("Números de Goroutinas meio : %v\n", runtime.NumGoroutine())

	wg.Wait()

	fmt.Printf("Números de CPUS fim: %v\n", runtime.NumCPU())
	fmt.Printf("Números de Goroutinas fim : %v\n", runtime.NumGoroutine())

}
