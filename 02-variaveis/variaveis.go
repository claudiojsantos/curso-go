package main

import "fmt"

func main() {
	var variavel01 string = "Variavel 01"
	fmt.Println(variavel01)

	variavel02 := "Variavel 02"
	fmt.Println(variavel02)

	var (
		variavel03 string = "Variavel03"
		variavel04 string = "Variavel04"
	)

	fmt.Println(variavel03, variavel04)

	variavel05, variavel06 := "Variavel05", "Variavel06"
	fmt.Println(variavel05, variavel06)

	const contante01 string = "Contante01"
	fmt.Println(contante01)

	variavel05, variavel06 = variavel06, variavel05
	fmt.Println(variavel05, variavel06)

}
