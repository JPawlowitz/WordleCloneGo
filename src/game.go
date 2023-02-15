package src

import "fmt"

func RunGame() {
	fmt.Printf("\x1b[%dm%s\x1b[0m", 34, "Correct letter, correct position! \n")
	fmt.Printf("\x1b[%dm%s\x1b[0m", 90, "Correct letter, false position! \n")

	word := GetRandomWord()

	for try := 0; try < 5; try++ {
		guess := getInput(&try)
		if CompareWords(word, guess) {
			fmt.Println("You won!")
			return
		}
	}

	fmt.Printf("You lost! \n The word was %s", word)
}

func getInput(numTries *int) string {
	fmt.Printf("Try %d/5 \n", *numTries)
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
