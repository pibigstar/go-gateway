package errx

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/glog"
	"github/pibigstar/go-gateway/app/const/code/locales"
	"strings"
)

var JsonMsg *gjson.Json

func init() {
	var err error
	JsonMsg, err = gjson.Load(locales.JsonPath())
	if err != nil {
		glog.Println("read json failed")
	}
}

type Coder interface {
	Code() int
}

type ErrorX struct {
	code Coder
	msg  string
}

func (e ErrorX) Error() string {
	if e.msg == "" && JsonMsg != nil {
		if msg, ok := JsonMsg.Get(fmt.Sprintf("%d", e.code)).(string); ok {
			e.msg = msg
		}
	}
	return e.msg
}

func (e ErrorX) Code() int {
	return e.code.Code()
}

func New(code Coder, msg ...string) ErrorX {
	e := ErrorX{
		code: code,
	}
	if len(msg) > 0 {
		e.msg = strings.Join(msg, ",")
	}
	return e
}
