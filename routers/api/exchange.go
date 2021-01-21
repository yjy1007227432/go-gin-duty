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

// @Summary 新增本人换班申请表信息
// @Produce  json
// @Param token query string true "token"
// @Param respondent query string true "申请对象"
// @Param request_time query string true "申请时间"
// @Param requested_time query string true "被申请时间"
// @Param exchange_type query string true "换班类型，1，晚班，2,周末白班，3，crm工作日特殊班，4，周末全天班"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/exchange/addMyExchange [post]
func AddMyExchange(c *gin.Context) {
	appG := app.Gin{C: c}
	respondent := c.Query("respondent")

	name := (&util.GetName{C: *c}).GetName()

	nowDay := time.Now().Format("2006-01-02")

	//如果过了八点半，则视为下一天
	if time.Now().Format("15:04") > "08:30" {
		nowDay = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	}
	//只能审批之前的换班申请
	if c.Query("request_time") <= nowDay || c.Query("requested_time") <= nowDay {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}

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
	if IsExistRequest == false || IsExistRequested == false {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_ROTA, nil)
		return
	}
	//换班的日期得的确存在申请人值班与被申请人值班
	IsExist, err := rota_service.CheckTwoExist(name, respondent, exchange.RequestTime, exchange.RequestedTime, nameGroup, exchange.ExchangeType)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_ROTAS_EXIST_FAIL, nil)
		return
	}
	if IsExist != true {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_ROTA, nil)
		return
	}

	err = exchange.AddExchange()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_EXCHANGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 查看所有的换班请求表
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/exchange/getAll   [post]
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

// @Summary 清空所有的换班请求表
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/exchange/deleteAll [post]
func DeleteAllExchange(c *gin.Context) {

	appG := app.Gin{C: c}
	err := (&exchange_service.Exchange{}).DeleteAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_EXCHANGE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 查看本人的换班请求表(未审批/已审批)
// @Produce  json
// @Param token query string true "token"
// @Param  state query int true "状态 0：未审批 1：已审批"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/exchange/myExchange [post]
func GetMyExchange(c *gin.Context) {
	var (
		exchanges []models.DutyExchange
		err       error
	)
	appG := app.Gin{C: c}

	state := c.Query("state")

	response, _ := strconv.Atoi(state)

	name := (&util.GetName{C: *c}).GetName()

	exchanges, err = (&exchange_service.Exchange{
		Proposer: name,
		Response: response,
	}).GetMyExchange()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"exchanges": exchanges,
	})
}

// @Summary 查看本人回复的换班申请表信息(未审批/已审批)
// @Produce  json
// @Param token query string true "token"
// @Param state query int true "状态 0：未审批 1：已审批"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/exchange/getMyExamine [post]
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

// @Summary 删除本人的未审批换班请求表
// @Produce  json
// @Param token query string true "token"
// @Param id query int true "Id"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/exchange/deleteMyExchange{id} [post]
func DeleteExchange(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.Query("id")

	if id == "" {
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
		return
	}

	idInt, _ := strconv.Atoi(id)

	name := (&util.GetName{C: *c}).GetName()

	var exchange = &exchange_service.Exchange{
		Id: idInt,
	}

	exchange1, err := exchange.GetExchangeById()

	if name != exchange1.Proposer {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_NOT_ME_EXCHANGE_FAIL, nil)
		return
	}

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

// @Summary 回复换班申请表
// @Produce  json
// @Param token query string true "token"
// @Param  id query int true "Id"
// @Param  response query  int true "回复，状态 0为默认、1为拒绝、2为同意"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/exchange/examineExchange [post]
func ExamineExchange(c *gin.Context) {
	appG := app.Gin{C: c}

	ID := c.Query("id")
	Response := c.Query("response")

	if ID == "" || Response == "" {
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
		return
	}

	id, _ := strconv.Atoi(ID)
	response, _ := strconv.Atoi(Response)

	name := (&util.GetName{C: *c}).GetName()
	group := (&util.GetName{C: *c}).GetGroup()

	exchange, err := (&exchange_service.Exchange{
		Id: id,
	}).GetExchangeById()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_EXCHANGE_FAIL, nil)
		return
	}

	if exchange.Response != 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXAMINA_EXCHANGE_FAIL, nil)
		return
	}

	nowDay := time.Now().Format("2006-01-02")
	//如果过了八点半，则视为下一天
	if time.Now().Format("15:04") > "08:30" {
		nowDay = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	}
	//只能审批之前的换班申请
	if exchange.RequestTime < nowDay || exchange.RequestedTime < nowDay {
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
