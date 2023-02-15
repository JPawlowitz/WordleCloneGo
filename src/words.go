package src

import (
	"bufio"
	"math/rand"
	"os"
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

func CompareWords(word *string, guess *string) {

}
