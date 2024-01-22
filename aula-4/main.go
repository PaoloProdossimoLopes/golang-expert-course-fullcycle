package main

import (
	"fmt" //Library to format (go native library)
)

type Id int

func main() {
	var e Id = 1
	fmt.Printf("O Tipo de é %T", e) //print the type
	fmt.Printf("O Tipo de é %v", e) // print the value
}
