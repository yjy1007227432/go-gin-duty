package api

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/rota_service"
	"go-gin-duty-master/util"
	"net/http"
)

func GetRotaByMonth(c *gin.Context) {
	appG := app.Gin{C: c}

	month := c.Query("month")

	//参数验证
	if month == "" {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	rotaService := rota_service.Rota{
		Datetime: month,
	}
	rotas, err := rotaService.GetThisMonth()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROTAS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rotas": rotas,
	})
}

func DeleteRotaByMonth(c *gin.Context) {
	appG := app.Gin{C: c}

	month := c.Query("month")

	//参数验证
	if month == "" {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	rotaService := rota_service.Rota{
		Datetime: month,
	}

	err := rotaService.DeleteThisMonth()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ROTAS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteRotaByDay(c *gin.Context) {
	appG := app.Gin{C: c}

	day := c.Query("datetime")

	//参数验证
	if day == "" {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	rotaService := rota_service.Rota{
		Datetime: day,
	}

	err := rotaService.DeleteThisDay()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ROTAS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddRotaByDay(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		form AddRotaForm
	)
	err := c.Bind(&form)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
		return
	}

	name := (&util.GetName{C: *c}).GetName()
	rotaService := rota_service.Rota{
		Datetime:          form.Datetime,
		Week:              form.Week,
		BillingLate:       form.BillingLate,
		BillingWeekendDay: form.BillingWeekendDay,
		CrmLate:           form.CrmLate,
		CrmWeekendDay:     form.CrmWeekendDay,
		CrmDutySpecial:    form.CrmDutySpecial,
		CreatedBy:         name,
	}

	exists, err := rotaService.ExistByDatetime()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_ROTA_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_ROTA, nil)
		return
	}

	err = rotaService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ROTA_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func ImportRota(c *gin.Context) {
	appG := app.Gin{C: c}
	file, _, err := c.Request.FormFile("file")

	//参数验证
	if file == nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	name := (&util.GetName{C: *c}).GetName()

	if err != nil {
		logs.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	rota := rota_service.Rota{}

	err = rota.Import(file, name)

	if err != nil {
		logs.Warn(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Code": "", "Msg": err.Error(), "Data": nil})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type AddRotaForm struct {
	Datetime          string `form:"datetime" binding:"required"`
	Week              string `form:"week" binding:"required"`
	BillingLate       string `form:"billing_late"`
	BillingWeekendDay string `form:"billing_weekend_day"`
	CrmLate           string `form:"crm_late"`
	CrmWeekendDay     string `form:"crm_weekend_day"`
	CrmDutySpecial    string `form:"crm_duty_special"`
	CreatedBy         string `form:"created_by"`
}
