package main

import "fmt"

//Check a string and return its bool slice
//FALSE = binary 0, TRUE = binary 1
func argumentToBinary(argument string) []bool {
	stringLength := len(argument)
	fmt.Print("Turning string into binary ")
	fmt.Println(stringLength)
	bits := make([]bool, stringLength)
	for position, char := range argument {
		if char == '0' {
			bits[position] = false
		} else if char == '1' {
			bits[position] = true
		}
	}
	return bits
}
