package middlewares

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateEncryptPasswd(Password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	encodePWD := string(hash)
	fmt.Println("oricode:" + Password)
	fmt.Println("encode:" + encodePWD)
	return encodePWD, nil
}

type MyClaims struct {
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

func GetSecret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil // 这是我的secret
	}
}

func GenerateJWT(phone string) (tokenString string, err error) {
	mySigningKey := []byte("AllYourBase")
	claim := MyClaims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour * time.Duration(1))), // 过期时间1小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err = token.SignedString(mySigningKey)
	fmt.Println(tokenString)
	return tokenString, err
}

func ParseToken(tokenss string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &MyClaims{}, GetSecret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginParams = LoginParams{}

		if err := c.BindJSON(&loginParams); err != nil {
			fmt.Println(err)
		}

		fmt.Println(loginParams.Username)
		fmt.Println(loginParams.Password)
		enPasswd, err := GenerateEncryptPasswd("Mac8.678")
		if err != nil {
			fmt.Println("GenerateEncryptPasswd error")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(enPasswd), []byte(loginParams.Password)); err != nil {
			fmt.Println("pwd wrong")
		} else {
			fmt.Println("pwd ok")
			GenerateJWT("123456789")
		}

	}

}
