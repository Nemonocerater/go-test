package main

var = NUM_WORKERS = 5

type Work struct {
	x, y, z int
}

func worker(in <- chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		Sleep(w.z)
		out <- w
	}
}

func run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < NUM_WORKERS; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
}

func sendLotsOfWork(in <- chan *Work) {
}

func receiveLotsOfResults(out chan<- *Work) {
}

func main() {
	run()
}