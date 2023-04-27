package myutil

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 把 http 请求返回的 resp 的 body 读取转换成 byte 类型，如果出现乱码，可能是存在未解压的问题，该方法将自动帮你解压，避免出现乱码问题
func AutoReaderBody2Byte(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, errors.New("resp 空指针异常")
	}
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		var err error
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		reader = resp.Body
	}
	body, err := NewReadAll(reader)
	if err != nil {
		return body, err
	}
	return body, nil
}

// 扫描 请求包中 Content-Length 的长度
func ScanBodyContentLength2int(data []byte) int {
	if data == nil {
		return 0
	}
	var contentLength = 0
	if bytes.Contains(data, []byte("\r\n\r\n")) {
		arr := bytes.SplitN(data, []byte("\r\n\r\n"), 2)
		if len(arr) > 1 {
			contentLength = len(arr[1])
		}
	} else if bytes.Contains(data, []byte("\n\n")) {
		arr := bytes.SplitN(data, []byte("\n\n"), 2)
		if len(arr) > 1 {
			contentLength = len(arr[1])
		}
	}
	return contentLength
}

// 给我一个 data 数据包，还你一个 request 和 err
func Data2Request(baseUrl string, data []byte) (*http.Request, error) {
	if data == nil {
		return nil, errors.New("data 数据为空")
	}
	contentLength := ScanBodyContentLength2int(data)
	req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(data)))
	if err != nil {
		return nil, err
	}
	req.RequestURI = ""
	req.ContentLength = int64(contentLength)
	fmt.Println(req.URL)
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	req.URL = u
	return req, nil
}
