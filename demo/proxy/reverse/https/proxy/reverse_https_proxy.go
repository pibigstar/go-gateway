package proxy

import (
	"crypto/tls"
	"github/pibigstar/go-gateway/demo/balance"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, //连接超时
		KeepAlive: 30 * time.Second, //长连接超时时间
	}).DialContext,
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 跳过证书验证
	//TLSClientConfig: func() *tls.Config {
	//	pool := x509.NewCertPool()
	//	caCertPath := ssl.Path("ca.crt")
	//	caCrt, _ := ioutil.ReadFile(caCertPath)
	//	pool.AppendCertsFromPEM(caCrt)
	//	return &tls.Config{RootCAs: pool}
	//}(),
	MaxIdleConns:          100,              //最大空闲连接
	IdleConnTimeout:       90 * time.Second, //空闲超时时间
	TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
	ExpectContinueTimeout: 1 * time.Second,  //100-continue 超时时间
}

func NewMultipleHostsReverseProxy(balance balance.Balance) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		targetURL, err := balance.Get()
		if err != nil {
			panic(err)
		}
		target, err := url.Parse(targetURL)
		if err != nil {
			panic(err)
		}

		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "user-agent")
		}
	}
	if err := http2.ConfigureTransport(transport); err != nil {
		panic(err)
	}
	return &httputil.ReverseProxy{Director: director, Transport: transport}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
