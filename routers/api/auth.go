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

// @Summary 获取用户名字
// @Produce  json
// @Param token query string true "token"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} string	 "{"code":500,"data":{},"msg":"ok"}"
// @Router /api/auth/getName   [post]
func GetNameByToken(c *gin.Context) {
	appG := app.Gin{C: c}

	token := c.Query("token") //

	username, err := util.DecrpytToken(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GENERATE_TOKEN_FAIL, nil)
		return
	}
	auth := auth_service.Auth{Username: username}

	name, _, err := auth.GetNameByUsername()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_NAME_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"name": name,
	})
}

// @Summary 新增个人用户
// @Param token query string true "token"
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param name query string true "姓calculate名"
// @Param group query string true "组：crm/"
// @Param telephone query string true "电话"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} string	 "{"code":500,"data":{},"msg":"ok"}"
// @Router /api/auth/AddAuth   [post]
func AddAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	auth := auth_service.Auth{}

	err := c.Bind(&auth)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
		return
	}

	IsExist, err := auth.IsExistUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_AUTH_FAIL, nil)
		return
	}

	if IsExist == true {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_USER_AUTH, nil)
		return
	}

	err = auth.AddAuth()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_AUTH_FAIL, nil)
		return
	}

}

// @Summary 通过用户名密码获取token
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth   [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	valid := validation.Validation{} //验证参数的有效性

	username := c.Query("username") //
	password := c.Query("password")
	//参数验证
	if username == "" || password == "" {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
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

// @Summary 注册个人用户
// @Produce  json
// @Param token query string true "token"
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param name query string true "名字"
// @Param group query string true "组：crm/calculate"
// @Param telephone query string true "电话"
// @Success 200 {string} string	 "{"code":200,"data":{},"msg":"ok"}"
// @Router /register   [post]
func Register(c *gin.Context) {
	appG := app.Gin{C: c}
	//获取用户名、密码
	name := c.Query("name")
	password := c.Query("password")
	//判断用户是否存在
	//存在输出状态1
	//不存在创建用户，保存密码与用户名
	Bool, err := (&auth_service.Auth{
		Name:     name,
		Password: password,
	}).IsExistName()
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
