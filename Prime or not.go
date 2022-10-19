package main

import "fmt"

func checkPrime(number int) {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf("%d is not prime number \n", number)
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				fmt.Printf("%d is prime number\n", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf("%d is a prime number\n", number)
		}
	}
}
func main() {
	checkPrime(0)
	checkPrime(10)
	checkPrime(3)
	checkPrime(2)
	checkPrime(1)
	checkPrime(5)
}
