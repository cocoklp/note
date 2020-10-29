package main

import (
	"fmt"
	"github.com/miekg/dns"
)

func main() {
	domain := "bds.jdcloud.com"
	cname := "bds.jdcloud.com-cb1946939e5d.jdcloudwaf.com"
	fmt.Println(doesDomainCnameMatch(domain, cname))
}

func doesDomainCnameMatch(domain, cname string) bool {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(domain+".", dns.TypeA)

	in, _, err := c.Exchange(m, "114.114.114.114:53")
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
