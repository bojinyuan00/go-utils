package dnsip

import (
	`fmt`
	`net`
	`os`
)

// Test 包测试
func Test() {
	fmt.Println("欢迎使用dns-ip相关转化工具包")
}

// DnsToIp dns-域名转ip
func DnsToIp(dnsString string) (ipStrings []string, err error) {
	// 解析cname
	cname, _ := net.LookupCNAME(dnsString)

	// 解析ip地址
	ipStrings, err = net.LookupHost(dnsString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	// 对域名解析进行控制判断
	// 有些域名通常会先使用cname解析到一个别名上，然后再解析到实际的ip地址上
	switch {
	case cname != "":
		fmt.Println("cname:", cname)
		if len(ipStrings) != 0 {
			fmt.Println("vips:")
			for _, n := range ipStrings {
				fmt.Fprintf(os.Stdout, "%s\n", n)
			}
		}
	case len(ipStrings) != 0:
		for _, n := range ipStrings {
			fmt.Fprintf(os.Stdout, "%s\n", n)
		}
	default:
		fmt.Println(cname, ipStrings)
	}
	return ipStrings, nil
}

// IpToDns dns-ip转化域名
func IpToDns(ipString string) (dnsStrings []string, err error) {
	// 反向解析(主机必须得能解析到地址)
	dnsStrings, _ = net.LookupAddr(ipString)
	fmt.Println("dns+1:", dnsStrings)
	return dnsStrings, nil
}
