//get questions from problems.csv
package main

import (
	"encoding/csv"
	"log"
	"os"
)

func getQuestions() map[string]string {
	problemSet := make(map[string]string)

	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error while reading problems file", err)
	}

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		problemSet[record[0]] = record[1]
	}

	return problemSet
}
