package main

import (
	"fmt"
)

type Client struct {
	name     string
	age      int
	isActive bool
}

// attach
func (customer Client) disable() {
	customer.isActive = false
}

func (customer Client) enable() {
	customer.isActive = true
}

func main() {
	customer := Client{
		name:     "Paolo",
		age:      24,
		isActive: false,
	}
	customer.isActive = true
	customer.enable()
	customer.disable()
	fmt.Printf("O %v tem %v e esta %v\n", customer.name, customer.age, func() string {
		if customer.isActive {
			return "ativo"
		}

		return "inativo"
	}())
}
