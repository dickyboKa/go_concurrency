package main

import "fmt"

func main() {
	//i.DataRace()

	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)

}
