package duty_vacation

import (
	"go-gin-duty-master/models"
	"time"
)

type Vacation struct {
	Id                   int       `json:"id"`
	Name                 string    `json:"name" binding:"required"`
	RemainVacation       float64   `json:"remain_vacation" binding:"required"`
	RemainAnnualVacation float64   `json:"remain_annual_vacation" binding:"required"`
	UpdateTime           time.Time `json:"update_time"`
}

func (t *Vacation) GetAll() ([]models.DutyVacation, error) {
	vacations, err := models.GetVacationAll()
	return vacations, err
}

func (t *Vacation) ReduceOne(vacationType int) error {
	vacation, _ := models.GetVacationByName(t.Name)
	data := make(map[string]interface{})
	switch vacationType {
	case 0:
		data["remain_vacation"] = vacation.RemainVacation - 1
		data["used_vacation"] = vacation.UsedVacation + 1
	case 1:
		data["remain_annual_vacation"] = vacation.RemainAnnualVacation - 1
		data["used_annual_vacation"] = vacation.UsedAnnualVacation + 1
	}

	err := models.EditVacation(vacation.Id, data)
	return err
}

func (t *Vacation) AddOne() error {
	vacation, _ := models.GetVacationByName(t.Name)
	data := make(map[string]interface{})
	data["remain_vacation"] = vacation.RemainVacation + 1
	data["duty_data"] = vacation.DutyData + 1
	err := models.EditVacation(vacation.Id, data)
	return err
}

func (t *Vacation) AddHalf() error {
	vacation, _ := models.GetVacationByName(t.Name)
	data := make(map[string]interface{})
	data["remain_vacation"] = vacation.RemainVacation + 0.5
	data["duty_data"] = vacation.DutyData + 0.5
	err := models.EditVacation(vacation.Id, data)
	return err
}
func (t *Vacation) AddOneHalf() error {
	vacation, _ := models.GetVacationByName(t.Name)
	data := make(map[string]interface{})
	data["remain_vacation"] = vacation.RemainVacation + 1.5
	data["duty_data"] = vacation.DutyData + 1.5
	err := models.EditVacation(vacation.Id, data)
	return err
}

func (t *Vacation) ReduceHalf(vacationType int) error {
	vacation, _ := models.GetVacationByName(t.Name)
	data := make(map[string]interface{})
	switch vacationType {
	case 0:
		data["remain_vacation"] = vacation.RemainVacation - 0.5
		data["used_vacation"] = vacation.UsedVacation + 0.5
	case 1:
		data["remain_annual_vacation"] = vacation.RemainAnnualVacation - 0.5
		data["used_annual_vacation"] = vacation.UsedAnnualVacation + 0.5
	}
	err := models.EditVacation(vacation.Id, data)
	return err
}

func (t *Vacation) GetByName() (models.DutyVacation, error) {
	vacation, err := models.GetVacationByName(t.Name)
	return vacation, err
}

func (t *Vacation) DeleteAll() error {
	err := models.DeleteVacationAll()
	return err
}

func (t *Vacation) DeleteByName() error {
	err := models.DeleteVacationByName(t.Name)
	return err
}

func (t *Vacation) AddRemainVacation() error {
	data := make(map[string]interface{})
	data["remain_vacation"] = t.RemainVacation + 1
	err := models.EditVacation(t.Id, data)
	return err
}

func (t *Vacation) AddRemainAnnualVacation() error {
	data := make(map[string]interface{})
	data["remain_annual_vacation"] = t.RemainAnnualVacation + 1
	err := models.EditVacation(t.Id, data)
	return err
}

func (t *Vacation) Edit() error {
	data := make(map[string]interface{})
	data["remain_vacation"] = t.RemainVacation
	data["remain_annual_vacation"] = t.RemainAnnualVacation
	data["update_time"] = time.Now()
	return models.EditVacation(t.Id, data)
}

func (t *Vacation) Add(quantity float64) error {
	vacation, _ := models.GetVacationByName(t.Name)
	data := make(map[string]interface{})
	data["remain_vacation"] = vacation.RemainVacation + quantity
	data["overtime_data"] = vacation.OvertimeData + quantity
	err := models.EditVacation(vacation.Id, data)
	return err
}
