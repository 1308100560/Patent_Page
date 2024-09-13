package dto

import (
	"encoding/json"
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
)

type TradingRecordReq struct {
	ID         int                    `json:"id"`
	PatentID   int                    `json:"patentID"`
	Properties map[string]interface{} `json:"properties"`
}

func (r *TradingRecordReq) Generate(record *models.TradingRecord) {
	data, _ := json.Marshal(r.Properties)
	record.ID = r.ID
	record.PatentID = r.PatentID
	record.Properties = string(data)
}

type TradingRecordPagesReq struct {
	Query     string `json:"query"`
	PatentIDS []int  `json:"patentIDs"`

	dto.Pagination
	TradingRecordReq
}
