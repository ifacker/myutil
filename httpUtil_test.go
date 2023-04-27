package myutil

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestScanBodyContentLength2int(t *testing.T) {
	str := `POST /mcp/pc/pcsearch HTTP/1.1
Host: ug.baidu.com
Cookie: BIDUPSID=DE728E5577FBF6CE396E7A9B6EB9FF7E; PSTM=1666859102; BAIDUID=DE728E5577FBF6CEDB3E4C42CE7D2961:FG=1; BAIDUID_BFESS=DE728E5577FBF6CEDB3E4C42CE7D2961:FG=1; ZFY=j7:Bn4Ni084v1YF9ehv2QDaQJ:AXEkMTRZXWfvL8qzAWA:C; H_PS_PSSID=38515_36547_38529_38469_38350_38368_38468_38486_37935_26350_38545; BA_HECTOR=ap8ha0852ka00l858l2ha0fe1i4jp8s1m; BDRCVFR[S4-dAuiWMmn]=I67x6TjHwwYf0; delPer=0; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598
Content-Length: 56
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
}

func TestData2Request(t *testing.T) {

	proxyUrl := "http://localhost:8080"
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

	str := `POST /mcp/pc/pcsearch HTTP/1.1
Host: ug.baidu.com
Cookie: BIDUPSID=DE728E5577FBF6CE396E7A9B6EB9FF7E; PSTM=1666859102; BAIDUID=DE728E5577FBF6CEDB3E4C42CE7D2961:FG=1; BAIDUID_BFESS=DE728E5577FBF6CEDB3E4C42CE7D2961:FG=1; ZFY=j7:Bn4Ni084v1YF9ehv2QDaQJ:AXEkMTRZXWfvL8qzAWA:C; H_PS_PSSID=38515_36547_38529_38469_38350_38368_38468_38486_37935_26350_38545; BA_HECTOR=ap8ha0852ka00l858l2ha0fe1i4jp8s1m; BDRCVFR[S4-dAuiWMmn]=I67x6TjHwwYf0; delPer=0; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598
Content-Length: 56
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
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	a, err := AutoReaderBody2Byte(resp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(a))
}
