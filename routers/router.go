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

	app.Use(jwt.JWT())
	{
		//导入值班表
		r.POST("/rota/import", api.ImportRota)
		//获取当月值班表
		r.POST("/rotas", api.GetRotaByMonth)

	}

	return r
}
