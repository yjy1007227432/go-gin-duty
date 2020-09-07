package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/service/exchange_service"
	"go-gin-duty-master/service/rota_service"
	"strings"
	"time"

	"go-gin-duty-master/util"
	"net/http"
	"strconv"
)

func AddMyExchange(c *gin.Context) {
	appG := app.Gin{C: c}

	token := c.Query("token")
	respondent := c.Query("respondent")
	username, err := util.DecrpytToken(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DECRYPT_TOKEN_FAIL, nil)
		return
	}

	name, err := auth_service.Auth.GetNameByUsername(username)

	if name == respondent {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXCHANGE_SAME_FAIL, nil)
		return
	}

	nameGroup, err := auth_service.Auth.GetGroupByName(name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AUTH_FAIL, nil)
		return
	}
	respondentGroup, err := auth_service.Auth.GetGroupByName(respondent)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AUTH_FAIL, nil)
		return
	}
	if nameGroup != respondentGroup {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXCHANGE_SAME_FAIL, nil)
		return
	}

	var exchange = &exchange_service.Exchange{}
	err = c.Bind(exchange)
	exchange.Proposer = name

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
		return
	}
	//todo
	//换班日期得存在在调休表中
	//两个换班日期的值班人员得是申请人与被申请人
	IsExistRequest, err := rota_service.Rota{
		Datetime: exchange.RequestTime,
	}.ExistByDatetime()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROTAS_FAIL, nil)
		return
	}
	IsExistRequested, err := rota_service.Rota{
		Datetime: exchange.RequestedTime,
	}.ExistByDatetime()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROTAS_FAIL, nil)
		return
	}
	if IsExistRequest && IsExistRequested == false {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_ROTA, nil)
		return
	}
	rotaRequest, err := rota_service.Rota{
		Datetime: exchange.RequestTime,
	}.GetRotaByDay()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROTAS_FAIL, nil)
		return
	}
	rotaRequested, err := rota_service.Rota{
		Datetime: exchange.RequestedTime,
	}.GetRotaByDay()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROTAS_FAIL, nil)
		return
	}
	if !(strings.Contains(rotaRequest.CrmLate, name) || strings.Contains(rotaRequest.BillingLate, name)) {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_ROTAS_FAIL, nil)
		return
	}
	if !(strings.Contains(rotaRequested.CrmLate, respondent) || strings.Contains(rotaRequested.BillingLate, respondent)) {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_ROTAS_FAIL, nil)
		return
	}
	//不能有涉及到这两天的未处理换班请求
	isExist, err := exchange.IsExistDay()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}
	if isExist == true {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_EXCHANGE_FAIL, nil)
		return
	}
	err = exchange.AddExchange()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_EXCHANGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetAllExchange(c *gin.Context) {

	var (
		exchanges []models.DutyExchange
		err       error
	)

	appG := app.Gin{C: c}

	exchanges, err = exchange_service.Exchange{}.GetAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"exchanges": exchanges,
	})
}

func DeleteAllExchange(c *gin.Context) {

	appG := app.Gin{C: c}
	err := exchange_service.Exchange{}.DeleteAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_EXCHANGE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetMyExchange(c *gin.Context) {
	var (
		exchanges []models.DutyExchange
		err       error
	)
	appG := app.Gin{C: c}

	state := c.Query("state")
	stateInt, _ := strconv.Atoi(state)
	token := c.Query("token")
	username, err := util.DecrpytToken(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DECRYPT_TOKEN_FAIL, nil)
		return
	}
	name, err := auth_service.Auth.GetNameByUsername(username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AUTH_FAIL, nil)
		return
	}

	exchanges, err = exchange_service.Exchange{
		Proposer: name,
		Response: stateInt,
	}.GetMyExchange()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"exchanges": exchanges,
	})
}

func GetNeedExamineExchanges(c *gin.Context) {
	var (
		exchanges []models.DutyExchange
		err       error
	)
	appG := app.Gin{C: c}

	state := c.Query("state")
	stateInt, _ := strconv.Atoi(state)
	token := c.Query("token")
	username, err := util.DecrpytToken(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DECRYPT_TOKEN_FAIL, nil)
		return
	}
	name, err := auth_service.Auth.GetNameByUsername(username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AUTH_FAIL, nil)
		return
	}

	exchanges, err = exchange_service.Exchange{
		Respondent: name,
		Response:   stateInt,
	}.GetMyExamineExchange()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"exchanges": exchanges,
	})
}

func DeleteExchange(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.Query("Id")

	idInt, _ := strconv.Atoi(id)

	var exchange = &exchange_service.Exchange{
		Id: idInt,
	}

	exchange1, err := exchange.GetExchangeById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}
	if exchange1.Response != 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_CHANGE_EXCHANGE_FAIL, nil)
		return
	}
	err = exchange.DeleteById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_EXCHANGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func ExamineExchange(c *gin.Context) {
	appG := app.Gin{C: c}

	requestTime := c.Query("request_time")
	requestedTime := c.Query("requested_time")
	nowDay := time.Now().Format("2006-01-02")
	if requestTime < nowDay || requestedTime < nowDay {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}

	id := c.Query("Id")
	response, _ := strconv.Atoi(c.Query("response"))
	idInt, _ := strconv.Atoi(id)

	token := c.Query("token")
	username, err := util.DecrpytToken(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DECRYPT_TOKEN_FAIL, nil)
		return
	}
	name, err := auth_service.Auth.GetNameByUsername(username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AUTH_FAIL, nil)
		return
	}

	exchange, err := exchange_service.Exchange{
		Id: idInt,
	}.GetExchangeById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}
	if exchange.Respondent != name {
		appG.Response(http.StatusInternalServerError, e.ERROR_RESPONCE_EXCHANGE_FAIL, nil)
		return
	}
	if exchange.Response != 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXAMINA_EXCHANGE_FAIL, nil)
		return
	}

	err = exchange_service.Exchange{
		Id:       idInt,
		Response: response,
	}.Edit()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPDATE_EXCHANGE_FAIL, nil)
		return
	}
	//todo
	//同意换班之后更新值班表

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
