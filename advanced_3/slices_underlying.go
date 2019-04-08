package main

import "fmt"

func main() {
	name := [4]string{
		"John",
		"Jane",
		"Mary",
		"Mark",
	}

	fmt.Println(name)

	var b = name[0:2]
	b[1] = "Lolcats"
	fmt.Println(name)

}
