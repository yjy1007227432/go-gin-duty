package jwt

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/util"
	"net/http"
)

func Identification() func(c *gin.Context) {

	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		token := c.Query("token")
		username, err := util.DecrpytToken(token)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_DECRYPT_TOKEN_FAIL, nil)
			c.Abort()
			return
		}
		name, group, err := (&auth_service.Auth{
			Username: username,
		}).GetNameByUsername()

		if err != nil || name == "" || group == "" {
			appG.Response(http.StatusInternalServerError, e.ERROR_GET_NAME_FAIL, nil)
			c.Abort()
			return
		}

		c.Set("name", name)
		c.Set("group", group)

		c.Next()
	}
}
