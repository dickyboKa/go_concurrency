package introduction

import (
	"fmt"
	"sync"
	"time"
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

func UnderstandSelectStatement() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}

	c1 := make(chan interface{}); close(c1)
	c2 := make(chan interface{}); close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Coun: %d\nc2Count: %d\n", c1Count, c2Count)

	done := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)
		close(done)
	}()

	workCounter := 0
	loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(1*time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signaled to stop. \n", workCounter)
}