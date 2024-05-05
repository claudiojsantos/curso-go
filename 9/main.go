package main

import "fmt"

func main() {
	fmt.Println(soma(10, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func soma(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
