package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/duty_overtime"
	"go-gin-duty-master/service/duty_vacation"
	"net/http"
)

func Test(c *gin.Context) {
	appG := app.Gin{C: c}

	overtimes, _ := (&duty_overtime.DutyOverTime{}).GetAllNowDay()

	for _, overtime := range overtimes {
		_ = (&duty_vacation.Vacation{
			Name: overtime.Proposer,
		}).Add(overtime.Quantity)
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rotas": "成功",
	})
}
