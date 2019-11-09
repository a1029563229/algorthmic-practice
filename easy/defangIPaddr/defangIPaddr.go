package main

import "fmt"

func defangIPaddr(address string) string {
	var invalidAddress string
	for _, s := range address {
		s := string(s)
		switch s {
		case ".":
			invalidAddress += "[.]"
		default:
			invalidAddress += s
		}
	}

	return invalidAddress
}

func main() {
	address := "1.1.1.1"
	invalidAddress := defangIPaddr(address)
	fmt.Println(invalidAddress)
}
