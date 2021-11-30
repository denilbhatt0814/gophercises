package main

import(
	"encoding/csv"
	"fmt"
	"os"
)

type empData struct{
	Name 	string
	Age 	string
	City 	string
}

func main() {
	// Opening the file
	csvFile, err := os.Open("emp.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Succesfully opened CSV file...")
	defer csvFile.Close()

	// Creating the reader & reading all lines
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i, line := range csvLines {
		if i == 0{
			continue
		}
		emp := empData{
			Name: line[0],
			Age: line[1],
			City: line[2],
		}
		fmt.Println(emp.Name + " " + emp.Age + " " + emp.City)
	}
}