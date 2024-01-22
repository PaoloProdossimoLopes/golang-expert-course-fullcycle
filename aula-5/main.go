package main

import "fmt"

func main() {
	var lista [3]string
	fmt.Println("Primeira possição antes de iniciar: ", lista[0])
	lista[0] = "posição 1"
	lista[1] = "posição 2"
	lista[2] = "posição 3"
	fmt.Println("Primeira possição pos de iniciar: ", lista[0])
	fmt.Println("O tamanho do meu array é:", len(lista))
	fmt.Println("O meu ultimo valor no array :", lista[len(lista)-1])

	println("Percorrendo minha lista ...")
	for indice, valor := range lista {
		fmt.Printf("O INDICE é %v e o VALOR é %v\n", indice, valor)
		//fmt.Printf("O INDICE é %d e o VALOR é %s\n", indice, valor) tbm funciona
	}
}
