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
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file.")
	}
	for i, record := range records {
		fmt.Printf("Problem #%d: %s = \n", i+1, record[0])
		fmt.Scan(&answer)
		if answer == record[1] {
			correctAnswers += 1
		}
	}
	results := fmt.Sprintf("You scored %d out of %d!", correctAnswers, len(records))
	fmt.Println(results)
}
