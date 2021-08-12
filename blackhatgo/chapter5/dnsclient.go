package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/miekg/dns"
)

type result struct {
	IPAddress string
	HostName  string
}

var wg sync.WaitGroup

///////////////////////////////////// FAN OUT CONCURRENCY PATTERN /////////////////////////////////////////

func main() {
	var flWorkerCount *int //flag.Int("c", 2, "workers to be user")
	var flWordlist *string //flag.String("w", "", "wordlist to be user")
	var flDomain *string

	domain := "microsoft.com"
	workerCount := 5
	wordlist := "wordlist"

	flWorkerCount = &workerCount
	flDomain = &domain //flag.String("d", "", "domain to be user")
	flWordlist = &wordlist

	flag.Parse()

	var results []result

	gather := make(chan []result)
	fqdns := make(chan string, *flWorkerCount)
	dnsServerAddr := "8.8.8.8:53"

	fh, err := os.Open(*flWordlist)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	// purpose of this go routine 
	// Launche the separte go routines i.e Workercount and do the job with each go routine
	go func() {
		wg.Add(*flWorkerCount) //To sync the number of go routines created 
		for i := 0; i < *flWorkerCount; i++ {
			go worker(gather, fqdns, dnsServerAddr) // launch the cpu intensive work n time n is number of threads
		}
		wg.Wait() // will wait for all go routines to complete their work which will be done at line 91
		close(gather) // will close the output channel
	}()

	go func() {
		for scanner.Scan() {
			val := scanner.Text()
			fqdns <- fmt.Sprintf("%s.%s", val, *flDomain) // will write continuously on the input channel
		}
		close(fqdns) //once the writing is complete close the channel
	}()

	for r := range gather { // loop over the output channel and store the result
		results = append(results, r...)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.HostName, r.IPAddress)
	}
	w.Flush()

}

func worker(gather chan []result, fqdns chan string, dnsServerAddr string) {

	for eachDnsReq := range fqdns { // continuosly reading from input channel which is being written at line 66
		results := lookup(eachDnsReq, dnsServerAddr) //my actual work
		if len(results) > 0 {
			gather <- results // continously writing to the output channel which is being read at line 71 by ranging over
		}
	}
	wg.Done() // once the input channel is close then close this go routine and will be called for each go routine working on this method
}

func lookup(fqdn string, dnsServerAddr string) []result {
	var results []result
	ips := lookupA(fqdn, dnsServerAddr)
	for _, ip := range ips {
		results = append(results, result{ip, dnsServerAddr})
	}
	return results
}

func lookupA(fqdn, dnsServerAddr string) []string {
	var ipaddress []string
	var msg dns.Msg
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		fmt.Println(err, ", ", fqdn)
		return nil
	}
	if len(in.Answer) < 1 {
		fmt.Println("No records")
	} else {
		for _, r := range in.Answer {
			fmt.Println(r)
			if a, ok := r.(*dns.A); ok {
				ipaddress = append(ipaddress, string(a.A))
			}
		}
	}
	return ipaddress
}
