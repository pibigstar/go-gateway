package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

// 使用 httputil 包里面的ReverseProxy做反向代理

func main() {
	// 当访问 127.0.0.1:7000 时，会被反向代理到 127.0.0.1:7002
	uri, err := url.Parse("http://127.0.0.1:7002")
	if err != nil {
		panic(err)
	}
	// 返回一个带 director 的 ReverseProxy
	reverseProxy := httputil.NewSingleHostReverseProxy(uri)
	// 自定义的修改返回值函数
	reverseProxy.ModifyResponse = modifyResponse
	// 自定义错误处理函数
	reverseProxy.ErrorHandler = errorHandler

	if err := http.ListenAndServe(":7000", reverseProxy); err != nil {
		panic(err)
	}

}

func modifyResponse(resp *http.Response) error {
	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 添加一些新的返回值
	payload = []byte("My Response:" + string(payload))

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
	resp.ContentLength = int64(len(payload))
	resp.Header.Set("Content-Length", strconv.FormatInt(int64(len(payload)), 10))
	return nil
}

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, "ErrorHandler error:"+err.Error(), 500)
}
