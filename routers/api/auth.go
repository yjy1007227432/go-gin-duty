package api

import (
	"go-gin-duty-master/e"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/auth_service"
	"go-gin-duty-master/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	valid := validation.Validation{} //验证参数的有效性

	username := c.PostForm("username") //
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a) //首先验证参数的有效性

	//把返回值塞在Context上下文中了 每个HTTP请求都会包含一个Context对象，Context应贯穿整个HTTP请求，包含所有上下文信息
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil) //400
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}

	isExist, err := authService.Check()

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

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

//注册
func Register(c *gin.Context) {
	appG := app.Gin{C: c}
	//获取用户名、密码
	name := c.Query("name")
	password := c.Query("password")
	//判断用户是否存在
	//存在输出状态1
	//不存在创建用户，保存密码与用户名
	Bool, err := auth_service.Auth{
		Name:     name,
		Password: password,
	}.IsExistName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AUTH_FAIL, nil)
		return
	}
	if Bool {
		appG.Response(http.StatusOK, e.ERROR_EXIST_AUTH, nil)
		return
	} else {
		auth := auth_service.Auth{}
		err = c.Bind(&auth)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
			return
		}
		err = auth.AddAuth()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_ADD_AUTH_FAIL, nil)
			return
		}
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"msg": "用户创建成功",
	})

}
