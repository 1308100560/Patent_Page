package models

import "go-admin/common/models"

type Package struct {
	PackageId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"packageId"`
	PackageName string `json:"packageName" gorm:"size:128;comment:专利包"`
	//Desc        string `json:"desc" gorm:"size:128;comment:描述"`
	Properties string `json:"properties" gorm:"comment:专利包属性"`
	Files      string `json:"files" gorm:"comment:专利包附件"`
	Type       string `json:"type" gorm:"size:128;comment:专利包类型"`
	models.ControlBy
	models.ModelTime

	PatentsNum int `json:"patentsNum" gorm:"-:all"`
	TotalPrice int `json:"totalPrice" gorm:"-:all"`
}

func (e *Package) TableName() string {
	return "package"
}

func (e *Package) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Package) GetId() interface{} {
	return e.PackageId
}
