package main

import "fmt"

func removeVowels(str string) string {
	var N string
	for _, s := range str {
		s := string(s)
		switch s {
		case "a", "e", "i", "o", "u":
		default:
			N += s
		}
	}
	return N
}

func main() {
	strInput := "leetcodeisacommunityforcoders"
	strOutput := removeVowels(strInput)
	fmt.Println(strOutput)
}
