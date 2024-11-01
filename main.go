package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
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

// ReadCsv takes in a pointer to a open file and returns a 2D slice of CSV data
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

func main() {
	filePtr := flag.String("f", "problems.csv", "Path to CSV file")

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

	correct := 0
	incorrect := 0
	for i, problem := range problems {
		var answer string
		// Display problems
		fmt.Printf("Problem #%d: %s\n", i+1, problem[0])

		// Get answers
		fmt.Scanln(&answer)

		// Keep track of correct and incorrect answers
		if answer == problem[1] {
			correct++
		} else {
			incorrect++
		}
	}

	// Display # of correct/incorrect answers
	fmt.Printf("%d correct, %d incorrect\n", correct, incorrect)
}
