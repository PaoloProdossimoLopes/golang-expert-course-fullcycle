package main

import "strconv"

func main() {
	println("O valor da soma de 2 + 5 =", sum(2, 3))
	println(sumTuple(2, 3))
	valor, isCorrect := sumTuple(2, 3)
	println(valor, isCorrect)

	block := func(value int) string {
		return "Olá " + strconv.Itoa(value)
	}

	clo(block)
}

// func sum(a int, b int) int {
func sum(a, b int) int { // Sintax suggar: when parameters have the same type
	return a + b
}

func sumTuple(a, b int) (int, bool) {
	total := a + b
	return total, total > 50
}

func clo(block func(int) string) {
	println(block(32))
}
