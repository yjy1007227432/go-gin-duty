package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/util"
	"net/http"
)

// @Summary 新增本人加班申请表信息
// @Produce  json
// @Param token query string true "token"
// @Param quantity query float64 true "加班时间"
// @Param reason query string true "加班理由"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/addMyRest [post]
func AddOvertime(c *gin.Context) {
	appG := app.Gin{C: c}
	name := (&util.GetName{C: *c}).GetName()

	var dutyOvertime = &models.DutyOverTime{
		Datetime: datetime,
		Proposer: name,
		Response: 0,
	}

	err = c.Bind(rest)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
		return
	}

	IsExist, err := rest.CheckIsExist()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}
	if IsExist == true {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_RESTS_FAIL, nil)
		return
	}

	err = rest.Add()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RESTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
