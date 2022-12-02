package myutil

import (
	"context"
	"errors"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// 初始化 proxy 配置，应传入 Transport 和 proxyURL【proxyURL 的格式为：socks5://localhost:8080 或 http://localhost:8080】
func InitProxy(tr *http.Transport, proxyURL string) error {
	// 全部转小写，方便判断
	proxy := strings.ToLower(proxyURL)
	if strings.Contains(proxy, "socks5") {
		return setSocks5Proxy(tr, proxy)
	} else if strings.Contains(proxy, "https") {
		return errors.New("[!] 代理配置无效，本项目暂不支持\"https\"协议，您可以尝试使用\"http\"协议或\"socks5\"协议")
	} else if strings.Contains(proxy, "http") {
		return setHttpProxy(tr, proxy)
	}
	return nil
}

// 设置 socks5 代理
func setSocks5Proxy(tr *http.Transport, proxyURL string) error {
	socksParse, proxyErr := url.Parse(proxyURL)
	if proxyErr != nil {
		return proxyErr
	}
	dialer, err := proxy.FromURL(socksParse, proxy.Direct)
	if err != nil {
		return err
	}
	dc := dialer.(interface {
		DialContext(ctx context.Context, network, addr string) (net.Conn, error)
	})
	if proxyErr == nil {
		tr.DialContext = dc.DialContext
	}
	return nil
}

// 设置 http 代理
func setHttpProxy(tr *http.Transport, proxyURL string) error {
	proxyPares, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	(*tr).Proxy = http.ProxyURL(proxyPares)
	return nil
}
