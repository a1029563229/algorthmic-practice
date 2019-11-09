package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	var N int
	var JMap = make(map[int32]int)
	for _, j := range J {
		JMap[j] = 1
	}
	for _, s := range S {
		if JMap[s] != 0 {
			N++
		}
	}
	return N
}

func main() {
	J := "aA"
	S := "aAAbbbb"
	N := numJewelsInStones(J, S)
	fmt.Println(N)
}
