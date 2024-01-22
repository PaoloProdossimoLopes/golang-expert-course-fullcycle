package main

import "fmt"

func main() {
	exemploDeMap := map[string]int{}
	fmt.Printf("%v\n", exemploDeMap)

	instanciaDeMapUsandoMake := make(map[string]int)
	fmt.Printf("%v\n", instanciaDeMapUsandoMake)

	salarios := map[string]int{
		"Paolo":    2_000,
		"Leticia":  5_000,
		"Fernando": 1200,
		"Carlos":   750,
	}
	fmt.Printf("%v\n", salarios)

	//Removendo um item
	delete(salarios, "Fernando")

	//Adicionado item
	salarios["Bruna"] = 500

	for nomeDoFuncionario, valorDoSalarioDoFuncionario := range salarios {
		fmt.Printf("O Salario de %v é de %v\n", nomeDoFuncionario, valorDoSalarioDoFuncionario)
	}

	for _, valorDoSalarioDoFuncionario := range salarios {
		fmt.Printf("O Salario é de %v\n", valorDoSalarioDoFuncionario)
	}
}
