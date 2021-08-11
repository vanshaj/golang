package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/miekg/dns"
)

type result struct {
	IPAddress string
	HostName  string
}

func main() {
	var flWorkerCount *int //flag.Int("c", 2, "workers to be user")
	var flWordlist *string //flag.String("w", "", "wordlist to be user")
	var flDomain *string

	domain := "microsoft.com"
	workerCount := 2
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

	for i := 0; i < *flWorkerCount; i++ {
		go worker(gather, fqdns, dnsServerAddr)
	}

	for scanner.Scan() {
		val := scanner.Text()
		fqdns <- fmt.Sprintf("%s.%s", val, *flDomain)
	}

	go func() {
		for r := range gather {
			results = append(results, r...)
		}
	}()

	close(fqdns)
	close(gather)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.HostName, r.IPAddress)
	}
	w.Flush()

}

func worker(gather chan []result, fqdns chan string, dnsServerAddr string) {
	for eachDnsReq := range fqdns {
		results := lookup(eachDnsReq, dnsServerAddr)
		if len(results) > 0 {
			gather <- results
		}
	}

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
