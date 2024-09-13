package dto

import (
	"encoding/json"
	"go-admin/app/user-agent/models"
	cDto "go-admin/common/dto"
	common "go-admin/common/models"
)

const PatentPriceBase = 8000

const (
	RelaGraphType = "rela"
	TechGraphType = "tech"

	AllClaimGraphScope = "all-claimed"
	AllFocusGraphScope = "all-focused"
	PackageScope       = "package"
)

type PatentReq struct {
	PatentId int `json:"patentId" gorm:"size:128;comment:专利ID"`
	PatentDetail
	UserProperties map[string]interface{} `json:"userProperties" gorm:"comment:用户属性"`
	common.ControlBy
}

func (s *PatentReq) GenerateList(model *models.Patent) {
	if s.PatentId != 0 {
		model.PatentId = s.PatentId
	}
	model.PNM = s.PNM
	model.ControlBy = s.ControlBy
	pbs, _ := json.Marshal(s)            //把s（json）转化为byte[]
	model.PatentProperties = string(pbs) //把byte[]转化为string
}

type PatentById struct {
	PatentId int `json:"PatentId" gorm:"size:128;comment:专利ID"`
	common.ControlBy
}

type PatentsIds struct {
	PatentId  int   `json:"patent_Id"`
	PatentIds []int `json:"patent_Ids"`
}

func (s *PatentsIds) GetPatentId() []int {
	s.PatentIds = append(s.PatentIds, s.PatentId)
	return s.PatentIds
}

type PatentBriefInfo struct {
	PatentId int    `json:"patentId" gorm:"size:128;comment:专利ID"`
	PNM      string `json:"PNM" gorm:"size:128;comment:申请号" vd:"len($)>0"`
}

type PatentDescReq struct {
	PNM       string `json:"PNM"`
	UserId    int    `json:"userId"`
	PackageID int    `json:"packageId"`
	Desc      string `json:"desc"`

	common.ControlBy
}

func (r *PatentDescReq) GeneratePatentPackage(model *models.PatentPackage) {
	model.PNM = r.PNM
	model.PackageId = r.PackageID
	model.Desc = r.Desc
}

type PatentPagesReq struct {
	cDto.Pagination
	Query string `json:"query"`
}

type FindPatentPagesReq struct {
	cDto.Pagination
	Query string `json:"query"`
}

type PatentICG struct {
	ICG   string `json:"ICG"`
	Count int    `json:"count"`
}

type PatentICGs []PatentICG

func (s PatentICGs) Len() int {
	return len(s)
}

func (s PatentICGs) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}

func (s PatentICGs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type PatentProp struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type PatentProps []PatentProp

func (s PatentProps) Len() int {
	return len(s)
}

func (s PatentProps) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}

func (s PatentProps) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
