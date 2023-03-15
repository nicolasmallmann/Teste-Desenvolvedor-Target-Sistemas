package main

import "fmt"

func main() {
	fmt.Print("digite um numero: ")
	var input int64
	fmt.Scanln(&input)

	a := int64(0)
	b := int64(1)

	for a < input {
		a, b = b, a+b
		if a == input {
			fmt.Println("e um numero fibonaci")
			return
		}
	}
	fmt.Println("nao e um numero fibonaci")
	return
}
