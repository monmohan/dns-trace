package main

import (
	"fmt"
	"net"
	"os"

	"github.com/miekg/dns"
)

var ErrNoAnswer = fmt.Errorf("No Answer Section")
var ErrNoAuthority = fmt.Errorf("No Authority Section")

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: $ go run trace.go <domain>")
		os.Exit(1)
	}
	domain := os.Args[1]
	ip := runQuery(domain)
	fmt.Printf("\nDomain %s resolved to IP(V4) %s\n", domain, ip.String())

}

func runQuery(domain string) (ipv4 net.IP) {
	rootServer := "199.9.14.201:53"
	c := new(dns.Client)
	return runQueryWithServer(c, domain, rootServer)

}

func runQueryWithServer(c *dns.Client, domain string, server string) (ipv4 net.IP) {
	fmt.Printf("\nResolving domain -> IP query for domain=%s with nameserver=%s\n", domain, server)
	in, _, err := c.Exchange(buildQuery(domain), server)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	ip := parseAnswer(in)

	if ip == nil {
		ip = parseAuthority(in) // find the nameserver IP
		if ip != nil {
			return runQueryWithServer(c, domain, ip.String()+":53")
		}

	}
	return ip
}
func parseAnswer(in *dns.Msg) (ip net.IP) {
	for _, e := range in.Answer {
		if d, ok := e.(*dns.A); ok {
			fmt.Printf("Go A record %s\n", d.A.String())
			return d.A
		}
		if cname, ok := e.(*dns.CNAME); ok {
			fmt.Printf("Looking for cname %s\n", cname.Target)
			return runQuery(cname.Target)

		}
	}
	return nil
}

func parseAuthority(in *dns.Msg) (ip net.IP) {
	for _, ns := range in.Ns {
		if d, ok := ns.(*dns.NS); ok {
			fmt.Printf("NS = %s ", d.Ns)
			//Look for IP in Additional Section
			for _, e := range in.Extra {
				if d, ok := e.(*dns.A); ok {
					fmt.Printf(" NS IP = %s", d.A.String())
					return d.A

				}
			}
			fmt.Printf("\nNS IP not found, Run Query for: %s ", d.Ns)
			return runQuery(d.Ns)

		}

	}
	return nil
}

func buildQuery(domain string) (question *dns.Msg) {
	m1 := new(dns.Msg)
	m1.Id = dns.Id()
	m1.RecursionDesired = true
	m1.Question = make([]dns.Question, 1)
	m1.Question[0] = dns.Question{domain, dns.TypeA, dns.ClassINET}
	return m1
}
