package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var filePath *string
	filePath = flag.String("file", "Data/problem.csv", "Quiz file path")
	flag.Parse()
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("unable to read file", err)
		return
	}

	counter := 0
	scannerR := bufio.NewScanner(os.Stdin)
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

		fmt.Printf("%s = ", record[0])
		scannerR.Scan()
		answer := scannerR.Text()

		if answer == record[1] {
			counter++
		}
	}
}
