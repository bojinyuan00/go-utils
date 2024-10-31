package dnsip

import (
	`fmt`
	`net`
	`net/url`
)

// Test 包测试
func Test() {
	fmt.Println("欢迎使用dns-ip相关转化工具包")
}

// DnsToIp dns-domain-ip
func DnsToIp(dnsString string) (ipStrings []string, err error) {
	//1、先判断当前传值是否为url
	_, err = url.ParseRequestURI(dnsString)
	if err != nil { //非url直接按照域名字符串解析
		fmt.Println("不是url，按照域名字符串解析：", err.Error())
		return domainToIp(dnsString)
	}

	//2、url时，则获取相关协议类型、域名地址、端口信息、其他信息
	return urlToIp(dnsString)
}

// IpToDns dns-ip-domain
func IpToDns(ipString string) (dnsStrings []string, err error) {
	// 反向解析(主机必须得能解析到地址)
	dnsStrings, _ = net.LookupAddr(ipString)
	fmt.Println("dns+1:", dnsStrings)
	return dnsStrings, nil
}
