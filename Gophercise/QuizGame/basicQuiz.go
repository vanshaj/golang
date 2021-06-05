package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var filePath *string
	var limit *int
	filePath = flag.String("file", "Data/problem.csv", "Quiz file path")
	limit = flag.Int("limit", 30, "Time Limit for quiz")
	flag.Parse()
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("unable to read file", err)
		return
	}
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	counter := 0
	scannerR := bufio.NewScanner(os.Stdin)
	csvReader := csv.NewReader(file)
	answerCh := make(chan string) // to write the answers read from stdin to this channel

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
		go func() {
			scannerR.Scan() // blocks the code thats why inside a goroutine so that timer can continue and update if the time is up
			answer := scannerR.Text()
			answerCh <- answer
		}()
		select {
		/*
			 select case is beautiful because we can check on 2 different channel if info is on one channel then do something if on another something else
			This blocks the code till any info is on any channel
			1. Till the user doesnot write the answer it will block case <- answerCh
			2. Till the time is not up it will block case <- timer.C
			So code will not move forward till info is avail on any channel(REMEMBER CHANNEL BLOCKS)
		*/
		case <-timer.C: // when timer sends info to the timer channel then games up
			fmt.Println("Time is up")
			fmt.Println("Your score is ", counter)
			return
		case eachAnswer := <-answerCh: // as soon as answer is being send by the user it continues and moves to next question
			if eachAnswer == record[1] {
				counter++

			}
		}
	}
}
