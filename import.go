package main

import (
	"fmt"
	helper "./helper"
	blah "./helper/blah"
)
var p = fmt.Println

func main() {
	v := helper.Vertex{0.1, 0.56}
	p(v.Abs())

	a := helper.Vertex{1, 3}
	b := helper.Vertex{3, 4}
	describe(a)
	p(a.Dist(b))

	p(blah.DoBlah())
}
