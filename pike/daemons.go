package main

func main() {
	go func() { // copy input to output
		// for-range drains the channel
		for val := range input {
			output <- input
		}
	}()
}