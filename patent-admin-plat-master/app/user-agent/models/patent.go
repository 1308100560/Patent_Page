package models

import (
	"go-admin/app/admin-agent/model"
	"go-admin/common/models"
)

type Patent struct {
	PatentId         int    `json:"patentId" gorm:"size:128;primaryKey;autoIncrement;comment:专利ID(主键)"`
	PNM              string `json:"PNM" gorm:"size:128;comment:申请号"`
	PatentProperties string `json:"patentProperties" gorm:"comment:专利详情"`
	models.ControlBy
	models.ModelTime

	UserProperties   map[string]interface{} `json:"userProperties" gorm:"-:all"`
	Price            int                    `json:"price,omitempty" gorm:"-:all"`
	EvalResult       *EvalResult            `json:"eval_result" gorm:"-:all"`
	PackageIDs       []int                  `json:"packageIDs" gorm:"-:all"`
	ReportTicketIDs  []int                  `json:"reportTicketIDs" gorm:"-:all"`
	RelaTypedReports map[string][]int       `json:"relaTypedReports" gorm:"-:all"`
}

func (Patent) TableName() string {
	return "patent"
}

func (e *Patent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Patent) GetId() interface{} {
	return e.PatentId
}

type EvalResult struct {
	EvalPrice  int           `json:"evalPrice"`
	EvalReport *model.Report `json:"evalReport"`
}
