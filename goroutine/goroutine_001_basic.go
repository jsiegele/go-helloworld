package main

import (
	"fmt"
)

func main() {
	// go - keyword - green Thread - goroutine
	go sayHello()
	// span goroutine - main() finishes - no output
	// time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
