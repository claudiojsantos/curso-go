package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}

	fmt.Printf("lent(s) é %d, cap(s) é %d de %v\n", len(s), cap(s), s)
	fmt.Printf("lent(s) é %d, cap(s) é %d de %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("lent(s) é %d, cap(s) é %d de %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("lent(s) é %d, cap(s) é %d de %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 60)
	fmt.Printf("lent(s) é %d, cap(s) é %d de %v\n", len(s), cap(s), s)
}
