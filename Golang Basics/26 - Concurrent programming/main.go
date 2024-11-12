package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("This happens asynchronously")
		wg.Done()
	}()
	fmt.Println("This is synchronous")
	matrix(50000)
	wg.Wait()
}

func matrix(routines int) {
	var wg sync.WaitGroup
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go runAsync(10000, &wg)
	}
	wg.Wait()
}

func runAsync(v int, wg *sync.WaitGroup) {
	for i := 0; i < v; i++ {
		fmt.Printf("%d", i)
	}
	wg.Done()
}
