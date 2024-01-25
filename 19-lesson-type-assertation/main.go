package main

import "fmt"
 


func main()  {
	var name interface{} = "Enio"

	res, ok := name.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
}





