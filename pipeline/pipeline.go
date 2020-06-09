package pipeline

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func multiply(done <-chan interface{}, intStream <-chan int, multiply int) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- i * multiply:
			}
		}
	}()
	return multipliedStream
}

func add(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
	addStream := make(chan int)
	go func() {
		defer close(addStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case addStream <- i + additive:
			}
		}
	}()
	return addStream
}

func ExperimentWithPipeline() {
	done := make(chan interface{})
	defer close(done)

	intStream := generateIntStream(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for p := range pipeline {
		fmt.Println(p)
	}
}

func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan int {
	addStream := make(chan int)
	go func() {
		defer close(addStream)
		for i := range intStream {
			notPrimeNumber := false
			if i == 0 || i == 1 {
				continue
			}
			for v := i - 1; v > 1; v-- {
				if i%v == 0 {
					notPrimeNumber = true
				}
			}
			if notPrimeNumber {
				continue
			}

			select {
			case <-done:
				return
			case addStream <- i:
			}
		}
	}()
	return addStream
}

func InefficientPrimeNumber() {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randFn := func() interface{} { return rand.Intn(5 * 1e7) }
	randIntStream := toInt(done, repeatFn(done, randFn))

	// here we fan-out spawn multiple go routine
	numFinders := runtime.NumCPU()
	finders := make([]<-chan int, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for num := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("%d\n", num)
	}
	fmt.Printf("Search took: %v\n", time.Since(start))
}
