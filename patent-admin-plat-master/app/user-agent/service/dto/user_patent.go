package dto

import (
	"encoding/json"
	"go-admin/app/user-agent/models"
	common "go-admin/common/models"
)

//user-patent

const (
	ClaimType = "认领"
	FocusType = "关注"
)

type UserPatentObject struct {
	UserId   int    `json:"userId" gorm:"size:128;comment:用户ID"`
	PatentId int    `form:"patentId" search:"type:exact;column:TagId;table:user_patent" comment:"专利ID" `
	Type     string `json:"type" gorm:"size:64;comment:关系类型（关注/认领）"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号"`
	//Desc     string `json:"desc" gorm:"size:128;comment:描述"`
	Properties map[string]interface{} `json:"properties" gorm:"comment:用户专利属性"`
	common.ControlBy
}

func (d *UserPatentObject) GetPatentId() interface{} {
	return d.PatentId
}

func (d *UserPatentObject) GetType() interface{} {
	return d.Type
}

func (d *UserPatentObject) GenerateUserPatent(g *models.UserPatent) {
	g.PatentId = d.PatentId
	g.UserId = d.UserId
	g.Type = d.Type
	g.PNM = d.PNM
	//g.Desc = d.Desc
	if d.Properties != nil {
		propBytes, _ := json.Marshal(d.Properties)
		g.Properties = string(propBytes)
	}
}

func NewUserPatentClaim(userId, patentId, createdBy, updatedBy int, PNM string, properties map[string]interface{}) *UserPatentObject {

	return &UserPatentObject{
		UserId:     userId,
		PatentId:   patentId,
		Type:       ClaimType,
		PNM:        PNM,
		Properties: properties,
		ControlBy: common.ControlBy{
			CreateBy: createdBy,
			UpdateBy: updatedBy,
		},
	}
}

func NewEmptyClaim() *UserPatentObject {
	return &UserPatentObject{
		Type: ClaimType,
	}
}

func NewUserPatentFocus(userId, patentId, createdBy, updatedBy int, PNM string, properties map[string]interface{}) *UserPatentObject {
	return &UserPatentObject{
		UserId:     userId,
		PatentId:   patentId,
		Type:       FocusType,
		PNM:        PNM,
		Properties: properties,
		ControlBy: common.ControlBy{
			CreateBy: createdBy,
			UpdateBy: updatedBy,
		},
	}
}

func NewEmptyFocus() *UserPatentObject {
	return &UserPatentObject{
		Type: FocusType,
	}
}

type ClaimReq struct {
	Desc string `json:"desc"`
}
