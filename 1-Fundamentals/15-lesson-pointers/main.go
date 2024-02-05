package main


func main()  {
 	value := 10
	
	println(value) 
	// return value

	println(&value)
	//  return address memory

	var pointers *int = &value
	println(pointers)

	*pointers = 500
	println(value)

	valueAddress := &value
	println(valueAddress)
	println(*valueAddress)


}




