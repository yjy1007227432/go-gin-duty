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

// @Summary //获取所有调休信息(获得审批同意的调休信息)
// @Produce  json
// @Param token query string true "token"
// @Param month query string true "month"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/getAllowedByMonth   [post]
func GetAllowedRests(c *gin.Context) {

	var (
		rests []models.DutyRest
		err   error
	)
	month := c.Query("month")
	appG := app.Gin{C: c}

	restService := rest_service.Rest{
		Datetime: month,
	}

	rests, err = restService.GetAllowedRestByMonth()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rests": rests,
	})
}

// @Summary 获取所有调休申请表信息
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/getAll   [post]
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

// @Summary 查看需要本人审核的未审核调休申请表信息
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/getNeedExamine   [post]
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

// @Summary 审批调休申请表
// @Produce  json
// @Param token query string true "token"
// @Param id query int true "Id"
// @Param response query string true "回复，状态 0为默认、1为拒绝、2为同意"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/examineRest   [post]
func ExamineRest(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.Query("id")

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

	nowDay := time.Now().Format("2006-01-02")

	if rest.Datetime < nowDay {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}

	//if rest.Checker != name {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_EXAMINE_RESTS_FAIL, nil)
	//	return
	//}
	if rest.Response != 0 {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXAMINA_RESTS_FAIL, nil)
		return
	}

	err = (&rest_service.Rest{
		Id:         idInt,
		Response:   response,
		Checker:    name,
		ResponseOn: time.Now(),
	}).Edit()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPDATE_RESTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 新增本人调休申请表信息
// @Produce  json
// @Param token query string true "token"
// @Param request_time query string true "申请调休时间"
// @Param type query int true "申请调休类型，0：上午，1：下午，2：全天"
// @Param vacation_type query int true "申请调休类型，0：调休，1：年休"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/addMyRest [post]
func AddRest(c *gin.Context) {
	appG := app.Gin{C: c}

	datetime := c.Query("request_time")
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

// @Summary 删除本人未审批调休申请表信息
// @Produce  json
// @Param token query string true "token"
// @Param  id query int true "Id"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/deleteMyRest [post]
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

	nowDay := time.Now().Format("2006-01-02")
	if rest1.Datetime < nowDay {
		appG.Response(http.StatusInternalServerError, e.ERROR_TIME_EARLY_FAIL, nil)
		return
	}

	//if rest1.Response != 0 {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_NOT_CHANGE_RESTS_FAIL, nil)
	//	return
	//}
	err = rest.DeleteById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_RESTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 获取本人调休申请表信息(未审批/已审批)
// @Produce  json
// @Param token query string true "token"
// @Param state query int false "0 未审批、1 拒绝、 2 同意 "
// @Success 200 {string} string	 "{"code":200,"data":{rest},"msg":"ok"}"
// @Router /api/rests/getMe [post]
func GetMyRest(c *gin.Context) {
	appG := app.Gin{C: c}

	state := c.Query("state")

	stateInt, _ := strconv.Atoi(state)

	name := (&util.GetName{C: *c}).GetName()

	var rests []models.DutyRest
	var err error

	if state == "" {
		rests, err = (&rest_service.Rest{
			Proposer: name,
		}).GetRestsByName()
	} else {
		rests, err = (&rest_service.Rest{
			Proposer: name,
			State:    stateInt,
		}).GetRestsByNameState()
	}

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rests": rests,
	})
}

// @Summary 删除某个员工某天的调休申请表信息
// @Produce  json
// @Param token query string true "token"
// @Param proposer query string true "proposer"
// @Param datetime query string true "proposer"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/deleteAll   [post]
func DeleteRestByProposer(c *gin.Context) {
	appG := app.Gin{C: c}

	restService := rest_service.Rest{}

	err := restService.DeleteAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_RESTS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 删除所有调休申请表信息
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/rests/deleteAll   [post]
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
