package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	jwt.StandardClaims
}

func GetJwtSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(global.JWTSetting.Expire).Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJwtSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	//parse define
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})

	if tokenClaims != nil {
		//check expire
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
