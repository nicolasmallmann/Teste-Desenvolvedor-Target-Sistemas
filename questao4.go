package main

import "fmt"

var valores = map[string]float64{

	"sp":     67836.43,
	"rj":     36678.66,
	"mg":     29229.88,
	"es":     27165.48,
	"outros": 19849.53,
}

var soma float64

func main() {
	for _, faturamento := range valores {
		soma = soma + faturamento

	}

	for estado, faturamento := range valores {
		var percentual = (faturamento / soma) * 100
		fmt.Printf("%s: %.3f%%\n", estado, percentual)
	}

}
