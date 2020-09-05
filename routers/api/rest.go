package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/service/rest_service"
	"go-gin-duty-master/util"
	"net/http"
	"strconv"
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

	rest, err := rest_service.Rest{
		Id: idInt,
	}.GetRestById()

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

	err = rest_service.Rest{
		Id:       idInt,
		Response: response,
	}.Edit()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPDATE_RESTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func AddRest(c *gin.Context) {
	appG := app.Gin{C: c}

	token := c.Query("token")

	username, err := util.DecrpytToken(token)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DECRYPT_TOKEN_FAIL, nil)
		return
	}

	name, err := auth_service.Auth.GetNameByUsername(username)

	var rest = &rest_service.Rest{}
	err = c.Bind(rest)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
		return
	}

	rest.Proposer = name
	err = rest.Add()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESTS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteRest(c *gin.Context) {
	appG := app.Gin{C: c}

	id := c.Query("Id")

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

	rests, err := rest_service.Rest{
		Proposer: name,
		State:    stateInt,
	}.GetRestsByName()

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
