package main

import (
	"fmt"
	"lesson-package/math"
	"github.com/google/uuid"
)

func main()  {
	result := math.sum(1,2)
	fmt.Printf("Result: %v", result)
	id := uuid.New()
	fmt.Printf("Result: %v", id)
}

