package main

import "fmt"

func main() {
	fmt.Println("Hello, World!");
	add, mult := calc(3, 8);
	fmt.Println(add);
	fmt.Println(mult);
	fmt.Println(calc(1,2));
	/*
	for i := 1; i < 5; i++ {
		fmt.Print(i, " ")
	}
	*/
	var sum = 1
	for ;sum < 1000; {
		sum += sum
		fmt.Print(sum, " ")
	}
}

func calc(x, y int) (a, b int) {
	a = x + y
	b = x * y
	return
}


