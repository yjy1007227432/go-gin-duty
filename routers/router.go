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
		//获取当月值班表
		r.POST("/rotas/getMonth", api.GetRotaByMonth)
		//获取本人调休申请表信息
		r.POST("/rests/getMe", api.GetMyRest)
		//新增本人调休申请表信息
		r.POST("/rests/addMyRest", api.AddRest)
		//删除本人未审批调休申请表信息
		r.POST("/rests/deleteMyRest", api.DeleteRest)
		//获取所有调休信息
		r.POST("/vacation/getAll", api.GetAllVacation)

	}

	app.Use(jwt.JWT()).Use(jwt.ADMIN())
	{
		//导入值班表
		r.POST("/rota/import", api.ImportRota)
		//删除月值班表
		r.POST("/rotas/deleteMonth", api.DeleteRotaByMonth)
		//删除日值班表
		r.POST("/rotas/deleteDay", api.DeleteRotaByDay)
		//添加日值班表
		r.POST("/rotas/addDay", api.AddRotaByDay)

		//删除所有调休申请表信息
		r.POST("/rests/deleteAll", api.DeleteRests)
		//获取所有调休申请表信息/
		r.POST("/rests/getAll", api.GetRests)
		//查看需要本人审核的未审核调休申请表信息
		r.POST("/rests/getAll", api.GetNeedExamineRests)
		//审批调休申请表
		r.POST("/rests/examineRest", api.ExamineRest)

		//清空所有调休信息
		r.POST("/vacation/deleteAll", api.DeleteAllVacation)
		//删除某人调休信息
		r.POST("/vacation/deleteByName", api.DeleteVacationByName)
		//修改某人调休信息
		r.POST("/vacation/editByName", api.EditVacationByName)

		//查看所有的换班请求表
		r.POST("/exchange/getAll", api.GetAllExchange)

	}

	return r
}
