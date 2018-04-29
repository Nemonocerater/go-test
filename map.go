package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m["Josh's House"] = Vertex{
		90.90, 38.829,
	}
	m["Jill's House"] = Vertex{
		1, 89.32,
	}
	delete(m, "Google")

	goog, goog_ok := m["Google"]
	if goog_ok {
		fmt.Print("Google (shouldn't be printing) ")
		fmt.Println(goog)
	}

	bell, bell_ok := m["Bell Labs"]
	if bell_ok {
		fmt.Print("Bell Labs ")
		fmt.Println(bell)
	}

	fmt.Println(m)
}