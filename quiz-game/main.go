package main

import(
	"encoding/csv"
	"fmt"
	"os"
	"flag"
)

var(
	score int = 0
)

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

	// Asking the problems
	qno := 0
	for _, problem := range csvLines {
		qno++ // increasingt the no. of Q count
		fmt.Printf("Problem #%d: %s = ", qno,problem[0])
		var ans string
		fmt.Scanf("%s\n",&ans)

		// Validating ans and scoring
		if ans == problem[1]{
			score++ // increasing score
		}
	}

	fmt.Printf("You scored %d of %d.\n", score, qno)
}