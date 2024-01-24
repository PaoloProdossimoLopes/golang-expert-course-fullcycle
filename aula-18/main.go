package main

import "fmt"

func main() {
	// Option 1
	for i := 0; i < 10; i++ {
		println(i)
	}

	// Option 2
	lista := []string{"um", "dois", "tres"}
	for i, v := range lista {
		fmt.Printf("pocicao=%v valor=%v\n", i, v)
	}

	// Option 3
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// Option 4
	for {
		print("Hellow")
	}
}
