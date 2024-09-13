package dto

import (
	"encoding/json"
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
)

type TechAchieveReq struct {
	ID         int                    `json:"id"`
	UserID     int                    `json:"-"`
	Properties map[string]interface{} `json:"properties"`
}

func (r *TechAchieveReq) Generate(record *models.TechAchieve) {
	data, _ := json.Marshal(r.Properties)
	record.ID = r.ID
	record.Properties = string(data)
	record.CreateBy = r.UserID
}

type TechAchievePagesReq struct {
	Query string `json:"query"`

	dto.Pagination
	TechAchieveReq
}
