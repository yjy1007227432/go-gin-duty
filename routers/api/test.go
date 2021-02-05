package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/timely_task"
	"net/http"
)

func Test(c *gin.Context) {
	appG := app.Gin{C: c}

	timely_task.ComputeVacation()

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rotas": "成功",
	})
}
