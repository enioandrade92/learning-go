package main

import (
	"errors"
	"fmt"
)

func main()  {
	total1, err1 := sum(1,2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(total1)
	}


	total2, err2 := sum(1,2,3,4,5)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(total2)
	}



}

func sum(numbers ...int) (int, error) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	if total < 5 {
		return total, errors.New("valor menor que 5")
	}
	return total, nil
}


