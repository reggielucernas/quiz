package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "set quiz time limit in seconds")
	shuffle := flag.Bool("shuffle", false, "enable random shuffling of questions")

	flag.Parse()

	t1 := time.NewTimer(time.Duration(*limit) * time.Second)

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the provided CSV file.")
		os.Exit(1)
	}

	problems := parseLines(lines)

	if *shuffle {
		shuffleLines(problems)
	}

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		select {
		case <-t1.C:
			fmt.Println("Time's up!")
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		default:
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Println("Finished!")
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func shuffleLines(problems []problem) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(problems) > 0 {
		n := len(problems)
		randIndex := r.Intn(n)
		problems[n-1], problems[randIndex] = problems[randIndex], problems[n-1]
		problems = problems[:n-1]
	}
}
