## 小工具类

### 文件读写操作
[查看 fileIO_test.go](fileIO_test.go)

### 数组去重
[查看 element_test.go](element_test.go)

### Set
Set，与 java 中的 Set 类似，特点是在添加数组的时候，不允许出现重复的 string 元素。  
[查看 set_test.go](set_test.go)

### 字符转换
把 unicode 转为 string。
#### Unicode2String   
#### StringToBytes 
#### BytesToString
[查看 codeUtil_test.go](codeUtil_test.go)

### NewReadAll
读取 io.reader 流，并返回 []byte 类型数据。  
[查看 newIO_test.go](newIO_test.go)

### 数据包处理
#### ScanBodyContentLength2int
传一个数据包进去，给你返回数据包中 body 的长度  
这个方法适用于请求包中没有 Content-Length 参数，但是是 POST 请求，这个时候用来计算 body 长度的。  
[查看 httpUtil_test.go](httpUtil_test.go)

#### Data2Request
给我一个 data 数据包，还你一个 request 和 err  
一般 http.ReadRequest 这个方法读取的数据包，有的时候是没办法直接用的，这个时候使用本方法，可以完美的修复了因为缺少 Content-Length 参数而无法使用的数据包。  
[查看 httpUtil_test.go](httpUtil_test.go)

#### AutoReaderBody2Byte
把 http 请求返回的 resp 的 body 读取转换成 byte 类型，如果出现乱码，可能是存在未解压的问题，该方法将自动帮你解压，避免出现乱码问题。  
[查看 httpUtil_test.go](httpUtil_test.go)  

### InitProxy
设置代理，同时支持并自动识别 socks5 协议和 http 协议。  
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
### v1.1.9
新增了文件操作的一些方法  
详细请看`fileIO_test.go`文件

### v1.1.4
新增了两个方法
1. ScanBodyContentLength2int
2. Data2Request

### v1.1.3
修复了 AutoReaderBody2Byte 方法的 bug

### v1.1.2
新更新的版本处理了各个方法的空指针异常返回报错

...