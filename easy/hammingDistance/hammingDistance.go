package main

import "fmt"

func hammingDistance(x int, y int) int {
	N := 0
	x2 := convertToBinary(x)
	y2 := convertToBinary(y)
	lx := len(x2)
	ly := len(y2)
	fmt.Println("x2: ", x2)
	fmt.Println("y2: ", y2)
	if lx > ly {
		for ly != len(x2) {
			if x2[0] == 1 {
				N++
			}
			x2 = x2[1:]
		}
	} else {
		for lx != len(y2) {
			if y2[0] == 1 {
				N++
			}
			y2 = y2[1:]
		}
	}

	for i := 0; i < len(x2); i++ {
		if x2[i] != y2[i] {
			N++
		}
	}

	return N
}

func convertToBinary(n int) []int {
	var binary []int
	bus, rem := n, 0
	for bus != 0 {
		bus, rem = computed(bus)
		binary = append([]int{rem}, binary...)
	}

	return binary
}

func computed(n int) (int, int) {
	bus := n / 2
	rem := n % 2
	return bus, rem
}

func main() {
	N := hammingDistance(3, 1)
	fmt.Println(N)
}
