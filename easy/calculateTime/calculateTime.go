package main

import (
	"fmt"
	"math"
)

func calculateTime(keyboard string, word string) int {
	var N int
	var keyboardMap = make(map[rune]int)
	for i, k := range keyboard {
		keyboardMap[k] = i
	}
	N = keyboardMap[rune(word[0])]
	for i := 0; i < len(word)-1; i++ {
		distance := keyboardMap[rune(word[i])] - keyboardMap[rune(word[i+1])]
		distance = int(math.Abs(float64(distance)))
		N += distance
	}
	return N
}

func main() {
	keyboard := "abcdefghijklmnopqrstuvwxyz"
	word := "cba"
	N := calculateTime(keyboard, word)
	fmt.Println(N)
}
