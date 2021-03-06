package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-gin-duty-master/cors"

	_ "go-gin-duty-master/docs"

	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-duty-master/middleware/jwt"
	"time"

	"go-gin-duty-master/routers/api"
)

func InitRouter() *gin.Engine {

	// 运行 nsq
	//	servers.NsqRun()
	r := gin.Default()
	r.Use(cors.CorsHandler())
	r.Use(cors.Cors())
	r.POST("/auth", api.GetAuth) //
	r.POST("/register", api.Register)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app := r.Group("/api").Use(jwt.TimeoutMiddleware(time.Second * 100))

	app.Use(jwt.JWT()).Use(jwt.Identification())
	{
		//测试接口
		app.POST("/test", api.Test)
		//获取本人姓名
		app.POST("/auth/getName", api.GetNameByToken)
		//获取所有人信息
		app.POST("/auth/getAll", api.GetAll)
		//获取当月值班表 1
		app.POST("/rotas/getMonth", api.GetRotaByMonth)
		//获取本人调休申请表信息(未审批/已审批) 1
		app.POST("/rests/getMe", api.GetMyRest)
		//新增本人调休申请表信息 //todo
		app.POST("/rests/addMyRest", api.AddRest)
		//删除本人未审批调休申请表信息 1
		app.POST("/rests/deleteMyRest", api.DeleteRest)
		//获取所有调休信息 1
		app.POST("/vacation/getAll", api.GetAllVacation)
		//获取所有调休信息(获得审批同意的调休信息)
		app.POST("/rests/getAllowedByMonth", api.GetAllowedRests)
		//查看本人的换班请求表(未审批/已审批) 1
		app.POST("/exchange/myExchange", api.GetMyExchange)
		//删除本人的未审批换班请求表 1
		app.POST("/exchange/deleteMyExchange", api.DeleteExchange)
		//回复换班申请表 1
		app.POST("/exchange/examineExchange", api.ExamineExchange)
		//查看本人回复的换班申请表信息(未审批/已审批) 1
		app.POST("/exchange/getMyExamine", api.GetNeedExamineExchanges)
		//新增本人换班申请表信息 1
		app.POST("/exchange/addMyExchange", api.AddMyExchange)
		//新增本人加班申请表信息 //todo
		app.POST("/overtime/addMyOvertime", api.AddOvertime)
		//获取本人加班申请表信息
		app.POST("/overtime/getMe", api.GetMyOvertime)
		//删除本人加班调休申请表信息 1
		app.POST("/overtime/deleteMyOvertime", api.DeleteOvertime)
	}

	app.Use(jwt.JWT()).Use(jwt.ADMIN()).Use(jwt.Identification())
	{
		//新增员工信息表
		app.POST("/auth/AddAuth", api.AddAuth)
		//导入值班表  1
		app.POST("/rota/import", api.ImportRota)
		//删除月值班表 1
		app.POST("/rotas/deleteMonth", api.DeleteRotaByMonth)
		//删除日值班表 1
		app.POST("/rotas/deleteDay", api.DeleteRotaByDay)
		//添加日值班表 1
		app.POST("/rotas/addDay", api.AddRotaByDay)
		//删除所有调休申请表信息
		app.POST("/rests/deleteAll", api.DeleteRests)
		//获取所有调休申请表信息/
		app.POST("/rests/getAll", api.GetRests)
		//查看需要本人审核的未审核调休申请表信息
		app.POST("/rests/getNeedExamine", api.GetNeedExamineRests)
		//审批调休申请表
		app.POST("/rests/examineRest", api.ExamineRest)
		//清空所有调休信息 1
		app.POST("/vacation/deleteAll", api.DeleteAllVacation)
		//删除某人调休信息 1
		app.POST("/vacation/deleteByName", api.DeleteVacationByName)
		//修改某人调休信息 1
		app.POST("/vacation/editByName", api.EditVacationByName)
		//查看所有的换班请求表 1
		app.POST("/exchange/getAll", api.GetAllExchange)
		//清空所有的换班请求表 1
		app.POST("/exchange/deleteAll", api.DeleteAllExchange)
		//查看需要审核的未审核加班申请表信息
		app.POST("/overtime/getNeedExamine", api.GetNeedExamineOvertime)
		//审批加班申请表信息 //todo
		app.POST("/overtime/examineOvertime", api.ExamineOverTime)
		//查看需要审核的未审核加班申请表信息
		app.POST("/overtime/getAll", api.GetALLOvertime)

	}

	return r
}
