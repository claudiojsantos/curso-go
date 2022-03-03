package main

import "fmt"

func main() {
	// Aritmeticos
	soma := 1 + 1
	subtracao := 1 - 1
	multiplicacao := 1 * 1
	divisao := 1 / 1

	fmt.Println(soma, subtracao, multiplicacao, divisao)

	var num01 int8 = 10
	var num02 int8 = 15
	sum := num01 + num02
	fmt.Println(sum)

	// Atribuição
	var variavel01 string = "String"
	variavel02 := "String2"
	fmt.Println(variavel01, variavel02)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Lógicos
	fmt.Println("--------------------------")
	fmt.Println(true && true)
	fmt.Println(false || true)
	fmt.Println(false || true)
	fmt.Println(!true)

	// Unários
	numero := 10
	numero++
	fmt.Println(numero)

	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)
}
