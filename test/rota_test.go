package test

import (
	"fmt"
	"go-gin-duty-master/service/rota_service"
	"os"
	"testing"
)

func TestImport(t *testing.T) {

	var rota = &rota_service.Rota{}

	f, _ := os.Open("C:\\Users\\Administrator\\Desktop\\test.xlsx")

	_ = rota.Import(f)

}

func TestGetThisMonth(t *testing.T) {
	var rota = &rota_service.Rota{
		Datetime: "2020-09",
	}
	rotas, _ := rota.GetThisMonth()
	fmt.Print(rotas)
}
