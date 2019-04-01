package main

import (
	"fmt"
	"math"
)

func main() {
	i, j := 1, 2
	f := math.Sqrt(float64(i*i + j*j))
	z := uint(f)

	fmt.Println(i, j, f, z)
}
