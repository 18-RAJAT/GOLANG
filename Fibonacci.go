package main

import "fmt"

func fib(n int) {
	var n3, n1, n2 int = 0, 0, 1
	for i := 1; i <= n; i++ {

		fmt.Printf(" %d ", n1)
		n3 = n1 + n2
		n1 = n2
		n2 = n3
	}
}

func main() {
	fib(10)
}
