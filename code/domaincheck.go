package main

import (
	"fmt"
	"golang.org/x/net/publicsuffix"
	"strings"
)

func CheckIsToAddDomainOK(originNum int64, limit int32, domainsOrig []string, toAddDomain string) (string, error) {
	// first remove .
	d := strings.Trim(toAddDomain, ".") // 删除首尾的 "."
	fmt.Println("d:", d)
	fmt.Println(len(toAddDomain), len(d))
	dSLD, err := publicsuffix.EffectiveTLDPlusOne(d)
	fmt.Println(dSLD, err)
	if err != nil {
		return "", err
	}
	partsSLD := strings.Split(dSLD, ".")
	fmt.Println(partsSLD)
	partsD := strings.Split(d, ".")
	fmt.Println(partsD)
	return "", nil
}
func main() {
	//domainori := []string{"..www.jd.com.."}
	//CheckIsToAddDomainOK(1, 2, domainori, "www.kd.w.baidu.com.cn")
	fmt.Println(publicsuffix.EffectiveTLDPlusOne("*.*.com"))
	fmt.Println(publicsuffix.EffectiveTLDPlusOne("www.map.google.com"))

	fmt.Println(publicsuffix.EffectiveTLDPlusOne("map.google.co.uk"))
	fmt.Println(publicsuffix.EffectiveTLDPlusOne("www.book.amazon.co.uk"))
	fmt.Println(publicsuffix.EffectiveTLDPlusOne("www.book.amadzon.co.uk"))
	fmt.Println(publicsuffix.EffectiveTLDPlusOne("book.amazon.co.uk"))
	fmt.Println(publicsuffix.EffectiveTLDPlusOne("amazon.co.uk"))

	fmt.Println(publicsuffix.PublicSuffix("www.book.amazon.co.uk"))
	fmt.Println(publicsuffix.PublicSuffix("www.book.amadzon.co.uk"))
	fmt.Println(publicsuffix.PublicSuffix("book.amazon.co.uk"))
	fmt.Println(publicsuffix.PublicSuffix("amazon.co.uk"))
	fmt.Println(publicsuffix.EffectiveTLDPlusOne("a.m.jd.com"))
	fmt.Println(publicsuffix.EffectiveTLDPlusOne("http://www.jd.com"))
	{
		domain := "*.a.b.com"
		dTmp := strings.TrimPrefix(domain, "*.")
		idx := strings.LastIndex(host, dTmp)
		if idx == -1 {
			continue
		}
		if !strings.HasPrefix(dnproto.Domain, "*.") && idx != 0 {
			continue
		}
		if dnproto.Proto != schema {
			err = common.NewWafApiError(common.ErrParamsFormat, "", "", "this domain does not support "+schema+" proto", "该域名未配置"+schema+"协议")
			continue
		}
		return true, nil

	}
}
