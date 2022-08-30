package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
		}

	}

}
