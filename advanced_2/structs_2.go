package main

import "fmt"

//Vertex is used for defining vertices in an argand plane
type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2}

	v.X = 12
	fmt.Println(v.X)
}
