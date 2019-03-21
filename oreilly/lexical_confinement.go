package main

// https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html

import (
	"fmt"
	"time"
)

func main() {
	//readOnlyChannel();
	writeOnlyChannel();
}

func writeOnlyChannel() {
	chanOwner := func() chan<- int {
		writer := make(chan int, 5)
		go func() {
			for input := range writer {
				fmt.Printf("Received: %d\n", input)
			}
			fmt.Println("Done receiving!")
		}()
		return writer
	}

	publisher := func(writer chan<- int) {
		defer close(writer)
		for i := 0; i <= 5; i++ {
			writer <- i
			time.Sleep(time.Second)
		}
	}

	writer := chanOwner()
	publisher(writer)
}

func readOnlyChannel() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
				time.Sleep(time.Second)
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}