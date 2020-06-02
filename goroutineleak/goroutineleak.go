package goroutineleak

import (
	"fmt"
	"time"
)

func ThisIsLeaking() {

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
func AvoidGoRoutineLeakWithForSelect() {
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
			default:
				fmt.Println("I'm doing nothing")
			}
		}
	}()
	return completed
}