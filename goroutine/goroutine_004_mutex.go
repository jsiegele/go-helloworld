package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		// durch dieses Konstrukt erhalten wir ein schönes (inperformates) sequentielles Programm.
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
}

func sayHello() {
	fmt.Println("Hello", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
