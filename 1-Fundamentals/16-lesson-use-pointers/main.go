package main

import "fmt"

func add(value int) {
	value = 11
}

func addInMemory(value *int) {
	*value = 11
}

func main()  {
 	value1 := 10

	println(value1) 
	// return value
	add(value1)
	fmt.Printf("Na memoria o valor de value1 ainda é o mesmo, que é: %d \n", value1)
	
	memoryAddressValue := &value1
	println(*memoryAddressValue)
	addInMemory(memoryAddressValue)
	fmt.Printf("Agora sim o valor é alterado na memoria: %d \n", value1)
	println(value1)
	




}




