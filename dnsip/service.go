package dnsip

import (
	`errors`
	`fmt`
	`github.com/bojinyuan00/go-utils/comfunc`
	`net`
	`net/url`
)

//url拆飞解析
func urlToIp(dnsString string) (ipStrings []string, err error) {
	//解析url
	u, err := url.Parse(dnsString)
	if err != nil {
		//fmt.Println("url解析错误，请检查后重新输入:", err.Error())
		return nil, errors.New("url解析错误，请检查后重新输入:" + err.Error())
	}

	//url=》协议类型判断
	requestType := u.Scheme //协议类型
	types := []string{"http", "https", "ftp", "ftps"}
	found := comfunc.StringInSlice(requestType, types)
	if !found {
		//fmt.Println("url协议类型错误，当前仅支持", types)
		return nil, errors.New("url协议类型错误，当前仅支持\"http\", \"https\", \"ftp\", \"ftps\"")
	}

	host := u.Hostname() //域名获取 =》后续用作ip地址解析
	port := u.Port()     //端口获取

	ips, err := domainToIp(host)
	if err != nil {
		return nil, err
	}

	//组装url完整参数
	for _, c := range ips {
		ipNew := requestType + "://" + c
		if len(port) == 0 {
			switch requestType {
			case "http":
				port = "80"
			case "https":
				port = "443"
			case "ftp":
				port = "21"
			case "ftps":
				port = "21"
			default:
				port = "-1"
			}
		}
		ipNew += ":" + port
		ipStrings = append(ipStrings, ipNew)
	}

	//fmt.Println(ipStrings)

	return ipStrings, nil
}

//域名解析
func domainToIp(dnsString string) (ipStrings []string, err error) {
	// 解析cname别名
	cname, err := net.LookupCNAME(dnsString)
	if err != nil {
		//fmt.Println("域名解析错误：", err.Error())
		return nil, errors.New("域名别名解析错误：" + err.Error())
	}

	// 解析ip地址
	ipStrings, err = net.LookupHost(dnsString)
	if err != nil {
		//fmt.Println("域名解析错误：", err.Error())
		return nil, errors.New("域名解析错误：" + err.Error())
	}

	// 对域名解析进行控制判断
	// 有些域名通常会先使用cname解析到一个别名上，然后再解析到实际的ip地址上
	switch {
	case cname != "":
		if len(ipStrings) != 0 {
			for _, n := range ipStrings {
				fmt.Println("cname域名匹配的ip地址：", n)
			}
		}
	case len(ipStrings) != 0:
		for _, n := range ipStrings {
			fmt.Println("域名匹配的ip地址：", n)
		}
	default:
		fmt.Println(cname, ipStrings)
	}

	return ipStrings, nil
}
