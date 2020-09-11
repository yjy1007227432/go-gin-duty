package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/service/exchange_service"
	"go-gin-duty-master/service/rota_service"
	"time"

	"go-gin-duty-master/util"
	"net/http"
	"strconv"
)

func AddMyExchange(c *gin.Context) {
	appG := app.Gin{C: c}
	respondent := c.Query("respondent")

	name := (&util.GetName{C: *c}).GetName()

	//自身不能换班
	if name == respondent {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXCHANGE_SAME_FAIL, nil)
		return
	}

	//必须是同组人员才可以换班
	nameGroup := (&util.GetName{C: *c}).GetGroup()

	respondentGroup, err := (&auth_service.Auth{
		Name: respondent,
	}).GetGroupByName()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AUTH_FAIL, nil)
		return
	}
	if nameGroup != respondentGroup {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXCHANGE_SAME_FAIL, nil)
		return
	}

	var exchange = exchange_service.Exchange{}
	exchange.Proposer = name
	err = c.Bind(&exchange)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
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

	//换班的日期存在在值班表中
	IsExistRequest, err := (&rota_service.Rota{
		Datetime: exchange.RequestTime,
	}).ExistByDatetime()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROTAS_FAIL, nil)
		return
	}
	IsExistRequested, err := (&rota_service.Rota{
		Datetime: exchange.RequestedTime,
	}).ExistByDatetime()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROTAS_FAIL, nil)
		return
	}
	if IsExistRequest && IsExistRequested == false {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_ROTA, nil)
		return
	}
	//换班的日期得的确存在申请人值班与被申请人值班
	IsExist, err := rota_service.CheckTwoExist(name, respondent, exchange.RequestTime, exchange.RequestedTime, nameGroup, exchange.ExchangeType)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_ROTAS_EXIST_FAIL, nil)
		return
	}
	if IsExist != true {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_ROTA, nil)
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

	appG := app.Gin{C: c}
	var (
		exchanges []models.DutyExchange
		err       error
	)

	exchanges, err = (&exchange_service.Exchange{}).GetAll()

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
	err := (&exchange_service.Exchange{}).DeleteAll()

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

	name := (&util.GetName{C: *c}).GetName()

	exchanges, err = (&exchange_service.Exchange{
		Proposer: name,
		Response: stateInt,
	}).GetMyExchange()

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
	name := (&util.GetName{C: *c}).GetName()

	exchanges, err = (&exchange_service.Exchange{
		Respondent: name,
		Response:   stateInt,
	}).GetMyExamineExchange()

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

	id := c.Query("id")

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

	id, _ := strconv.Atoi(c.Query("Id"))
	response, _ := strconv.Atoi(c.Query("response"))
	name := (&util.GetName{C: *c}).GetName()
	group := (&util.GetName{C: *c}).GetGroup()

	exchange, err := (&exchange_service.Exchange{
		Id: id,
	}).GetExchangeById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}
	nowDay := time.Now().Format("2006-01-02")
	//如果过了八点半，则视为下一天
	if time.Now().Format("15:04") > "08:30" {
		nowDay = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	}
	//只能审批之前的换班申请
	if exchange.RequestTime <= nowDay || exchange.RequestedTime <= nowDay {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}
	//只能审批本人的换班申请
	if exchange.Respondent != name {
		appG.Response(http.StatusInternalServerError, e.ERROR_RESPONCE_EXCHANGE_FAIL, nil)
		return
	}
	if exchange.Response != 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXAMINA_EXCHANGE_FAIL, nil)
		return
	}

	err = (&exchange_service.Exchange{
		Id:       id,
		Response: response,
	}).ExchangeTwo(group)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPDATE_EXCHANGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
