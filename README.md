## 小工具类

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

---
### v1.1.2
新更新的版本处理了各个方法的空指针异常返回报错

...