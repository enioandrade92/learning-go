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



func main()  {
	enio := Client {
		name: "Enio",
		age: 30,
		status: true,
	}
	enio.disable()

}

func (client Client) disable() {
	client.status = false
	fmt.Printf("Disabled client: %s\n", client.name)
}



