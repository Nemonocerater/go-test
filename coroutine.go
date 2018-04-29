package main

import "fmt"
var p = fmt.Println

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func buff() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	p(<-ch)
	p(<-ch)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x,y)

	buff()
}