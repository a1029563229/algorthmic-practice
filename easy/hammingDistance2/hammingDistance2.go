package main

import "fmt"

func hammingDistance(x int, y int) int {
	res := 0
	n := x ^ y
	fmt.Println("n:", n)
	for n != 0 {
		res++
		n = n & (n - 1)
	}
	return res
}

func main() {
	N := hammingDistance(3, 1)
	fmt.Println(N)
}
