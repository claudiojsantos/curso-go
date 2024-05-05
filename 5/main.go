package main

import "fmt"

func main() {
	var meuArray [3]int // array fixo
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray) - 1)
	fmt.Println(len(meuArray))
	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Println(i, v)
	}
}
