package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("Data/problem.csv")
	if err != nil {
		fmt.Println("unable to read file", err)
		return
	}

	counter := 0
	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if record == nil && err == io.EOF {
			fmt.Println("Quiz is up")
			fmt.Println("Your score is ", counter)
			return
		} else if err != nil {
			fmt.Println("Unable to read single csv record ", err)
			return
		}
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("%s = ", record[0])
		scanner.Scan()
		answer := scanner.Text()

		if answer == record[1] {
			counter++
		}
	}
}
