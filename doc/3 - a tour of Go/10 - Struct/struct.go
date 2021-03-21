package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implict
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{2, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	fmt.Println(v1, p, v2, v3, v1)
}
