package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secretKey    []byte = []byte(beego.AppConfig.String("secretkey"))
	expTime, err        = beego.AppConfig.Int64("exptime")
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func ObtainJWT(username string) string {
	claims := Claims{
		username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + expTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return ss
}
func RefreshJWT(strToken string) string {
	token, err := jwt.ParseWithClaims(strToken, &Claims{}, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		println(err)
		return ""
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		println("trst")
		return ""
	}
	if err := token.Claims.Valid(); err != nil {
		print("test")
		return ""
	}
	//user, err := models.GetUser(claims.Username)
	//if err != nil {
	//	println(err)
	//	return "user not exist"
	//}
	newToken := ObtainJWT(claims.Username)
	return newToken
}
func VerifyJWT(strToken string) string {
	token, err := jwt.ParseWithClaims(strToken, &Claims{}, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		logs.Error(err)
		return ""
	}
	if err := token.Claims.Valid(); err != nil {
		print("test")
		return ""
	}
	return "ok"

}
