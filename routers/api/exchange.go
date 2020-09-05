package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/exchange_service"
	"net/http"
)

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
