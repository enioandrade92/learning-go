package main

import "fmt"

type Client struct {
	name string
}

func (client *Client) walked() {
	client.name = "osk"
	fmt.Printf("Client %s walked\n", client.name)
}

func main()  {
	enio := Client {
		name: "Enio",
	}
	enio.walked()

	fmt.Printf("Client %s name\n", enio.name)

}




