package main

import "fmt"

func printNumberFactors(x int) {
	fmt.Printf("Number factors: ")
	divisor := 2
	for x != 1 {
		if x%divisor == 0 {
			fmt.Printf("%d ", divisor)
			x /= divisor
		} else {
			divisor++
		}
	}
}

func main() {
	var x int
	fmt.Scanf("%d", &x)
	printNumberFactors(x)
	anotherMain()
}

func anotherMain() {
	msg := "Kappa"
	fmt.Print(msg)
}
