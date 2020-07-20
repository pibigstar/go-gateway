package errx

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/glog"
	"github/pibigstar/go-gateway/app/consts/code/locales"
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
	args []interface{}
}

func (e ErrorX) Error() string {
	msg := fmt.Sprintf("code: %d", e.code)
	if JsonMsg != nil {
		if format, ok := JsonMsg.Get(fmt.Sprintf("%d", e.code)).(string); ok {
			msg = format
			if len(e.args) > 0 {
				msg = fmt.Sprintf(format, e.args...)
			}
		}
	}
	return msg
}

func (e ErrorX) Code() int {
	return e.code.Code()
}

func New(code Coder, args ...interface{}) ErrorX {
	return ErrorX{
		code: code,
		args: args,
	}
}
