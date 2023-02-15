package main

import (
	"WordleCloneGo/src"
)

func main() {
	src.LoadWords()
	src.InitAlphabet()
	src.RunGame()
}
