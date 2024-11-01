# go-utils
go相关工具包

# 拉取命令
## 常规
```go get github.com/bojinyuan00/go-utils```
## 默认
```go get github.com/bojinyuan00/go-utils@latest```
## 指定版本
```go get github.com/bojinyuan00/go-utils@v1.0.0```

# 功能划分
## dns domain ip transform
dnsip ==>实现域名转ip、ip转化域名相关功能<br>
	dnsip.DnsToIp("music.daydaylove.top")<br>
	dnsip.DnsToIp("http://www.baidu.com")<br>
	dnsip.DnsToIp("http://www.baidu.com:8011")<br>
  	dnsip.DnsToIp("http://www.baidu.com:8011/admin/users")<br>
  	dnsip.IpToDns("127.0.0.1")

## comfunc 公用的封装方法，独立做包使用/后期其它功能模块调用
判断字符串是否在切片数组内 comfunc.StringInSlice(a string, list []string) bool
