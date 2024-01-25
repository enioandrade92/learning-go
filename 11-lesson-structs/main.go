package main

// structs não é uma classe, mas é o que mais se assemelha a uma
import (
	"fmt"
)

type Client struct {
	name string
	age int
	status bool
}

func main()  {

	enio := Client {
		name: "Enio",
		age: 30,
		status: true,
	}

	fmt.Printf("Name: %s, Age: %d, Status: %t\n", enio.name, enio.age, enio.status)

	enio.status = false

	fmt.Printf("Name: %s, Age: %d, Status: %t", enio.name, enio.age, enio.status)

}



