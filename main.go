// quiz game with a built in timer
// main driver function with quiz loop
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	tlimit := 30
	if len(os.Args) > 1 {
		tlimit = getTimeClock(os.Args[1])
	}
	fmt.Println(tlimit)

	fmt.Print("Press Enter to begin quiz.\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	questions := getQuestions()
	score := quizLoop(questions, tlimit)
	fmt.Printf("Your score is: %d out of %d\n", score, len(questions))
}

// primary loop for the quiz
func quizLoop(questions map[string]string, tlimit int) int {
	var score int

	ansChan := make(chan string)

out:
	for question, ans := range questions {
		fmt.Print(question + "= ")
		go getInput(ansChan)

		select {
		case <-time.After(time.Duration(tlimit) * time.Second):
			fmt.Println()
			break out
		case input := <-ansChan:
			if input == ans {
				score++
			}
		}
	}
	return score
}

func getTimeClock(s string) int {
	tlimit, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return tlimit
}
