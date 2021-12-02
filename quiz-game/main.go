package main

import(
	"encoding/csv"
	"fmt"
	"os"
	"flag"
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
			ans: line[1],
		}
	}
	return set
}


func main() {
	// Setting the problem file
	filename := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
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

	// Asking the problems
	score := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.ques)
		var ans string
		fmt.Scanf("%s\n",&ans)

		// Validating ans and scoring
		if ans == problem.ans {
			score++ // increasing score
		}
	}
	
	fmt.Printf("You scored %d of %d.\n", score, len(problems))
}