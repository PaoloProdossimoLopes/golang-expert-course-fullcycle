package main

var b bool

// Declaration
var (
	a string
	c int
	e float64
)

// With default parameters
var (
	j      = 2
	w bool = true
)

func main() {
	println("bool:", b)
	b = true
	b = false

	println("string:", a)
	a = "ola"
	a = "Mundo"

	println("int:", c)
	c = 1
	c = 4

	println("float:", 2)
	e = 1.2
	e = 4

	//Short hand to declare a variavel (V A R I A B L E)
	shortHandVariable := 23
	println("short hand variable:", shortHandVariable)
}
