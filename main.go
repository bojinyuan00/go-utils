package main

import (
	`fmt`
	`github.com/bojinyuan00/go-utils/dnsip`
)

func main() {
	ipStrings, _ := dnsToIpTest("www.baidu.com")
	fmt.Println(ipStrings)
}

//事例-调用测试
func dnsToIpTest(dnsString string) (ipStrings []string, err error) {
	ipStrings, err = dnsip.DnsToIp(dnsString)
	if err != nil {
		return
	}
	return
}
