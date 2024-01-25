package main

// structs não é uma classe, mas é o que mais se assemelha a uma
import (
	"fmt"
)

type Address struct {
	street string
	number int
	city string
	state string
}


type Client struct {
	name string
	age int
	status bool
	Address Address
}

// no caso abaixo é uma composição você consegue acessar cliente.city
// type Client struct {
// 	name string
// 	age int
// 	status bool
// 	Address
// }

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



