package main

import "fmt"

//Compl is a function to denote complex numbers
type Compl struct {
	Real      float64
	Imaginary float64
}

func main() {

	com := Compl{12, -15}
	p := &com
	p.Real = 1e9
	fmt.Println(com)
}
