package errx

import (
	"github/pibigstar/go-gateway/app/consts/code"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(code.Error_Server_Error)
	t.Log(err)
}
