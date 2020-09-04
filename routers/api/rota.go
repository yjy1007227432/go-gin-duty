package api

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/rota_service"
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

func AddRota(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddRotaForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	rotaService := rota_service.Rota{
		Datetime:           form.Datetime,
		Week:               form.Week,
		BillingLate:        form.BillingLate,
		BillingWeekendLate: form.BillingWeekendLate,
		CrmLate:            form.CrmLate,
		CrmWeekendLate:     form.CrmWeekendLate,
		CrmDuty:            form.CrmDuty,
		CreatedBy:          form.CreatedBy,
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
	Datetime           string `form:"datetime '' comment('日期') VARCHAR(50)"`
	Week               string `form:"week '' comment('日期') VARCHAR(50)"`
	BillingLate        string `form:"billing_late '' comment('计费晚班人员') VARCHAR(50)"`
	BillingWeekendLate string `form:"billing_weekend_late '' comment('计费周末晚班人员') VARCHAR(50)"`
	CrmLate            string `form:"crm_late '' comment('crm晚班人员') VARCHAR(50)"`
	CrmWeekendLate     string `form:"crm_weekend_late '' comment('crm周末晚班人员') VARCHAR(50)"`
	CrmDuty            string `form:"crm_duty '' comment('crm值班人员') VARCHAR(50)"`
	CreatedBy          string `form:"created_by '' comment('创建人') VARCHAR(100)"`
}
