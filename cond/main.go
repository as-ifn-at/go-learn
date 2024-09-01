package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	a := 6
	c := sync.NewCond(&sync.Mutex{})
	wg := &sync.WaitGroup{}
	// wg.Add(1)
	go func(a *int, c *sync.Cond, wg *sync.WaitGroup) {
		// defer wg.Done()
		defer gorountineEnded()
		c.L.Lock()
		for *a == 6 {
			fmt.Println("inside infinite loop going to sleep")
			c.Wait()
			// time.Sleep(2 * time.Second)
			fmt.Println("go routine woke up")
		}
		c.L.Unlock()
	}(&a, c, wg)

	time.Sleep(5 * time.Second)
	a++
	time.Sleep(1 * time.Second)
	go func() {
		if a != 6 {
			c.Signal()
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Printf("a: %v\n", a)
	// c.Broadcast()
	// wg.Wait()
	time.Sleep(10 * time.Second)
}

func gorountineEnded() {
	fmt.Println("gorountineEnded called!")
}
