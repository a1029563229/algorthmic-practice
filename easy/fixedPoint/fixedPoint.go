package main

import "fmt"

func fixedPoint(A []int) int {
	N := -1
	L := 0
	R := len(A) - 1
	for L < R {
		mid := (L + R) / 2
		if A[mid] > mid {
			R = mid
		}
		if A[mid] < mid {
			L = mid + 1
		}
	}
	fmt.Println(L)
	return N
}

func main() {
	A := []int{0, 1, 3, 7, 8, 9}
	N := fixedPoint(A[:])
	fmt.Println(N)
}
