## 小工具类

### Set
Set，与 java 中的 Set 类似，特点是在添加数组的时候，不允许出现重复的 string 元素
```go
	// add
	mySet := make(Set)
	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("apple") // 这个不会重复添加

	fmt.Println(mySet.Contains("apple"))  // 输出 true
	fmt.Println(mySet.Contains("orange")) // 输出 false

	// addAll
	stringSet := make(Set)
	stringSet.AddAll([]string{"apple", "banana", "apple", "orange"})

	fmt.Println("Set contains 'apple':", stringSet.Contains("apple"))   // 输出 "Set contains 'apple': true"
	fmt.Println("Set contains 'orange':", stringSet.Contains("orange")) // 输出 "Set contains 'orange': true"
	fmt.Println("Set contains 'grape':", stringSet.Contains("grape"))   // 输出 "Set contains 'grape': false"


	// Contains
	fmt.Println(stringSet.Contains("apple"))
	fmt.Println(stringSet.Contains("applePen"))

	// Remove
	stringSet.Remove("banana")
```

### AutoReaderBody2Byte
把 http 请求返回的 resp 的 body 读取转换成 byte 类型，如果出现乱码，可能是存在未解压的问题，该方法将自动帮你解压，避免出现乱码问题
```go
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, sendUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyByte, err := myutil.AutoReaderBody2Byte(resp)  // 划重点
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bodyByte))
```

### Unicode2String  
把 unicode 转为 string
```go
	unicode := `\u60a8\u4f3c\u4e4e\u5df2\u7ecf\u7b7e\u5230\u8fc7\u4e86..`
	str, err := Unicode2String(unicode)
	fmt.Println(str, err)
```

### StringToBytes
把 string 转为 byte[]
```go
	a := "abc"
	b := StringToBytes(a)
	fmt.Println(b)
```

### BytesToString
把 []byte 转为 string
```go
	a := []byte{'a', 'b', 'c'}
	b := BytesToString(a)
	fmt.Println(b)
```

### NewReadAll
读取 io.reader 流，并返回 []byte 类型数据
```go
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	buf := bufio.NewReader(file)
	result, _ := NewReadAll(buf)
	fmt.Println(string(result))
```

### InitProxy
设置代理，同时支持并自动识别 socks5 协议和 http 协议  
传入代理的格式是：  
socks5://localhost:8080  
http://localhost:8080  
```go
	proxyUrl := "socks5://localhost:1080"
	//proxyUrl := "http://localhost:1080"
	tr := &http.Transport{
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 500,
		MaxConnsPerHost:     500,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	// 为 tr 添加代理
	err := InitProxy(tr, proxyUrl)
	if err != nil {
		fmt.Println(err)
	}
	clinet := http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("GET", "https://baidu.com", nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := clinet.Do(req)
	if err != nil {
		log.Println(err)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
```

### ScanBodyContentLength2int
传一个数据包进去，给你返回数据包中 body 的长度  
这个方法适用于请求包中没有 Content-Length 参数，但是是 POST 请求，这个时候用来计算 body 长度的
```go
	str := `POST /mcp/pc/pcsearch HTTP/1.1
Host: ug.baidu.com
Cookie: BIDUPSID=DE728E5577FBF6CE396E7A9B6EB9FF7E; PSTM=1666859102; BAIDUID=DE728E5577FBF6CEDB3E4C42CE7D2961:FG=1; BAIDUID_BFESS=DE728E5577FBF6CEDB3E4C42CE7D2961:FG=1; ZFY=j7:Bn4Ni084v1YF9ehv2QDaQJ:AXEkMTRZXWfvL8qzAWA:C; H_PS_PSSID=38515_36547_38529_38469_38350_38368_38468_38486_37935_26350_38545; BA_HECTOR=ap8ha0852ka00l858l2ha0fe1i4jp8s1m; BDRCVFR[S4-dAuiWMmn]=I67x6TjHwwYf0; delPer=0; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598
Sec-Ch-Ua: "Chromium";v="103", ".Not/A)Brand";v="99"
Sec-Ch-Ua-Mobile: ?0
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36
Sec-Ch-Ua-Platform: "macOS"
Content-Type: application/json
Accept: */*
Origin: https://www.baidu.com
Sec-Fetch-Site: same-site
Sec-Fetch-Mode: cors
Sec-Fetch-Dest: empty
Referer: https://www.baidu.com/s?wd=2023%E5%B9%B4%E6%98%A5%E5%AD%A3%E4%B8%AD%E5%9B%BD%E5%85%83%E9%A6%96%E5%A4%96%E4%BA%A4%E7%BA%AA%E4%BA%8B&sa=fyb_n_homepage&rsv_dl=fyb_n_homepage&from=super&cl=3&tn=baidutop10&fr=top1000&rsv_idx=2&hisfilter=1
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Connection: close

{"invoke_info":{"pos_1":[{}],"pos_2":[{}],"pos_3":[{}]}}`
	a := ScanBodyContentLength2int([]byte(str))
	fmt.Println(a)
```

### Data2Request
给我一个 data 数据包，还你一个 request 和 err  
一般 http.ReadRequest 这个方法读取的数据包，有的时候是没办法直接用的，这个时候使用本方法，可以完美的修复了因为缺少 Content-Length 参数而无法使用的数据包
```go
	str := `POST /mcp/pc/pcsearch HTTP/1.1
Host: ug.baidu.com
Cookie: BIDUPSID=DE728E5577FBF6CE396E7A9B6EB9FF7E; PSTM=1666859102; BAIDUID=DE728E5577FBF6CEDB3E4C42CE7D2961:FG=1; BAIDUID_BFESS=DE728E5577FBF6CEDB3E4C42CE7D2961:FG=1; ZFY=j7:Bn4Ni084v1YF9ehv2QDaQJ:AXEkMTRZXWfvL8qzAWA:C; H_PS_PSSID=38515_36547_38529_38469_38350_38368_38468_38486_37935_26350_38545; BA_HECTOR=ap8ha0852ka00l858l2ha0fe1i4jp8s1m; BDRCVFR[S4-dAuiWMmn]=I67x6TjHwwYf0; delPer=0; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598
Sec-Ch-Ua: "Chromium";v="103", ".Not/A)Brand";v="99"
Sec-Ch-Ua-Mobile: ?0
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36
Sec-Ch-Ua-Platform: "macOS"
Content-Type: application/json
Accept: */*
Origin: https://www.baidu.com
Sec-Fetch-Site: same-site
Sec-Fetch-Mode: cors
Sec-Fetch-Dest: empty
Referer: https://www.baidu.com/s?wd=2023%E5%B9%B4%E6%98%A5%E5%AD%A3%E4%B8%AD%E5%9B%BD%E5%85%83%E9%A6%96%E5%A4%96%E4%BA%A4%E7%BA%AA%E4%BA%8B&sa=fyb_n_homepage&rsv_dl=fyb_n_homepage&from=super&cl=3&tn=baidutop10&fr=top1000&rsv_idx=2&hisfilter=1
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Connection: close

{"invoke_info":{"pos_1":[{}],"pos_2":[{}],"pos_3":[{}]}}`
	req, err := Data2Request("https://ug.baidu.com/mcp/pc/pcsearch", []byte(str))
	if err != nil {
		log.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	a, err := AutoReaderBody2Byte(resp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(a))
```

---
### v1.1.4
新增了两个方法
1. ScanBodyContentLength2int
2. Data2Request

### v1.1.3
修复了 AutoReaderBody2Byte 方法的 bug

### v1.1.2
新更新的版本处理了各个方法的空指针异常返回报错

...