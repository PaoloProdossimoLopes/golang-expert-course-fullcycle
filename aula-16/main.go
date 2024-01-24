package main

// Generic sem constraints
//func soma[T int | float64](m map[string]T) T {
//	var total T
//	for _, valor := range m {
//		total += valor
//	}
//	return total
//}

// Generic com constraint
type Number interface {
	int | float64 | float32 // Nao considera as derivaçoes de tipo
	//~int | float64 | float32//considera a as deriva'voes de tipo para o inteiro
}

func soma[T Number](m map[string]T) T {
	var total T
	for _, valor := range m {
		total += valor
	}
	return total
}

type Mu int

func (m Mu) muuu() {

}

func main() {
	var mu Mu = 1
	mu.muuu()
	soma(map[string]int{"Paolo": 200, "Leticia": 50})
	soma(map[string]float64{"Paolo": 12.3, "Leticia": 72.1})
}
