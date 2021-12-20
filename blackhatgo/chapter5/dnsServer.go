package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

var records map[string]string

func parse(filename string) (map[string]string, error) {
	records := make(map[string]string)
	fh, err := os.Open(filename)
	if err != nil {
		return records, err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ",", 2)
		if len(parts) < 2 {
			return records, fmt.Errorf("%s is not a valid line", line)
		}
		records[parts[0]] = parts[1]
	}
	return records, scanner.Err()

}

func dnsRequestHandler(w dns.ResponseWriter, req *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(req)
	if len(req.Question) < 1 {
		dns.HandleFailed(w, req)
		return
	}
	name := strings.TrimSuffix(req.Question[0].Name, ".")
	parts := strings.Split(name, ".")

	if len(parts) > 1 {
		name = strings.Join(parts[len(parts)-2:], ".")
	}
	match, ok := records[name]
	name = fmt.Sprintf("%v.", name)
	if !ok {
		dns.HandleFailed(w, req)
		return
	}
	msg.Authoritative = true
	msg.Answer = append(msg.Answer, &dns.A{
		Hdr: dns.RR_Header{
			Name:   name,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    60,
		},
		A: net.ParseIP(match),
	})
	// if err != nil {
	// 	dns.HandleFailed(w, req)
	// 	return
	// }
	if err := w.WriteMsg(&msg); err != nil {
		fmt.Println("Handle failed because of ", err)
		dns.HandleFailed(w, req)
		return
	}
}

func main() {
	records, _ = parse("proxy.config")
	dns.HandleFunc(".", dnsRequestHandler)
	log.Fatal(dns.ListenAndServe(":53", "udp", nil))
}
