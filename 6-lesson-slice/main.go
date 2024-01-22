package main

import "fmt"

func main()  {
	slice := []int{2,4,6,8,10,20,30,40,50,60,70,80}

	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(slice[:5]), cap(slice[:5]), slice[:5])
	fmt.Printf("len=%d cap=%d %v\n", len(slice[5:]), cap(slice[5:]), slice[5:])

	slice = append(slice, 90, 100)
	fmt.Print(slice)

}
