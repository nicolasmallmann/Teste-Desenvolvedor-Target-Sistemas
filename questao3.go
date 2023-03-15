package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
)

type Dias struct {
	Dia   int     `json:"dia"`
	Valor float64 `json:"valor"`
}

func main() {
	dadosjson, err := os.Open("dados.json")
	if err != nil {
		log.Fatal(err)
	}
	defer dadosjson.Close()

	var dias []Dias
	decoder := json.NewDecoder(dadosjson)
	err = decoder.Decode(&dias)
	if err != nil {
		log.Fatal(err)
	}

	var soma, maiorvalor float64
	var numdias int16
	menorvalor := math.MaxFloat64

	for _, dia := range dias {
		if dia.Valor > 0 {
			soma = soma + dia.Valor
			numdias++

			if dia.Valor > maiorvalor {
				maiorvalor = dia.Valor
			}

			if dia.Valor < menorvalor {
				menorvalor = dia.Valor
			}
		}
	}

	media := soma / float64(numdias)

	var numeroDiasSuperiorMedia float64
	for _, dia := range dias {
		if dia.Valor > media {
			numeroDiasSuperiorMedia++
		}
	}

	fmt.Println("o menor valor de faturamento em um dia do mes e:", menorvalor)
	fmt.Println("o maior valor de faturamento em um dia do mes e:", maiorvalor)
	fmt.Println("o numero de dias no mes em que o valor diario foi superior a media mensal e:", numeroDiasSuperiorMedia)
}
