package main

import "fmt"

type InterfaceVazia interface{}

// Todos implemetam
// Utilize com moderação

func main() {
	// Note como a interface vazia pode ser qualquer coisa
	var x interface{} = 10
	var y interface{} = "ok"

	showType(x)
	showType(y)
}

// aceita qualquer coisa
func showType(any interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor e %v", any, any)
}
