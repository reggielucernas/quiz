package main

import (
	// "encoding/csv"
	"flag"
	"fmt"
	"os"
	// "strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	fmt.Println("filename:", *csvFilename)

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *csvFilename)
		os.Exit(1)
	}
	fmt.Println(*file)

}
