package pipeline

import (
	"fmt"
	"math/rand"
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

	intStream := toIntStream(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for p := range pipeline {
		fmt.Println(p)
	}
}

func primeFinder(done <-chan interface{}, intStream <-chan interface{}) <-chan int {
	addStream := make(chan int)
	go func() {
		defer close(addStream)
		for i := range intStream {
			notPrimeNumber := false
			valInt := i.(int)
			if valInt == 0 || valInt == 1 {
				continue
			}
			for v := valInt - 1; v > 1; v-- {
				if valInt%v == 0 {
					notPrimeNumber = true
				}
			}
			if notPrimeNumber {
				continue
			}

			select {
			case <-done:
				return
			case addStream <- valInt:
			}
		}
	}()
	return addStream
}

func InefficientPrimeNumber() {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()
	for num := range take(done, primeFinder(done, repeatFn(done, func() interface{} { return rand.Intn(10) })), 10) {
		fmt.Printf("%d\n", num)
	}
	fmt.Printf("Search took: %v\n", time.Since(start))
}
