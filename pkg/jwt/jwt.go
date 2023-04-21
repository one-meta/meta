package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	// Domain   string `json:"domain"`
	// Project  string `json:"project"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// CreateJWT 创建
func CreateJWT(project, username, key string, expire int, method string) (string, error) {
	var hmac jwt.SigningMethod
	switch method {
	case "HS384":
		hmac = jwt.SigningMethodHS384
	case "HS512":
		hmac = jwt.SigningMethodHS512
	default:
		// "HS256"
		hmac = jwt.SigningMethodHS256
	}
	now := time.Now()
	// Create the Claims
	claims := &CustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * time.Duration(expire))),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(hmac, claims)
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return ss, nil
}

// ParseJWT 解析
func ParseJWT(tokenStr, key string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if tokenStr == "" {
			return nil, fmt.Errorf("token not found")
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// 销毁，未过期，则丢到黑名单缓存中
// 检查token是否销毁
