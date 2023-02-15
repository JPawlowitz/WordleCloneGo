package src

import "fmt"

var letters []string

func InitAlphabet() {
	for i := 97; i <= 122; i++ {
		letters = append(letters, string(i))
	}
}

func PrintRemainingLetters() {
	fmt.Print(letters[0])

	for i := 1; i < len(letters); i++ {
		fmt.Printf(", %s", letters[i])
	}

	fmt.Print("\n")
}
