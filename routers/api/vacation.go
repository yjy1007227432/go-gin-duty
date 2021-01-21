package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/duty_vacation"
	"net/http"
	"time"
)

// @Summary 获取所有员工的调休信息
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/vacation/getAll [post]
func GetAllVacation(c *gin.Context) {
	var (
		vacations []models.DutyVacation
		err       error
	)
	appG := app.Gin{C: c}
	vacationService := &duty_vacation.Vacation{}

	vacations, err = vacationService.GetAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_VACATION_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"vacations": vacations,
	})
}

// @Summary 清空所有调休信息
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/vacation/deleteAll   [post]
func DeleteAllVacation(c *gin.Context) {

	appG := app.Gin{C: c}

	vacationService := &duty_vacation.Vacation{}

	err := vacationService.DeleteAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_VACATION_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 删除某人调休信息
// @Produce  json
// @Param token query string true "token"
// @Param  name query string true "姓名"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/vacation/deleteByName   [post]
func DeleteVacationByName(c *gin.Context) {

	appG := app.Gin{C: c}

	name := c.Query("name")

	if name == "" {
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
		return
	}

	vacationService := &duty_vacation.Vacation{
		Name: name,
	}

	err := vacationService.DeleteByName()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_VACATION_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 修改某人调休信息
// @Produce  json
// @Param token query string true "token"
// @Param  name query int true "姓名"
// @Param  remain_vacation  query float64 true "剩余调休"
// @Param  remain_annual_vacation query float64 true "剩余年休"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/vacation/editByName   [post]
func EditVacationByName(c *gin.Context) {
	appG := app.Gin{C: c}

	vacation := duty_vacation.Vacation{}

	err := c.Bind(&vacation)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
		return
	}

	vaca, err := (&duty_vacation.Vacation{
		Name:       vacation.Name,
		UpdateTime: time.Now(),
	}).GetByName()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_VACATION_FAIL, nil)
		return
	}

	vacation.Id = vaca.Id

	err = vacation.Edit()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_VACATION_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
