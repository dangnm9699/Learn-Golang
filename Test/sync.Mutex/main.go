// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var total = 0

// func add(wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()
// 	total = total + 1
// 	m.Unlock()
// 	wg.Done()
// }

// func sub(wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()
// 	total = total - 1
// 	m.Unlock()
// 	wg.Done()
// }

// func main() {
// 	var w sync.WaitGroup
// 	var m sync.Mutex
// 	for i := 0; i < 1000000; i++ {
// 		w.Add(1)
// 		if i%2 == 0 {
// 			go add(&w, &m)
// 		} else {
// 			go sub(&w, &m)
// 		}
// 	}
// 	w.Wait()
// 	// Want 0, return 0, correct
// 	fmt.Println("Value of x", total)
// }

package main

import (
	"fmt"
	"sync"
)

var total = 0

func add(wg *sync.WaitGroup) {
	total = total + 1
	wg.Done()
}

func sub(wg *sync.WaitGroup) {
	total = total - 1
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
	// Want 0, return -35/230/-65/..., wrong
	fmt.Println("Value of x", total)
}
