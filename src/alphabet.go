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

func RemoveLetters(lettersToRemove []string) {
	for i := 0; i < len(lettersToRemove); i++ {
		for index, letter := range letters {
			if letter == lettersToRemove[i] {
				letters = append(letters[:index], letters[index+1])
				break
			}
		}
	}
}
