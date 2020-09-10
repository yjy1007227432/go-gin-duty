package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/rest_service"
	"go-gin-duty-master/service/rota_service"
	"go-gin-duty-master/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetRests(c *gin.Context) {

	var (
		rests []models.DutyRest
		err   error
	)

	appG := app.Gin{C: c}

	restService := rest_service.Rest{}

	rests, err = restService.GetAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rests": rests,
	})
}

func GetNeedExamineRests(c *gin.Context) {
	var (
		rests []models.DutyRest
		err   error
	)
	appG := app.Gin{C: c}

	name := (&util.GetName{C: *c}).GetName()

	restService := rest_service.Rest{
		Checker: name,
	}
	rests, err = restService.GetByChecker()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rests": rests,
	})
}

func ExamineRest(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.Query("Id")

	datetime := c.Query("datetime")

	nowDay := time.Now().Format("2006-01-02")
	if datetime < nowDay {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}

	response, _ := strconv.Atoi(c.Query("response"))

	idInt, _ := strconv.Atoi(id)

	name := (&util.GetName{C: *c}).GetName()

	rest, err := (&rest_service.Rest{
		Id: idInt,
	}).GetRestById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}

	if rest.Checker != name {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXAMINE_RESTS_FAIL, nil)
		return
	}
	if rest.Response != 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXAMINA_RESTS_FAIL, nil)
		return
	}

	err = (&rest_service.Rest{
		Id:       idInt,
		Response: response,
	}).Edit()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPDATE_RESTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddRest(c *gin.Context) {
	appG := app.Gin{C: c}

	datetime := c.Query("datetime")
	nowDay := time.Now().Format("2006-01-02")
	if datetime < nowDay {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}

	rota, err := (&rota_service.Rota{
		Datetime: datetime,
	}).GetRotaByDay()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROTAS_FAIL, nil)
		return
	}

	if rota.Week == "星期六" || rota.Week == "星期日" {
		appG.Response(http.StatusInternalServerError, e.ERROR_REST_WEEKEND_FAIL, nil)
		return
	}

	name := (&util.GetName{C: *c}).GetName()

	if strings.Contains(rota.BillingLate, name) || strings.Contains(rota.CrmDutySpecial, name) || strings.Contains(rota.CrmLate, name) {
		appG.Response(http.StatusInternalServerError, e.ERROR_ROTA_REST_FAIL, nil)
		return
	}

	var rest = &rest_service.Rest{
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
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteRest(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.Query("id")

	idInt, _ := strconv.Atoi(id)

	var rest = &rest_service.Rest{
		Id: idInt,
	}

	rest1, err := rest.GetRestById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}
	if rest1.Response != 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_CHANGE_RESTS_FAIL, nil)
		return
	}
	err = rest.DeleteById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_RESTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetMyRest(c *gin.Context) {
	appG := app.Gin{C: c}

	state := c.Query("state")
	stateInt, _ := strconv.Atoi(state)

	name := (&util.GetName{C: *c}).GetName()

	rests, err := (&rest_service.Rest{
		Proposer: name,
		Response: stateInt,
	}).GetRestsByName()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rests": rests,
	})
}

func DeleteRests(c *gin.Context) {
	appG := app.Gin{C: c}

	restService := rest_service.Rest{}

	err := restService.DeleteAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_RESTS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
