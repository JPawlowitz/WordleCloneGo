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

func CompareWords(word string, guess string) GuessResult {
	var result GuessResult
	var wrongLetters []string

	for i := 0; i < len(word); i++ {
		letter := IndexedChar{char: guess[i], index: i}

		if word[i] == guess[i] {
			result.correctPosition = append(result.correctPosition, letter)
			continue
		} else {
			for index, _ := range word {
				if guess[i] == word[index] {
					result.incorrectPosition = append(result.incorrectPosition, letter)
					continue
				}
			}
		}

		wrongLetters = append(wrongLetters, string(guess[i]))
	}

	RemoveLetters(wrongLetters)
	stringBuilder := strings.Builder{}

	for i := 0; i < len(guess); i++ {
		for _, correctPosition := range result.correctPosition {
			if guess[i] == correctPosition.char && i == correctPosition.index {
				stringBuilder.WriteString(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, string(guess[i])))
				break
			}

			stringBuilder.WriteString(string(guess[i]))
		}
	}

	fmt.Print(stringBuilder.String() + "\n")

	return result
}

type GuessResult struct {
	correctPosition   []IndexedChar
	incorrectPosition []IndexedChar
}

type IndexedChar struct {
	char  uint8
	index int
}
