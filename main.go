package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

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
	timer := time.NewTimer(45 * time.Second)
problemloop: //a label that labels the for loop.
	for i, record := range records {
		fmt.Printf("Problem #%d: %s = \n", i+1, record[0])
		answerCh := make(chan string) //creates an answer channel
		go func() {
			var answer string
			fmt.Scan(&answer)
			answerCh <- answer // when we get an answer, we are sending it to the answer channel
		}()
		select {
		case <-timer.C: //timer.C blocks code until the timer is out and sends a message of completion on channel C.
			// ^^^ if there is a message from timer.C channel, then go through this case.
			fmt.Println("You ran out of time!")
			break problemloop // breaks this specific loop.
		case answer := <-answerCh: // if we get an answer from the answer channel
			if answer == record[1] {
				correctAnswers++
			}
		}
	}
	results := fmt.Sprintf("You scored %d out of %d!", correctAnswers, len(records))
	fmt.Println(results)
}
