package utils

import (
	"testing"
)

func TestToken(t *testing.T) {
	token, err := GenJwtToken("123456")
	if err != nil {
		t.Error(err)
	}

	t.Log("token:", token)

	isToken := CheckJwtToken(token)
	t.Log("isToken:", isToken)

	if uid, found := GetUserInfoFromToken(token); found {
		t.Log("用户id：", uid)
	}
}
