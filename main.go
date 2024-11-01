package main

import (
	`fmt`
	`github.com/bojinyuan00/go-utils/dnsip`
)

func main() {
	ipStrings, err := dnsToIpTest("123344s")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("解析后的信息：", ipStrings)
}

//事例-调用测试(url、域名解析为ip地址)
func dnsToIpTest(dnsString string) (ipStrings []string, err error) {
	ipStrings, err = dnsip.DnsToIp(dnsString)
	if err != nil {
		return
	}
	return
}
