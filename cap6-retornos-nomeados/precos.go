package main

import "fmt"

func PrecoFinal(precoCusto float64) (precoDolar, precoReal float64) {
	fatorLucro := 1.33
	taxaConversao := 5.25

	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	return precoDolar, precoReal
}

func main() {
	precoDolar, precoReal := PrecoFinal(34.99)

	fmt.Printf("Preço final em dólar: %.2f\n"+
		"Preço final em reais: %.2f\n",
		precoDolar, precoReal)
}
