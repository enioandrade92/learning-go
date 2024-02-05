package main

import (
	"errors"
	"fmt"
)

func main()  {
	fmt.Println(sum(1,2))
	fmt.Println(sum(50,2))
	fmt.Println(sum(3,2))
	fmt.Println(sum(50,1))
	
	value1, err1 := sumOrError(1,2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(value1)
	}

	value2, err2 := sumOrError(50,2)
	if err2 != nil {
		fmt.Println(value2, err2)
	} else {
		fmt.Println(value2)
	}


}

func sum(a int, b int) (int, bool) {
	sum := a + b
	if sum >= 50 {
		return sum, true
	}
	return sum, false
}

func sumOrError(a int, b int) (int, error) {
	sum := a + b
	if sum >= 50 {
		return sum, errors.New("A soma é maior que 50")
	}
	return sum, nil
}
