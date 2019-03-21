package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		// This should print if the infinite loop below isn't blocking
		fmt.Println("Async thingy finished")
	}()

	fmt.Println("Sync thingy starting")

	a := 0
	for {
		if a > 100 {
			break;
		}
	}

	// This should never actually be hit
	fmt.Println("Sync thingy finished")
}