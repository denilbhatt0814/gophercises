package main

import(
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"time"
	"strings"
)

type problem struct {
	ques string
	ans	 string
}

func parseProblem(lines [][]string) []problem {
	set := make([]problem, len(lines))
	for i, line := range lines{
		set[i] = problem{
			ques: line[0],
			ans: strings.TrimSpace(line[1]),
		}
	}
	return set
}


func main() {
	// Setting the problem file
	filename := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	flag.Parse()

	// Opening the problems file
	csvFile, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	// Read the problems
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	problems := parseProblem(csvLines)

	// Declaring & initailzing a timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// Asking the problems
	score := 0
	t := time.Now()
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.ques)

		// Making a channel to recieve ans from diff. goroutine
		answerChan := make(chan string)
		go func() {
			/*This is go routine so that scanf doesn't 
				block the main goroutine which has a timer running*/
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select{
			/* This select waits for either of response (since no default declared)*/
		case <- timer.C:
			/* If the timer ends*/
			fmt.Printf("\nTime is up!!")
			fmt.Printf("\nYou scored %d of %d.\n", score, len(problems))
			fmt.Printf("You took %s\n", time.Since(t))
			return
		case answer := <- answerChan:
			/* If ans submited before timer ends */
			// Validating ans and scoring
			if answer == problem.ans {
				score++ // increasing score
			}				
		}
	}
	
	fmt.Printf("You scored %d of %d.\n", score, len(problems))
	fmt.Printf("You took %s\n", time.Since(t))
}