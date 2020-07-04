package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/goinggo/mapstructure"
)

func GenSaltPassword(salt, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	return fmt.Sprintf("%x", s2.Sum(nil))
}

//MD5 md5加密
func MD5(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MapToStruct(m interface{}, dst interface{}) error {
	return mapstructure.Decode(m, dst)
}
