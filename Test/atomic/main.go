package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total int64 = 0

func add(wg *sync.WaitGroup) {
	atomic.AddInt64(&total, 1)
	wg.Done()
}

func sub(wg *sync.WaitGroup) {
	atomic.AddInt64(&total, -1)
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		w.Add(1)
		if i%2 == 0 {
			go add(&w)
		} else {
			go sub(&w)
		}
	}
	w.Wait()
	// Value of x: 0
	fmt.Println("Value of x: ", total)
}
