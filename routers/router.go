package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/middleware/jwt"
	"go-gin-duty-master/servers"
	"time"

	"go-gin-duty-master/routers/api"
)

func InitRouter() *gin.Engine {

	// 运行 nsq
	servers.NsqRun()

	r := gin.Default()
	r.POST("/auth", api.GetAuth) //
	r.POST("/register", api.Register)

	app := r.Group("/api").Use(jwt.TimeoutMiddleware(time.Second * 2))
	app.Use(jwt.JWT())
	{
		//获取当月值班表
		r.POST("/rotas/getMonth", api.GetRotaByMonth)

		//获取本人调休申请表信息(未审批/已审批)
		r.POST("/rests/getMe", api.GetMyRest)

		//新增本人调休申请表信息
		r.POST("/rests/addMyRest", api.AddRest)
		//todo  周末和法定节假日

		//删除本人未审批调休申请表信息
		r.POST("/rests/deleteMyRest", api.DeleteRest)

		//获取所有调休信息
		r.POST("/vacation/getAll", api.GetAllVacation)

		//查看本人的换班请求表(未审批/已审批)
		r.POST("/exchange/myExchange", api.GetMyExchange)
		//删除本人的未审批换班请求表
		r.POST("/exchange/deleteMyExchange", api.DeleteExchange)
		//回复换班申请表
		r.POST("/exchange/examineExchange", api.ExamineExchange)
		//查看本人回复的换班申请表信息(未审批/已审批)
		r.POST("/exchange/getMyExamine", api.GetNeedExamineExchanges)
		//新增本人换班申请表信息
		r.POST("/exchange/addMyExchange", api.AddMyExchange)
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
		//清空所有的换班请求表
		r.POST("/exchange/deleteAll", api.DeleteAllExchange)

	}

	return r
}
