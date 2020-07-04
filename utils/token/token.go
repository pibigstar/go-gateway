package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/consts"
	"github/pibigstar/go-gateway/app/consts/code"
	"github/pibigstar/go-gateway/app/response"
	"github/pibigstar/go-gateway/utils"
	"github/pibigstar/go-gateway/utils/errx"
	"time"
)

const (
	// 加密的key值
	secretKey = "pibigstar"
	// token有效期,jwt默认key
	ClaimTokenEXP = "exp"
	// token使用的范围
	ClaimTokenScope = "web"
	ClaimTokenAdmin = "admin"

	// 用户信息
	ClaimTokenUserId = "userId"
)

// 生成token
func GenJwtToken(userInfo interface{}) string {
	claims := make(jwt.MapClaims)
	// 有效期
	claims[ClaimTokenEXP] = time.Now().Add(24 * time.Hour).Unix()
	claims[ClaimTokenUserId] = userInfo

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, _ := token.SignedString([]byte(secretKey))
	return s
}

// 检查token是否有效
func CheckJwtToken(tokenString string) bool {
	if tokenString == "" {
		return false
	}
	if err := CheckJwtTokenExpected(tokenString); err != nil {
		return false
	}
	return true
}

// 检查token是否过期
func CheckJwtTokenExpected(tokenString string) error {
	token, err := ParseJwtToken(tokenString)
	if err != nil {
		return err
	}
	return token.Claims.Valid()
}

// 解析token
func ParseJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("unexpected token claims")
		}
		return []byte(secretKey), nil
	})

	return token, err
}

func GetValueFromToken(t string) (interface{}, bool) {
	token, err := ParseJwtToken(t)
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if v, ok := claims[ClaimTokenUserId]; ok {
			return v, true
		}
	}
	return nil, false
}

// 从token中拿到用户信息
func GetUserInfoFromToken(tokenString string) (*response.AdminInfo, error) {
	if v, b := GetValueFromToken(tokenString); b {
		adminInfo := &response.AdminInfo{}
		if err := utils.MapToStruct(v, adminInfo); err == nil {
			return adminInfo, nil
		}
	}
	return nil, errx.New(code.Error_Not_Login)
}

// 从cookie中拿到用户信息
func GetUserInfoFromCookie(r *ghttp.Request) (*response.AdminInfo, error) {
	c, err := r.Request.Cookie("token")
	if err != nil {
		return nil, errx.New(code.Error_Not_Login)
	}
	return GetUserInfoFromToken(c.Value)
}

// 从session中拿到用户信息
func GetUserInfoFromSession(r *ghttp.Request) (*response.AdminInfo, error) {
	t, ok := r.Session.Get(consts.UserTokenSessionKey).(string)
	if !ok || t == "" {
		return nil, errx.New(code.Error_Not_Login)
	}
	return GetUserInfoFromToken(t)
}
