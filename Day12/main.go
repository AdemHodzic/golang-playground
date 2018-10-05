package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   int
}

var (
	//	waitGroup sync.WaitGroup
	score int
)

func readProblems(result chan Problem) {

	//	defer waitGroup.Done()
	problemsFile, _ := os.Open("problems.csv")
	problemsReader := csv.NewReader(bufio.NewReader(problemsFile))

	var problems []Problem

	for {
		line, err := problemsReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		question := line[0]
		answer, _ := strconv.Atoi(line[1])

		problem := Problem{question, answer}
		problems = append(problems, problem)
	}

	for _, p := range problems {
		result <- p
		time.Sleep(time.Second * 1)
	}
	close(result)
	_ = problemsFile.Close()
}

func solveProblems(result chan Problem) {

	//	defer waitGroup.Done()

	scanner := bufio.NewScanner(os.Stdin)

	for problem := range result {
		fmt.Printf("%s=", problem.question)

		scanner.Scan()

		input := strings.Trim(scanner.Text(), " ")

		givenAnswer, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if givenAnswer == problem.answer {
			fmt.Printf("Correct!\n")
			score++
		} else {
			fmt.Printf("Incorrect!\n")
		}
	}

	return
}

func startTimer(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {

	problems := make(chan Problem, 12)

	//	waitGroup.Add(1)
	go readProblems(problems)
	//	waitGroup.Add(1)
	go solveProblems(problems)
	//	waitGroup.Wait()
	startTimer(5)
	fmt.Printf("\nFinal score: %d\n", score)
}
