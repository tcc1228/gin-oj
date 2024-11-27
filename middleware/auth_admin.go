package middleware

import (
	"gin-gorm-oj/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthAdminCheck is a middleware that checks if the user is authenticated with admin role
func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Check if user is admin
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Authorization",
			})
			c.Abort()
			return
		}
		if userClaim == nil || userClaim.IsAdmin != 1 {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Admin",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
