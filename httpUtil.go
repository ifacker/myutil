package myutil

import (
	"compress/gzip"
	"io"
	"net/http"
)

// 把 http 请求返回的 resp 的 body 读取转换成 byte 类型，如果出现乱码，可能是存在未解压的问题，该方法将自动帮你解压，避免出现乱码问题
func AutoReaderBody2Byte(resp *http.Response) ([]byte, error) {
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
