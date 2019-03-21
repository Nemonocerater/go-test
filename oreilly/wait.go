package main

import (
	"fmt"
	"sync"
	"time"
)

type bob struct {
	b *time.Time
}

func printData(wg *sync.WaitGroup, num int) {
	defer wg.Done()

	time.Sleep(time.Second * time.Duration(num))

	// Change this to a number that is used to see it seg fault and crash the whole process
	if num == 1000 {
		var b bob
		fmt.Printf("%d", b.b.Day())
	}

	fmt.Printf("%d\n", num)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printData(&wg, i)
	}

	wg.Wait()
}
