package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func DomainHasRecord(d string) (bool, error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(d+".", dns.TypeA)

	in, _, err := c.Exchange(m, "114.114.114.114:53")
	if err != nil {
		fmt.Println(err)
	}
	if len(in.Answer) > 0 {
		fmt.Println(in.Answer)
		return true, nil
	}
	if err == nil {
		return false, nil
	}

	return false, nil
}

func checkDnsStatus(domain, cname string) bool {
	c := new(dns.Client)
	m := new(dns.Msg)
	// for test domain = "oa.chnau99999.com"
	m.SetQuestion(domain+".", dns.TypeA)
	in, _, err := c.Exchange(m, "172.16.16.16:53")
	if err != nil {
		return false
	}

	for _, r := range in.Answer {
		if r.Header().Rrtype == dns.TypeCNAME {
			fmt.Println(r.(*dns.CNAME).Target)
			if r.(*dns.CNAME).Target == cname+"." {
				return true
			}
		}
	}

	return false
}

func main() {
	d := os.Args[1]
	fmt.Println(checkDnsStatus(d, "test"))
	fmt.Println(DomainHasRecord(d))
}
