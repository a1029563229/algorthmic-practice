package main

import "fmt"

func anagramMappings(A []int, B []int) []int {
	C := make([]int, len(A))
	M := map[int]int{}
	for i, b := range B {
		M[b] = i
	}
	for i, a := range A {
		C[i] = M[a]
	}
	return C
}

func main() {
	A := []int{12, 28, 46, 32, 50}
	B := []int{50, 12, 32, 46, 28}
	C := anagramMappings(A, B)
	fmt.Println(C)
}
