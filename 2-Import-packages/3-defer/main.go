package main

import (
	"io"
	"net/http"
)

func print() {
	// defer executa apenas no final da função ou processo
	defer println("First")
	println("Second")
	println("Third")
}

func main() {
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))

	print()
}
