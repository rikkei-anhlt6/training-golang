package main

import (
	"fmt"
	"sync"
)

var so = 0

func in(wg *sync.WaitGroup, m *sync.RWMutex) {
	fmt.Printf("Số: %d\n", so)
	m.RUnlock()
	wg.Done()
}
func change(wg *sync.WaitGroup, m *sync.RWMutex) {
	so++
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.RWMutex
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock()
		go in(&wg, &m)
		m.Lock()
		go change(&wg, &m)
	}
	wg.Wait()
	fmt.Println("End main")
}
