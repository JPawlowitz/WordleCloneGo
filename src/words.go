package src

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var wordMap map[int]string

func LoadWords() {
	wordMap = make(map[int]string)

	textFile, err := os.Open("/Users/jan/Documents/Code/GolandProjects/WordleCloneGo/words_alpha.txt")
	if err != nil {
		panic("Error loading word file!")
	}

	defer textFile.Close()

	scanner := bufio.NewScanner(textFile)
	index := 0

	for scanner.Scan() {
		word := scanner.Text()

		if len(word) == 5 {
			wordMap[index] = word
			index++
		}
	}
}

func GetRandomWord() string {
	randomIndex := rand.Intn(len(wordMap))
	return wordMap[randomIndex]
}

func CompareWords(word string, guess string) bool {
	correctPositions := 0
	stringBuilder := strings.Builder{}

	for i := 0; i < len(word); i++ {
		if word[i] == guess[i] {
			stringBuilder.WriteString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, string(guess[i])))
			correctPositions++
			continue
		} else {
			isIncorrect := true
			for index := range word {
				if guess[i] == word[index] {
					stringBuilder.WriteString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 90, string(guess[i])))
					isIncorrect = false
					break
				}
			}

			if isIncorrect {
				stringBuilder.WriteString(string(guess[i]))
			}
		}
	}

	fmt.Print(stringBuilder.String() + "\n")

	return correctPositions == len(word)
}
