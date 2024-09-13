package models

import (
	cModel "go-admin/common/models"
)

type TradingRecord struct {
	ID         int    `json:"id" gorm:"size:128;primaryKey;autoIncrement;comment:工单ID(主键)"`
	PatentID   int    `json:"patentID" gorm:"size:128;comment:专利ID"`
	Properties string `json:"properties" gorm:"comment:交易记录详情"`
	cModel.ModelTime

	Patent *Patent `json:"patent" gorm:"-"`
}
