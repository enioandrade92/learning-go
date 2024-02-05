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

type Person interface {
	disable()
}

type Client struct {
	name string
	age int
	status bool
	Address Address
}

func (client Client) disable() {
	client.status = false
	fmt.Printf("Disabled client: %s\n", client.name)
}

func Disabling(person Person)  {
	person.disable()
}


func main()  {
	enio := Client {
		name: "Enio",
		age: 30,
		status: true,
	}
	Disabling(enio)
	fmt.Printf("Disabled client: %t\n", enio.status)

}




