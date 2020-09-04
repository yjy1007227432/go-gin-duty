package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/middleware/jwt"

	"go-gin-duty-master/routers/api"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/auth", api.GetAuth) //

	app := r.Group("/api")
	{
		//导入标签
		r.POST("/rota/import", api.ImportRota)
	}

	app.Use(jwt.JWT())

	return r
}
