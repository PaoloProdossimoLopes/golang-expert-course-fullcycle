package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])   // no reduce the cap of slice
	fmt.Printf("len=%d cap=%d %v\n\n", len(s[:4]), cap(s[:4]), s[:4]) // no reduce the cap of slice

	fmt.Printf("len=%d cap=%d %v\n\n", len(s[2:]), cap(s[2:]), s[2:]) // reduce the cap of slice

	fmt.Printf("len=%d cap=%d %v\n\n", len(s[2:4]), cap(s[2:4]), s[2:4]) // reduce the cap of slice

	s = append(s, 110) // This operation shows the duplications of slice internaly managed by golang to resize it
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:4]), cap(s[2:4]), s[2:4])
}
