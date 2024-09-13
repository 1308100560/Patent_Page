package dto

import (
	"encoding/json"
	"go-admin/app/other/apis"
	"go-admin/app/user-agent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

const (
	FilesAdd    = "add"
	FilesDelete = "del"
)

const (
	PackageClaimType = "claim"
	PackageFocusType = "focus"
)

type PackageListReq struct {
	UserId int `json:"-" form:"desc" search:"type:order;column:created_at;table:package"`
}

type PackagePagesReq struct {
	dto.Pagination
	UserId int    `json:"-" form:"desc" search:"type:order;column:created_at;table:package"`
	Query  string `json:"query"`
	Type   string `json:"type"`
}

type PackageFindReq struct {
	UserId int    `json:"-" form:"desc" search:"type:order;column:created_at;table:package"`
	Query  string `json:"query"`
}

type PackageInsertReq struct {
	PackageId   int                    `json:"packageId" comment:"专利包ID"` // 专利包ID
	PackageName string                 `json:"packageName" comment:"专利包名" vd:"len($)>0"`
	Type        string                 `json:"type" comment:"专利包类型" vd:"len($)>0"`
	Properties  map[string]interface{} `json:"properties" comment:"专利包属性"`
	UserID      int                    `json:"-"`
	//common.ControlBy
}

func (s *PackageInsertReq) Generate(model *models.Package) {
	if s.PackageId != 0 {
		model.PackageId = s.PackageId
	}
	model.PackageName = s.PackageName
	model.Type = s.Type
	model.CreateBy = s.UserID

	if s.Properties != nil {
		propsBytes, _ := json.Marshal(s.Properties)
		model.Properties = string(propsBytes)
	}
}

func (s *PackageInsertReq) GetId() interface{} {
	return s.PackageId
}

type PackageUpdateReq struct {
	PackageId   int                    `json:"packageId" comment:"专利包ID"` // 专利包ID
	PackageName string                 `json:"packageName" comment:"专利包名"`
	Properties  map[string]interface{} `json:"properties" comment:"专利包属性"`
	FilesOpt    string                 `json:"filesOpt" comment:"文件操作"`
	Files       []apis.FileResponse    `json:"files" comment:"专利包附件"`
	common.ControlBy
}

func (s *PackageUpdateReq) Generate(model *models.Package) {
	if s.PackageId != 0 {
		model.PackageId = s.PackageId
	}
	model.PackageName = s.PackageName
	if s.Properties != nil {
		propsBytes, _ := json.Marshal(s.Properties)
		model.Properties = string(propsBytes)
	}
}

func (s *PackageUpdateReq) GenerateAndAddFiles(model *models.Package) {
	s.Generate(model)
	if len(model.Files) == 0 {
		fbs, _ := json.Marshal(s.Files)
		model.Files = string(fbs)
	} else {
		files := make([]apis.FileResponse, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)
		files = append(files, s.Files...)
		fbs, _ := json.Marshal(files)
		model.Files = string(fbs)
	}
}

func (s *PackageUpdateReq) GenerateAndDeleteFiles(model *models.Package) {
	s.Generate(model)
	if len(model.Files) != 0 {
		files := make([]apis.FileResponse, 0)
		_ = json.Unmarshal([]byte(model.Files), &files)

		needToDel := make(map[string]struct{})
		for _, df := range s.Files {
			needToDel[df.FullPath] = struct{}{}
		}

		slow := 0
		for _, f := range files {
			if _, ok := needToDel[f.FullPath]; !ok {
				files[slow] = f
				slow++
			}
		}
		files = files[:slow]
		fbs, _ := json.Marshal(files)
		model.Files = string(fbs)
	}
}

func (s *PackageUpdateReq) GetId() interface{} {
	return s.PackageId
}

type PackageById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *PackageById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

func GetDemoPackages() []string {
	return []string{
		"石油化工装置工艺设计包内容",
		"设计基础",
		"工艺说明",
		"物料平衡",
		"消耗量",
		"界区条件表",
		"卫生、安全、环保说明",
		"分析化验项目表",
		"工艺管道及仪表流程图 (PID)",
		"建议的设备布置图及说明",
		"工艺设备表",
		"工艺设备",
		"自控仪表",
		"特殊管道",
		"主要安全泄放设施数据表",
		"有关专利文件目录",
		"工艺手册",
		"工艺说明",
		"正常操作程序",
		"开车准备工作程序",
		"开车程序",
		"正常停车程序",
		"事故处理原则",
		"催化剂装卸",
		"采样",
		"工艺危险因素分析及控制措施",
		"环境保护",
		"设备检查与维护",
		"分析化验手册",
	}
}
