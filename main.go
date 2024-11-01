package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// OpenFile takes in a path to a file and returns a pointer to the open file
func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", file.Name())
		return nil, err
	}

	return file, nil
}

// ReadCsv takes in a pointer to an open file and returns a 2D slice of CSV data
func ReadCsv(file *os.File) ([][]string, error) {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse file: %s\n", file.Name())
		return nil, err
	}
	defer file.Close()

	return data, nil
}

func StartTimer() {

}

func main() {
	filePtr := flag.String("f", "problems.csv", "Path to CSV file")
	shufflePtr := flag.Bool("s", false, "Shuffle output")
	timeLimit := flag.Int("l", 10, "The time limit fot the quiz")

	flag.Parse()

	// Open quiz problems
	file, err := OpenFile(*filePtr)
	if err != nil {
		panic(err)
	}

	// Parse CSV
	problems, err := ReadCsv(file)
	if err != nil {
		panic(err)
	}

	// Shuffle problems
	if *shufflePtr {
		for i := range problems {
			j := rand.Intn(i + 1)
			problems[i], problems[j] = problems[j], problems[i]
		}
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	incorrect := 0
	for i, problem := range problems {
		// Display problems
		fmt.Printf("Problem #%d: %s\n", i+1, problem[0])

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answer = strings.TrimSpace(answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time's up!")
			fmt.Printf("%d correct, %d incorrect\n", correct, incorrect)
			return
		case answer := <-answerCh:
			// Keep track of correct and incorrect answers
			if answer == problem[1] {
				correct++
			} else {
				incorrect++
			}
		}
	}

	// Display # of correct/incorrect answers
	fmt.Printf("%d correct, %d incorrect\n", correct, incorrect)
}
