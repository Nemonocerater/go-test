package main

import (
	"fmt"
	"bufio"
	"os"
)
var p = fmt.Println

var pos = 0

func moveLeft() {
	if pos > 0 {
		pos--
	}
}

func moveRight() {
	if pos < 10 {
		pos++
	}
}

func render() {
	bytes := []byte{'.','.','.','.','.','.','.','.','.','.'}
	bytes[pos] = 'X'
	var out  = string(bytes) + "\r"
	fmt.Print(out)
}

func main() {
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		lines <- scan.Text()
	}

	for {
		select {
		default:
			render()
		}
	}
}