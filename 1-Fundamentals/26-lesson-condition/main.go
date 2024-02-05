package main

import "fmt"

func main()  {
	a := 1
	b := 1
	c := 2

	if a == b {
		fmt.Printf("equal \n")
	// not exist else if
	} else {
		fmt.Printf("any \n")
	}

	if a+b == c {
		fmt.Printf("true \n")
	}

	switch a {
		case 1:
			println("a \n")
		case 2:
			println("b")
		case 3:
			println("b")
		default:
			println("d")
	}
		

}
