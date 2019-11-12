package main

import "fmt"

func oddCells(n int, m int, indices [2][2]int) int {
	var N int
	arr := make([][]int, n, m)
	fmt.Println(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Println(i)
			fmt.Println(j)
			arr[i][j] = 10
			fmt.Println(arr)
		}
	}
	fmt.Println(arr)
	return N
}

func main() {
	indices := [2][2]int{
		{0, 1},
		{1, 1}}
	N := oddCells(2, 3, indices)
	fmt.Println(N)
}
