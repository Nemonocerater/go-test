package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func Sleep(dur int64) {
	time.Sleep(2 * time.Second)
}

/******* Worker *******/

type Worker struct {
	requests chan Request
	pending int
	index int
}

func (w *Worker) work(done chan *Worker) {
	for {
		fmt.Println("work(" + strconv.Itoa(w.index) + ")")
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
			b.dispatch(req)
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
	fmt.Println("Dispatching work to " + strconv.Itoa(worker.index) + "[" + strconv.Itoa(worker.pending) + "]")
	worker.requests <- req
}

func (b *Balancer) complete(worker *Worker) {
	//fmt.Println("Completing worker " + strconv.Itoa(worker.index))
	worker.pending--
	fmt.Println("complete: " + strconv.Itoa(worker.index) + "[" + strconv.Itoa(worker.pending) + "]")
}


////////////////////////


func requester(work chan<- Request) {
	c := make(chan int)
	Sleep(rand.Int63n(10000))

	//fmt.Println("...requesting work")
	work <- Request{fn: workfn, c: c}
	//fmt.Println("...requested work")
	result := <-c
	furtherProcess(result)
}

func workfn() int {
	Sleep(10000)
	return int(time.Now().Unix())
}

func furtherProcess(result int) {
	fmt.Println("Result: " + strconv.Itoa(result))
}

func main() {
	input := make(chan Request)

	for i := 0; i < 3; i++ {
		go requester(input)
	}

	var balancer Balancer
	balancer.init(1)
	balancer.balance(input)


/*
	d := make(chan *Worker)
	w := Worker{requests: input, pending: 0, index: 124}
	go w.work(d)

	Sleep(10000)
	for {
		select{
		case req := <-input:
			req.c <- req.fn()
		case w := <-d:
			fmt.Println("ending..." + strconv.Itoa(w.index))
		default:
		}
	}
*/
}
