package confinemen

import (
	"bytes"
	"fmt"
	"sync"
)

// Not Good
func AdHocConfinemen() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

// Good
func LexicalConfinemen() {
	results := chanOwner()
	consumer(results)
}

func chanOwner() <-chan int {
	results := make(chan int, 5)
	go func() {
		defer close(results)
		for i := 0; i <= 5; i++ {
			results <- i
		}
	}()
	return results
}

func consumer(results <-chan int) {
	for result := range results {
		fmt.Printf("Recieved: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

func LexicalConfinemenBuffer() {
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}

func printData(wg *sync.WaitGroup, data []byte) {
	defer wg.Done()

	var buff bytes.Buffer
	for _, d := range data {
		fmt.Fprintf(&buff, "%c", d)
	}
	fmt.Println(buff.String())
}