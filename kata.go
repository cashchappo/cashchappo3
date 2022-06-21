package main

import "fmt"

func main() {
	h := 0
	m := 4
	s := 1

	fmt.Println(Past(h, m, s))

}

func Past(h, m, s int) int {

	h1 := h * 3600000
	m1 := m * 60000
	s1 := s * 1000

	result := h1 + m1 + s1

	return result
}
