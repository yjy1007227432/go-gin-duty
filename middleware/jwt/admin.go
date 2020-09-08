package jwt

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/util"
	"net/http"
)

func ADMIN() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code    int
			data    interface{}
			isAdmin int
		)

		code = e.SUCCESS
		token := c.Query("token")
		username, err := util.DecrpytToken(token)

		if err != nil {
			code = e.ERROR_DECRYPT_TOKEN_FAIL
		}

		isAdmin, err = (&auth_service.Auth{
			Username: username,
		}).IsAdmin()

		if isAdmin == 0 {
			code = e.ERROR_NOT_ADMIN
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
