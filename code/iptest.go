package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
)

func stringHex2Binary(strHex string) (string, error) {
	strByte := []byte(strHex)
	strBianry := ""
	for _, data := range strByte {
		str, err := strconv.ParseInt(string(data), 16, 10)
		if err != nil {
			return "", err
		}
		strBianry = fmt.Sprintf("%s%b", strBianry, str)
	}
	return strBianry, nil
}

func FullIPv6(ip net.IP) string {
	dst := make([]byte, hex.EncodedLen(len(ip)))
	_ = hex.Encode(dst, ip)
	tmpRet := string(dst[0:4]) + ":" +
		string(dst[4:8]) + ":" +
		string(dst[8:12]) + ":" +
		string(dst[12:16]) + ":" +
		string(dst[16:20]) + ":" +
		string(dst[20:24]) + ":" +
		string(dst[24:28]) + ":" +
		string(dst[28:])
	return tmpRet
}

func FullIPv6WithMask(ipFull string, mask string) (string, error) {
	maskB, err := stringHex2Binary(mask)
	if err != nil {
		return "", err
	}
	retIp := make([]byte, 0, len(maskB))
	OriIp := []byte(ipFull)
	ipk := 0
	var k int
	for k = 1; k <= len(maskB); k++ {
		if maskB[k-1] == '0' {
			break
		}
		if k%4 == 0 {
			if k != 0 && k%16 == 0 {
				retIp = append(retIp, OriIp[ipk])
				ipk++
			}
			retIp = append(retIp, OriIp[ipk])
			ipk++
		}
	}
	fmt.Println(k, ipk)
	if (k-1)%16 != 0 {
		return "", fmt.Errorf("ip format not support")
	}
	return fmt.Sprintf("%s:/%d", string(retIp), k-1), nil
}

func getIP(ipOri net.IP) (string, error) {
	ipV4 := ipOri.To4()
	if ipV4 != nil {
		return ipV4.String(), nil
	}
	ipV6 := ipOri.To16()
	if ipV6 != nil {
		return FullIPv6(ipV6), nil
	}
	return "", fmt.Errorf("ip %s invalid", ipOri.String())
}

var allowMask = map[string]map[int]bool{
	"ipv4": {8: true, 16: true, 24: true, 32: true},
	"ipv6": {64: true, 128: true},
}

func checkAndGetIP(ipOri string) (string, error) {
	ip := net.ParseIP(ipOri)
	if ip != nil {
		return getIP(ip)
	} else {
		ipSpefic, ipNet, err := net.ParseCIDR(ipOri)
		if err != nil {
			return "", err
		}
		speficMask, maxMask := ipNet.Mask.Size()
		if speficMask == maxMask {
			return getIP(ipSpefic)
		}
		if ipSpefic.To4() != nil {
			if !allowMask["ipv4"][speficMask] {
				fmt.Printf("ip  mask err %s\n", ipOri)
			}
			return ipNet.String(), nil
		}
		ipTmp := ipSpefic.To16()
		if ipTmp != nil {
			speficMasks := ipNet.Mask.String()
			if !allowMask["ipv6"][speficMask] {
				fmt.Printf("ip  mask err %s \n", ipOri)
			}
			ipv6Full, err := getIP(ipSpefic)
			ipv6, err := FullIPv6WithMask(ipv6Full, speficMasks)
			return ipv6, err

		}
		return ipNet.String(), nil
	}
	return "", fmt.Errorf("unknow error")
}

func main() {
	ipOri := os.Args[1]
	//retIp, err := checkAndGetIP(ipOri)
	//fmt.Println("output")
	//fmt.Println(retIp, err)
	{
		ip4 := ipOri
		ipv4 := net.ParseIP(ip4)
		fmt.Println(ipv4, ipv4.To4())
		fmt.Println(ipv4, ipv4.To16())
	}
}
