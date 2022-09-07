package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeConsuming() gin.HandlerFunc {
	return func(c *gin.Context) {
		start_time := time.Now().UnixNano()
		c.Next()
		end_time := time.Now().UnixNano()
		fmt.Println(end_time - start_time)
		c.Set("time", end_time-start_time)

	}

}
