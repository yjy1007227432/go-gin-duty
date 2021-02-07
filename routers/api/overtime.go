package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/duty_overtime"
	"go-gin-duty-master/util"
	"net/http"
	"strconv"
	"time"
)

// @Summary 新增本人加班申请表信息
// @Produce  json
// @Param token query string true "token"
// @Param quantity query float64 true "加班时间"
// @Param reason query string true "加班理由"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/overtime/addMyOvertime [post]
func AddOvertime(c *gin.Context) {
	appG := app.Gin{C: c}
	name := (&util.GetName{C: *c}).GetName()

	quantity, _ := strconv.ParseFloat(c.Query("quantity"), 64)
	//reason := c.Query("reason")

	var dutyOvertime = &duty_overtime.DutyOverTime{
		Proposer: name,
		Quantity: quantity,
	}

	err := c.Bind(dutyOvertime)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
		return
	}
	err = dutyOvertime.AddOverTime()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_OVERTIME_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 审批加班申请表
// @Produce  json
// @Param token query string true "token"
// @Param id query int true "Id"
// @Param response query string true "回复，状态 0为默认、1为拒绝、2为同意"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/overtime/examineOvertime   [post]
func ExamineOverTime(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.Query("id")
	response, _ := strconv.Atoi(c.Query("response"))

	idInt, _ := strconv.Atoi(id)

	dutyOverTime, err := (&duty_overtime.DutyOverTime{
		Id: idInt,
	}).GetOverTimeById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_OVERTIME_FAIL, nil)
		return
	}

	if dutyOverTime.Response != 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXAMINA_OVERTIME_FAIL, nil)
		return
	}

	err = (&duty_overtime.DutyOverTime{
		Id:         idInt,
		Response:   response,
		ResponseOn: time.Now(),
	}).Edit()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPDATE_DUTYOVERTIME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 查看需要审核的未审核加班申请表信息
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/overtime/GetNeedExamineOvertime   [post]
func GetNeedExamineOvertime(c *gin.Context) {
	var (
		overtimes []models.DutyOvertime
		err       error
	)
	appG := app.Gin{C: c}

	overtimes, err = (&duty_overtime.DutyOverTime{}).GetAllNeedExamine()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_OVERTIME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"overtimes": overtimes,
	})
}

// @Summary 获取本人加班申请表信息(未审批/已审批)
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{rest},"msg":"ok"}"
// @Router /api/overtime/getMe [post]
func GetMyOvertime(c *gin.Context) {
	appG := app.Gin{C: c}

	name := (&util.GetName{C: *c}).GetName()

	var (
		overtimes []models.DutyOvertime
		err       error
	)

	overtimes, err = (&duty_overtime.DutyOverTime{
		Proposer: name,
	}).GetOvertimesByName()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_OVERTIME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"overtimes": overtimes,
	})
}

// @Summary 获取所有加班申请表信息
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/overtime/getAll   [post]
func GetALLOvertime(c *gin.Context) {
	var (
		overtimes []models.DutyOvertime
		err       error
	)

	appG := app.Gin{C: c}

	overtimes, err = (&duty_overtime.DutyOverTime{}).GetAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_OVERTIME_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"overtimes": overtimes,
	})
}

// @Summary 删除本人加班申请表信息
// @Produce  json
// @Param token query string true "token"
// @Param  id query int true "Id"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/overtime/deleteMyOvertime [post]
func DeleteOvertime(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.Query("id")

	idInt, _ := strconv.Atoi(id)

	var dutyOvertime = &duty_overtime.DutyOverTime{
		Id: idInt,
	}

	overtime, err := dutyOvertime.GetOverTimeById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_OVERTIME_FAIL, nil)
		return
	}

	nowDay := time.Now().Format("2006-01-02")
	overtimeDay := overtime.CreatedOn.Format("2006-01-02")
	if overtimeDay < nowDay {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}

	if time.Now().Hour() > 23 {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}

	//if rest1.Response != 0 {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_NOT_CHANGE_RESTS_FAIL, nil)
	//	return
	//}

	err = dutyOvertime.DeleteOvertimeById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_OVERTIME_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
