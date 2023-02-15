package src

import "fmt"

func RunGame() {
	word := GetRandomWord()

	for tries := 0; tries <= 5; tries++ {
		guess := getInput()
		CompareWords(&word, &guess)
	}
}

func getInput() string {
	PrintRemainingLetters()
	fmt.Println("Please enter next word:")

	var userInput string
	for {
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("Wrong input, letters only!")
			continue
		}

		if len(userInput) == 5 {
			return userInput
		}

		fmt.Println("Wrong input, letters only!")
	}
}
