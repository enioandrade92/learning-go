package main

import "fmt"

const a = "Hello, World!"

func main()  {
	var arrayNumber [3]int
	arrayNumber[0] = 1
	arrayNumber[1] = 2
	arrayNumber[2] = 10
	fmt.Println(arrayNumber)
	fmt.Println(arrayNumber[0])
	fmt.Println(arrayNumber[len(arrayNumber) - 1])

	for i, v := range arrayNumber {
		fmt.Printf("O array no indice %d tem o valor de: %d\n", i, v)
	}
}
