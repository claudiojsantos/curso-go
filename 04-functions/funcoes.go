package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(text string) string {
		fmt.Println(text)

		return text
	}

	resultado := f("Texto da Função 01")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(20, 10)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSum, _ := calculosMatematicos(20, 10)
	fmt.Println(resultadoSum)

	_, resultadoSub := calculosMatematicos(20, 10)
	fmt.Println(resultadoSub)
}
