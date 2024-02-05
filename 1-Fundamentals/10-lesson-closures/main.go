package main

import (
	"fmt"
)

func main()  {
	total := func() (int) {
		return sum(1,3) * 2
	}()

	fmt.Println(total)
	

}

func sum(numbers ...int) (int) {
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}


