package main

import "fmt"

func main() {
	salarios := map[string]float64{"Claudio": 1000.0, "João": 2000.0, "Maria": 3000.0}
	fmt.Println(salarios)
	fmt.Println(salarios["Claudio"])
	delete(salarios, "Maria")
	fmt.Println(salarios)
	salarios["Pedro"] = 1200.0
	fmt.Println(salarios)

	sal := make(map[string]float64)
	sal["José"] = 1500.0
	fmt.Println(sal)

	sal1 := map[string]float64{}
	sal1["Paulo"] = 1500.0
	fmt.Println(sal1)

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	// blank identifier
	for _, salario := range salarios {
		fmt.Println(salario)
	}
}
