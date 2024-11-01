package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", file.Name())
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	problems, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse file: %s\n", file.Name())
	}

	fmt.Println(problems)

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, problem[0])
		fmt.Printf("Answer: %s\n", problem[1])
	}
}
