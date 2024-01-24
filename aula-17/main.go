package main

// Criando um comparador generico
// utilize o `compareble` para testar igualdade entre tipos
func compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	compare(1, 2)
	compare(1, 1.1)
}
