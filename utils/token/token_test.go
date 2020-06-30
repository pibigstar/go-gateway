package token

import (
	"testing"
)

func TestToken(t *testing.T) {
	token := GenJwtToken("123456")

	t.Log("token:", token)

	isToken := CheckJwtToken(token)
	t.Log("isToken:", isToken)

	if uid, found := GetUserInfoFromToken(token); found {
		t.Log("用户id：", uid)
	}
}
