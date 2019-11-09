package main

import "fmt"

func game(guess []int, answer []int) int {
	var N int = 0
	for i, n := range guess {
		if n == answer[i] {
			N++
		}
	}
	return N
}

func main() {
	guess := [3]int{1, 2, 3}
	answer := [3]int{1, 2, 3}
	N := game(guess[:], answer[:])
	fmt.Println(N)
}
