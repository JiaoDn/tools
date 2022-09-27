package sessions

import (
	"github.com/gin-gonic/gin"
)

func SetAuthCookie(c *gin.Context, sessionId string) gin.HandlerFunc {
	c.SetCookie("user_cookie", sessionId, 1000, "/", "127.0.0.1", false, true)
	return nil
}
