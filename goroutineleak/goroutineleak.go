package goroutineleak

import (
	"fmt"
	"math/rand"
	"time"
)

func GouRoutineLeakReadChannel() {

	doWork(nil)
	// Perhaps more work is done here
	time.Sleep(1 * time.Second)
	fmt.Println("Done.")
}

func doWork(strings <-chan string) <-chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("doWork Exited")
		defer close(completed)
		for s := range strings {
			fmt.Println(s)
		}
	}()
	return completed
}

/*
The way this work is, the parent goroutine signal cancellation to its children.
By convention, this signal is usually a read-only channel name done.
The parent goroutine passes this channel to the child goroutine and then
closes the channel when it wants to cancel the child goroutine. Here's an example
*/
func AvoidGouRoutineLeakReadChannel() {
	done := make(chan interface{})
	terminated := doWorkWithDone(done, nil)

	go func() {
		// Cancel the operation after 1second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine")
		close(done)
	}()
	// wait for it to be done
	<-terminated
	fmt.Println("Done.")
}

func doWorkWithDone(done <-chan interface{}, stringsStream <-chan string) <-chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("doWork Exited")
		defer close(completed)
		// here we use the ubiquitous for-select pattern
		for {
			select {
			case s := <-stringsStream:
				// Do something useful
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}()
	return completed
}

func GouRoutineLeakWriteChannel() {
	randStream := newRandStream()
	fmt.Println("3 Random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	// Simulate on going work
	time.Sleep(1 * time.Second)
}

func newRandStream() <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			randStream <- rand.Int()
		}
	}()
	return randStream
}

/*
If a goroutine is responsible for creating a goroutine, it is also responsible
for ensuring it can stop the goroutine
*/
func AvoidGoRoutineLeakWriteValue() {
	done := make(chan interface{})
	randStream := newRandStreamWithDone(done)
	fmt.Println("3 Random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	// Simulate on going work
	time.Sleep(1 * time.Second)
}

func newRandStreamWithDone(done <-chan interface{}) <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			select {
			case randStream <- rand.Int():
			case <-done:
				return
			}

		}
	}()
	return randStream
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
