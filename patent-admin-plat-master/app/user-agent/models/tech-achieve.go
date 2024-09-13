package models

import "go-admin/common/models"

type TechAchieve struct {
	ID         int    `json:"id" gorm:"size:128;primaryKey;autoIncrement;comment:id(主键)"`
	Properties string `json:"properties" gorm:"comment:科技成果说明"`
	models.ControlBy
	models.ModelTime
}

func (e *TechAchieve) TableName() string {
	return "tech_achieve"
}

func (e *TechAchieve) Generate() models.ActiveRecord {
	return e
}

func (e *TechAchieve) GetId() interface{} {
	return e.ID
}
