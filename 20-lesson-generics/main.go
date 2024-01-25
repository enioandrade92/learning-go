package main

import "fmt"

// não faz sentindo criar duas funções iguais por conta do tipo
func Sum(m map[string]int) int {
	var sum int

	for _, value := range m {
		sum += value
	} 
	return sum
}

 
func SumFloat(m map[string]float64) float64 {
	var sum float64

	for _, value := range m {
		sum += value
	} 
	return sum
}

func main()  {
	salaries := map[string]int{"Enio": 100, "Priscilla": 150, "Laura": 200}
	salariesFloat := map[string]float64{"Enio": 100.00, "Priscilla": 150.50, "Laura": 200.20}

	fmt.Println(Sum(salaries))
	fmt.Println(SumFloat(salariesFloat))
}
