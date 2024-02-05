package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	// creating file
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	// add string
	sizeString, err := f.WriteString("Hello world!\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created successfully, size: %d bytes\n", sizeString)

	// writing data
	sizeByte, err := f.Write([]byte("Writing data\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created successfully, size: %d bytes\n", sizeByte)
	f.Close()

	// reading file
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	// convert from byte to string
	fmt.Println(string(file))

	// reading in parts
	readFile, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(readFile)
	buffer := make([]byte,1)

	for {
		n, err :=reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}

}
