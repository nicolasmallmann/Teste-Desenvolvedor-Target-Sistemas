package main

import "fmt"

func reverter(str string) string {
	var result string
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func main() {
	fmt.Print("escreva algo: ")
	var input string
	fmt.Scanln(&input)

	fmt.Println("a palavra invertida e:", reverter(input))
}
