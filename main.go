package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var answer string
var correctAnswers int

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
		fmt.Println("There is an error", err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	for _, record := range records {
		fmt.Println("Please answer this question", record[0])
		fmt.Scan(&answer)
		if answer == record[1] {
			correctAnswers += 1
		}
	}
	results := fmt.Sprintf("You received a %d out of %d!", correctAnswers, len(records))
	fmt.Println(results)
}
