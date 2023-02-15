package src

import "fmt"

func RunGame() {
	var guessedWords [5]string
	word := GetRandomWord()

	for try := 0; try < 5; try++ {
		guessedWords[try] = getInput(&guessedWords, &try)
		if CompareWords(word, guessedWords[try]) {
			fmt.Println("You won!")
			return
		}
	}

	fmt.Printf("You lost! \n The word was %s", word)
}

func getInput(pastWords *[5]string, numTries *int) string {
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
