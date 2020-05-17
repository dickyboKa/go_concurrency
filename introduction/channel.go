package introduction

import (
	"fmt"
	"sync"
)

func PlayAroundWithChannel() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)

	// try to read from closed channel
	intStream := make(chan int)
	close(intStream)
	i, ok := <- intStream
	fmt.Printf("(%v): %v\n", ok, i)

	// range/loop over channel
	intStream = make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for i := range intStream {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	// example of unblocking multiple channel all together
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}