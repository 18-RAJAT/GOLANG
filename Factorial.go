package main

import "fmt"

func main() {
	var nFact, fact int
	fact = 1

	fmt.Println("Enter the number ")
	fmt.Scanln(&nFact)

	for i := 1; i <= nFact; i++ {
		fact = fact * i
	}
	fmt.Println("factorial is", nFact, " =", fact)
}
