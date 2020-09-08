package api

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

//@Summary Get Auth
//@Produce  json
//@Param username query string true "userName"
//@Param password query string true "password"
//@Success 200 {object} app.Response
//@Failure 500 {object} app.Response
//@Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	valid := validation.Validation{} //验证参数的有效性

	username := c.Query("username") //
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, err := valid.Valid(&a) //首先验证参数的有效性

	//把返回值塞在Context上下文中了 每个HTTP请求都会包含一个Context对象，Context应贯穿整个HTTP请求，包含所有上下文信息
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil) //400
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}

	isExist, _ := authService.Check()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	fmt.Print(token)
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
