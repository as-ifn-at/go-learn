package main

import (
	"fmt"
	"sync"
)

func main() {

	var count int
	increment := func() {
		fmt.Println("increment called!")
		count++
	}
	// The sync Package | 57
	var once sync.Once
	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
