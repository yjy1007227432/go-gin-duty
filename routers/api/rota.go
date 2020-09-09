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

	day := c.Query("day")

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

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	name := (&util.GetName{C: *c}).GetName()
	rotaService := rota_service.Rota{
		Datetime:           form.Datetime,
		Week:               form.Week,
		BillingLate:        form.BillingLate,
		BillingWeekendLate: form.BillingWeekendLate,
		CrmLate:            form.CrmLate,
		CrmWeekendLate:     form.CrmWeekendLate,
		CrmDuty:            form.CrmDuty,
		CreatedBy:          name,
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

	if err != nil {
		logs.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	rota := rota_service.Rota{}

	err = rota.Import(file)

	if err != nil {
		logs.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_IMPORT_ROTA_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type AddRotaForm struct {
	Datetime           string `form:"datetime"`
	Week               string `form:"week"`
	BillingLate        string `form:"billing_late"`
	BillingWeekendLate string `form:"billing_weekend_late"`
	CrmLate            string `form:"crm_late"`
	CrmWeekendLate     string `form:"crm_weekend_late"`
	CrmDuty            string `form:"crm_duty"`
	CreatedBy          string `form:"created_by"`
}
