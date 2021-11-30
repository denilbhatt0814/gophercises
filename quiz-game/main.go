package main

import(
	"encoding/csv"
	"fmt"
	"os"
)

var(
	filename string = "problems.csv"
	score int = 0
)

func main() {
	// Opening the problems file
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	// Read the problems
	csvLines, err := csv.NewReader(csvFile).ReadAll()

	// Asking the problems
	qno := 0
	for _, problem := range csvLines {
		fmt.Printf(problem[0] + " = ")
		qno++ // increasingt the no. of Q count
		var ans string
		fmt.Scanf("%s\n",&ans)

		// Validating ans and scoring
		if ans == problem[1]{
			score++ // increasing score
		}
	}

	fmt.Printf("Score: %d/%d \n", score, qno)
}