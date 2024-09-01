package main

import (
	"fmt"
	"sync"
	_ "sync/atomic"
	"time"
)

var (
	count int64
	lock  sync.Mutex
)

func main() {

	for i := 0; i < 1000; i++ {
		go increment()
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("count: %v\n", count)

}

func increment() {
	// atomic.AddInt64(&count, 1)
	lock.Lock()
	count++
	lock.Unlock()
}
