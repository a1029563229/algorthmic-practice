package main

import "fmt"

func balancedStringSplit(s string) int {
	N := 0
	num := 0
	for _, str := range s {
		if string(str) == "L" {
			num++
		} else {
			num--
		}
		if num == 0 {
			N++
		}
	}
	return N
}

func main() {
	s := "RLRRLLRLRL"
	N := balancedStringSplit(s)
	fmt.Println(N)
}
