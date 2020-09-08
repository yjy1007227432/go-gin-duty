package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-duty-master/e"
	"go-gin-duty-master/models"
	"go-gin-duty-master/pkg/app"
	"go-gin-duty-master/service/duty_vacation"
	"net/http"
)

func GetAllVacation(c *gin.Context) {
	var (
		vacations []models.DutyVacation
		err       error
	)
	appG := app.Gin{C: c}
	vacationService := &duty_vacation.Vacation{}

	vacations, err = vacationService.GetAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_VACATION_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"vacations": vacations,
	})
}

func DeleteAllVacation(c *gin.Context) {

	appG := app.Gin{C: c}

	vacationService := &duty_vacation.Vacation{}

	err := vacationService.DeleteAll()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_VACATION_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteVacationByName(c *gin.Context) {

	name := c.Query("name")

	appG := app.Gin{C: c}

	vacationService := &duty_vacation.Vacation{
		Name: name,
	}

	err := vacationService.DeleteByName()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_VACATION_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func EditVacationByName(c *gin.Context) {
	appG := app.Gin{C: c}

	vacation := duty_vacation.Vacation{}

	err := c.Bind(&vacation)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_BIND_DATA_FAIL, nil)
		return
	}

	vaca, err := (&duty_vacation.Vacation{
		Name: vacation.Name,
	}).GetByName()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_VACATION_FAIL, nil)
		return
	}

	vacation.Id = vaca.Id

	err = vacation.Edit()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_VACATION_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
