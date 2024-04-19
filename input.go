package main

import (
	"fmt"
)

func getInput(InputChan chan string) {
	var userInput string
	fmt.Scanln(&userInput)
	InputChan <- userInput
}
