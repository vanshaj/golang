package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {
	fmt.Println("vim-go")
	portS := make(chan int, 100)
	results := make(chan int)

	for i := 1; i < cap(portS); i++ {
		go startScan(portS, results)
	}

	go func() {
		for i := 1; i < 1023; i++ {
			portS <- i
		}
	}()
	var openPorts []int
	for i := 1; i < 1023; i++ {
		p := <-results
		if p != 0 {
			openPorts = append(openPorts, p)
		}
	}
	//close(portS)
	//close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}

func startScan(portS, results chan int) {
	for port := range portS {
		location := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", location)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}
