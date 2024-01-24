package main

import "fmt"

func main() {
	//Type assertion - coinversao de tipo
	var minhavar interface{} = 10
	println(minhavar.(string))

	result, isOK := minhavar.(int)
	fmt.Printf("result=%v | isOK=%v", result, isOK)
}
