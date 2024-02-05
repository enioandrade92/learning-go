package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		println(i)
	}
	numbers := []int{10,20,30,40,50,60,70}


	for k,v := range numbers {
		println(k,v)
	}
	

	fmt.Printf("Index and value \n")
	for k,v := range numbers {
		println(k,v)
	}

	fmt.Printf("Only value \n")
	for _,v := range numbers {
		println(v)
	}

	fmt.Printf("While mod \n")
	i := 0
	for i < 10 {
		fmt.Println(i)
	}

	

}
