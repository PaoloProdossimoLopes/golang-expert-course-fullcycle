package matematica

// Letras maisuculas incida que sao exportados
func Soma[T int | float64](a, b T) T {
	return a + b
}
