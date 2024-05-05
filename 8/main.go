package main

import "fmt"

func main() {
	fmt.Println(soma(10, 2))
}

func soma(a, b int) (int, bool) {
	if a+b > 10 {
		return a + b, true
	}

	return a + b, false
}
