package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/const/code"
	"github/pibigstar/go-gateway/app/response"
	"github/pibigstar/go-gateway/utils/errx"
	"time"
)

const (
	// 加密的key值
	secretKey = "pibigstar"
	// token有效期,jwt默认key
	TokenClaimEXP = "exp"
	// token使用的范围
	TokenClaimScope = "web"
	TokenClaimAdmin = "admin"

	// 将用户userId存放到token中
	TokenClaimUserId = "userId"
)

// 生成token
func GenJwtToken(userInfo interface{}) string {
	claims := make(jwt.MapClaims)
	// 有效期
	claims[TokenClaimEXP] = time.Now().Add(24 * time.Hour).Unix()
	claims[TokenClaimUserId] = userInfo

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

// 从token中拿到用户信息
func GetUserInfoFromToken(tokenString string) (value interface{}, found bool) {
	token, err := ParseJwtToken(tokenString)
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if v, ok := claims[TokenClaimUserId]; ok {
			return v, true
		}
	}

	return nil, false
}

// 从token中拿到用户信息
func GetUserInfoFromCookie(r *ghttp.Request) (*response.AdminInfo, error) {
	c, err := r.Request.Cookie("token")
	if err != nil {
		return nil, errx.New(code.Error_Not_Login)
	}

	token, err := ParseJwtToken(c.Value)
	if err != nil {
		return nil, errx.New(code.Error_Token_Expired)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if v, ok := claims[TokenClaimUserId]; ok {
			if userInfo, ok := v.(*response.AdminInfo); ok {
				return userInfo, nil
			}
		}
	}

	return nil, errx.New(code.Error_Not_Login)
}
