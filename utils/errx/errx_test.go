package errx

import (
	"github/pibigstar/go-gateway/app/const/code"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(code.Error_Server_Error)
	t.Log(err)

	err = New(code.Error_Server_Error, "hello world")
	t.Log(err)
}
