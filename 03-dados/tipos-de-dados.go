package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero02 uint32 = 1000
	fmt.Println(numero02)

	// aliases
	// int32 = rune
	var numero03 rune = 12345
	fmt.Println(numero03)

	// byte = uint8
	var numero04 byte = 123
	fmt.Println(numero04)

	// Numeros Reais
	var numeroReal01 float32 = 123.45
	fmt.Println(numeroReal01)

	var numeroReal02 float64 = 12334555444.44
	fmt.Println(numeroReal02)

	numeroReal03 := 1234.44
	fmt.Println(numeroReal03)

	// Strings
	var str string = "Claudio"
	fmt.Println(str)

	str2 := "Santos"
	fmt.Println(str2)

	char := 'C'
	fmt.Println(char)

	// Valor Zero
	var texto string
	fmt.Println(texto)

	var texto02 int16
	fmt.Println(texto02)

	// Boolean
	var booleano01 bool = true
	fmt.Println(booleano01)

	// Error
	var erro error
	fmt.Println(erro)

	var erro01 error = errors.New("Erro interno")
	fmt.Println(erro01)
}
