package internet

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestInitProxy(t *testing.T) {
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
}
