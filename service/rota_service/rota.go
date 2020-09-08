package rota_service

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"go-gin-duty-master/models"
	"io"
	"reflect"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Rota struct {
	Id                 int
	Datetime           string
	Week               string
	BillingLate        string
	BillingWeekendLate string
	CrmLate            string
	CrmWeekendLate     string
	CrmDuty            string
	CreatedOn          time.Time
	CreatedBy          string
	ModifiedOn         time.Time
	ModifiedBy         string
}

func (t *Rota) ExistByDatetime() (bool, error) {
	return models.ExistRotaByDatetime(t.Datetime)
}

func (t *Rota) Add() error {
	var m map[string]interface{}     //声明变量，不分配内存
	m = make(map[string]interface{}) //必可不少，分配内存
	m["datetime"] = t.Datetime
	m["week"] = t.Week
	m["billing_late"] = t.BillingLate
	m["billing_weekend_late"] = t.BillingWeekendLate
	m["crm_late"] = t.CrmLate
	m["crm_weekend_late"] = t.CrmWeekendLate
	m["crm_duty"] = t.CrmDuty
	m["created_by"] = t.CreatedBy
	err := models.AddDutyRota(m)
	return err
}

func (t *Rota) GetThisMonth() ([]models.DutyRota, error) {
	var (
		rotas []models.DutyRota
		err   error
	)

	rotas, err = models.GetMonth(t.Datetime)

	if err != nil {
		return nil, err
	}

	return rotas, err
}

func (t *Rota) GetRotaByDay() (models.DutyRota, error) {
	var (
		rota models.DutyRota
		err  error
	)

	rota, err = models.GetRotaByDay(t.Datetime)

	if err != nil {
		return models.DutyRota{}, err
	}

	return rota, err
}

func (t *Rota) DeleteThisMonth() error {
	err := models.DeleteMonth(t.Datetime)
	if err != nil {
		return err
	}
	return nil
}

func (t *Rota) DeleteThisDay() error {

	err := models.DeleteDay(t.Datetime)

	if err != nil {
		return err
	}

	return nil
}

func (t *Rota) Import(r io.Reader) error {

	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	index := xlsx.GetActiveSheetIndex()

	for i := 1; i <= index; i++ {
		rows := xlsx.GetRows(xlsx.GetSheetName(i))

		if len(rows) <= 2 {
			return errors.New("表格标题格式有误")
		}

		if !correctTitle(rows[0], rows[1]) {
			return errors.New("表格标题格式有误")
		}

		for irow, row := range rows {
			if irow > 1 {
				for i := 0; i < 5; i++ {
					if row[i] == "" {
						return errors.New("表格内容不能为空")
					}
				}
				var rota = Rota{}
				if row[1] == "星期六" || row[1] == "星期日" {
					rota = Rota{
						Datetime:           convertToFormatDay(row[0]),
						Week:               row[1],
						BillingLate:        row[2],
						BillingWeekendLate: row[2],
						CrmLate:            row[3],
						CrmWeekendLate:     row[3],
						CrmDuty:            row[4],
					}
				} else {
					rota = Rota{
						Datetime:           convertToFormatDay(row[0]),
						Week:               row[1],
						BillingLate:        row[2],
						BillingWeekendLate: "",
						CrmLate:            row[3],
						CrmWeekendLate:     "",
						CrmDuty:            row[4],
					}
				}
				exists, err := rota.ExistByDatetime()

				if err != nil {
					logs.Info("查询日期值班情况失败")
				} else if exists {
					logs.Info("导入失败，表格日期值班情况已存在")
				} else {
					err = rota.Add()
				}
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func convertToFormatDay(excelDaysString string) string {
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 //在网上工具计算的天数需要加2天，什么原因没弄清楚
	curDiffDay := excelDaysString
	b, _ := strconv.Atoi(curDiffDay)
	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := b - baseDiffDay
	//fmt.Println("realDiffDay:",realDiffDay)
	// 距离2006-01-02 秒数
	realDiffSecond := realDiffDay * 24 * 3600
	//fmt.Println("realDiffSecond:",realDiffSecond)
	// 2006-01-02 15:04:05距离1970-01-01 08:00:00的秒数 网上工具可查出
	baseOriginSecond := 1136185445
	resultTime := time.Unix(int64(baseOriginSecond+realDiffSecond), 0).Format("2006-01-02")
	return resultTime
}

func correctTitle(str1, str2 []string) bool {
	title1 := []string{"日期", "星期", "计费组", "CRM/OSS组", ""}
	title2 := []string{"", "", "晚班", "晚班", "值班"}

	return reflect.DeepEqual(str1, title1) && reflect.DeepEqual(str2, title2)

}
