package main

import "fmt"

func main()  {
	salarios := map[string]int{"Enio":100, "Priscilla":200, "Laura":300}
	fmt.Println(salarios["Enio"])
	delete(salarios, "Enio")
	salarios["osk"] = 500
	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}
