package main

import (
	"fmt"
)

type Address struct {
	city    string
	street  string
	zipcode string
}

type Position struct {
	pName string
}

type Client struct {
	name     string
	age      int
	isActive bool
	Address  //Composition
	position Position
}

func main() {
	customer := Client{
		name:     "Paolo",
		age:      24,
		isActive: false,
	}
	customer.isActive = true
	fmt.Printf("O %v tem %v e esta %v\n", customer.name, customer.age, func() string {
		if customer.isActive {
			return "ativo"
		}

		return "inativo"
	}())
}
