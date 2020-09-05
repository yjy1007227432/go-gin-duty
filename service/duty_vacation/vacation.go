package duty_vacation

import (
	"go-gin-duty-master/models"
	"time"
)

type Vacation struct {
	Id                   int
	Name                 string
	RemainVacation       int
	RemainAnnualVacation int
	UpdateTime           time.Time
}

func (t *Vacation) GetAll() ([]models.DutyVacation, error) {
	vacations, err := models.GetVacationAll()
	return vacations, err
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
