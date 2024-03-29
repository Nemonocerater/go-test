package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func Sleep() {
	time.Sleep(2 * time.Second)
}

var NUM_WORKERS = 5

/******* Request *******/

type Request struct {
	fn func() int
	c chan int
}

/******* Worker *******/

type Worker struct {
	requests chan Request
	pending int
	index int
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		done <- w
	}
}

/******* Balancer *******/

type Balancer struct {
	pool []*Worker
	done chan *Worker
}

func (b *Balancer) init(numWorkers int) {
	b.pool = make([]*Worker, numWorkers)
	b.done = make(chan *Worker)
	for i := 0; i < numWorkers; i++ {
		reqs := make(chan Request)
		w := Worker{requests: reqs, pending: 0, index: i}
		b.pool[i] = &w
		go w.work(b.done)
	}
	fmt.Println("Balancer created with " + strconv.Itoa(len(b.pool)) + " workers")
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case w := <-b.done:
			b.complete(w)
		case req := <-work:
			go b.dispatch(req)
		default:
			fmt.Print("...\r")
		}
	}
}

func (b *Balancer) dispatch(req Request) {
	var worker *Worker = b.pool[0]
	for i := 1; i < len(b.pool); i++ {
		if worker.pending > b.pool[i].pending {
			worker = b.pool[i]
		}
	}

	worker.pending++
	worker.requests <- req
	fmt.Println("Dispatching work to " + strconv.Itoa(worker.index) + "[" + strconv.Itoa(worker.pending) + "]")
}

func (b *Balancer) complete(worker *Worker) {
	worker.pending--
}


////////////////////////


func requester(work chan<- Request) {
	c := make(chan int)
	Sleep()

	work <- Request{fn: workfn, c: c}
	result := <-c
	furtherProcess(result)
}

func workfn() int {
	Sleep()
	return int(time.Now().Unix())
}

func furtherProcess(result int) {
	fmt.Println("Result: " + strconv.Itoa(result))
}

func main() {
	input := make(chan Request)

	for i := 0; i < 17; i++ {
		go requester(input)
	}

	var balancer Balancer
	balancer.init(NUM_WORKERS)
	balancer.balance(input)
}
