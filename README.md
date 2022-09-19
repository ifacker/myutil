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