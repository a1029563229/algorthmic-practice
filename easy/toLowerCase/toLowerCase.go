package main

import "fmt"

func toLowerCase(str string) string {
	lowerS := ""
	for _, s := range str {
		if s >= 65 && s <= 90 {
			s += 32
		}
		lowerS += string(s)
	}
	return lowerS
}

func main() {
	s := toLowerCase("Hello")
	fmt.Println(s)
}
