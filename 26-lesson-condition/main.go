package main

import "fmt"

func main()  {
	a := 1
	b := 1
	c := 2

	if a == b {
		fmt.Printf("equal")
	// not exist else if
	} else {
		fmt.Printf("any")
	}

	if a+b == c {
		fmt.Printf("true")
	}

	switch a {
		case 1:
			println("a")
		case 2:
			println("b")
		case 3:
			println("b")
		default:
			println("d")
	}
		

}
