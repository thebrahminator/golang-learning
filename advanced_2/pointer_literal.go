package main

import "fmt"

//Names is used to denote name of a person
type Names struct {
	firstName string
	lastName  string
}

var (
	v1 = Names{"12", "21"}
	v2 = Names{firstName: "21"}
	v3 = Names{}
	p  = &Names{"15", "21"}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
