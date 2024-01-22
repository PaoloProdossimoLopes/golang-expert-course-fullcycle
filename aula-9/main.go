package main

func main() {
	println("O valor da soma de 2 + 5 =", sum(2, 3))
	println(sumTuple(2, 3))
	valor, isCorrect := sumTuple(2, 3)
	println(valor, isCorrect)
	println("total:", sumAll(1, 2, 3, 4, 5))
}

// func sum(a int, b int) int {
func sum(a, b int) int { // Sintax suggar: when parameters have the same type
	return a + b
}

func sumTuple(a, b int) (int, bool) {
	total := a + b
	return total, total > 50
}

func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
