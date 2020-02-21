package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		fmt.Println("init:\t", j)
		wg.Add(2)
		go func(ch <-chan int) {
			i := <-ch
			fmt.Println("pop:\t", i)
			wg.Done()
		}(ch)
		go func(ch chan<- int) {
			fmt.Println("push:\t", j)
			ch <- j
			wg.Done()
		}(ch)
		wg.Wait()
	}
}
