package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	csvFile := flag.String("csv", "questionanswer.csv", "csv File for question and answer")
	flag.Parse()

	openFile, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file %s \n", *csvFile))
	}

	readFile := csv.NewReader(openFile)
	lines, err := readFile.ReadAll()
	if err != nil {
		fmt.Println("Failed to read provided csv")
	}

	//io
	questions := parseLine(lines)

	counter := 0
	for i, quest := range questions {
		fmt.Printf("Question Number #%d. %s:\n", i+1, quest.q)

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == quest.a {
			counter++
		}
	}

	if counter == 5 {
		fmt.Printf("Perfect!, You score %d out of %d", counter, len(questions))
	} else if counter == 4 {
		fmt.Printf("Not bad!, You score %d out of %d", counter, len(questions))
	} else {
		fmt.Printf("Idiot!, You score %d out of %d", counter, len(questions))
	}

}

type question struct {
	q string
	a string
}

// process
func parseLine(lines [][]string) []question {
	questionSlice := make([]question, len(lines))
	for i, line := range lines {
		questionSlice[i] = question{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return questionSlice
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
