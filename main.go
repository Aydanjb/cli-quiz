package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func openFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %s\n", file.Name())
		return nil, err
	}

	return file, nil
}

func readCsv(file *os.File) ([][]string, error) {
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

	file, err := openFile("problems.csv")
	if err != nil {
		panic(err)
	}

	problems, err := readCsv(file)
	if err != nil {
		panic(err)
	}

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, problem[0])
		fmt.Printf("Answer: %s\n", problem[1])
	}
}
