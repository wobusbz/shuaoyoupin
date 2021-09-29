package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SIGNED_KEY = "wuhuaroua"
	ISSUER     = "wobusbzs"
	EXPIRESAT  = 7 * 86400
)

type AuthToken struct {
	UserInter interface{} // 传入用户私人信息
	jwt.StandardClaims
}

func NewAuthToken() *AuthToken {
	auth := &AuthToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + EXPIRESAT, // 过期时间,
			IssuedAt:  time.Now().Unix(),             // 签发时间,
			Issuer:    ISSUER,
			NotBefore: time.Now().Unix(), // 生效时间,
		},
	}
	return auth
}

/********生成token**********/
func (a *AuthToken) EncodeToken(userInter interface{}) (token string, err error) {
	a.UserInter = userInter
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, a)
	if token, err = jwtToken.SignedString([]byte(SIGNED_KEY)); err == nil {
		return token, nil
	}
	return "", errors.New("Encode Token failed")
}

/********解码token**********/
func (a *AuthToken) DecodeToken(token string) (*AuthToken, error) {

	jwtToken, err := jwt.ParseWithClaims(token, &AuthToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGNED_KEY), nil
	})

	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors&jwt.ValidationErrorMalformed != 0 { // Token is malformed
				return nil, errors.New("验证错误格式不正确")
			} else if e.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("验证错误过期")
			} else if e.Errors&jwt.ValidationErrorIssuer != 0 {
				return nil, errors.New("验证错误发行人")
			} else if e.Errors&jwt.ValidationErrorIssuedAt != 0 {
				return nil, errors.New("签发时间错误")
			} else if e.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token 未生效")
			} else {
				return nil, err
			}
		}
	}

	if claims, ok := jwtToken.Claims.(*AuthToken); ok && jwtToken.Valid {
		return claims, nil
	}

	return nil, errors.New("token 无法解析")
}
