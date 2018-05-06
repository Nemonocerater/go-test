package main

import (
    "fmt"
        "time"
        )

func main() {

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	fmt.Println("Sleeping...")
	time.Sleep(2 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

}