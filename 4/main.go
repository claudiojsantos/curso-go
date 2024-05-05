package main

import "fmt"

type ID int

var (
	e float64 = 10.3
	f ID      = 5
)

func main() {
	fmt.Printf("O tipo de e é %T\n", e)
	fmt.Printf("O valor de e é %v\n", e)
	fmt.Printf("O tipo de f é %T\n", f)
	fmt.Printf("O valor de f é %v\n", f)
}
