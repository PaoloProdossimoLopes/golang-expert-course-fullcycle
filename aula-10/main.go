package main

func main() {
	valorDaClosureExecutada := func() int {
		return 3
	}()
	println(valorDaClosureExecutada)

	closure := func() string {
		return "isso é um exemplo de closure"
	}
	println(closure())
}
