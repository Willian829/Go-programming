package main

import (
	"Test/calc"
	"fmt"
)

func main() {
	fmt.Println("Digite um número: ")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Println("Digite outro número: ")
	var num2 int
	fmt.Scanln(&num2)

	result, err := calc.Dividir(num1, num2)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado", result)
}

