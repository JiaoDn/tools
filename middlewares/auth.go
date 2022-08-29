package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginParams = LoginParams{}

		if err := c.BindJSON(&loginParams); err != nil {
			fmt.Println(err)
		}

		fmt.Println(loginParams.Username)
		fmt.Println(loginParams.Password)

	}

}
