package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/user-agent/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
)

const (
	defaultPackageType = "claim"
)

type Package struct {
	api.Api
}

// GetPackagePages
// @Summary 获取当前用户专利包列表
// @Description 获取JSON
// @Tags 专利包
// @Router /api/v1/user-agent/package [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Param type query string true "type"
// @Security Bearer
func (e Package) GetPackagePages(c *gin.Context) {
	s := service.Package{}
	pps := service.PatentPackage{}
	ps := service.Patent{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		MakeService(&ps.Service).
		MakeService(&pps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.PackagePagesReq{}
	req.UserId = user.GetUserId(c)
	req.PageIndex, _ = strconv.Atoi(c.Query("pageIndex"))
	req.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	req.Query = c.Query("query")
	if len(c.Query("type")) != 0 {
		req.Type = c.Query("type")
	} else {
		req.Type = defaultPackageType
	}

	list := make([]models.Package, 0)
	var count int64
	err = s.GetPages(&req, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	calTotalPrice := func(patents []models.Patent) int {
		totalPrice := 0
		for _, p := range patents {
			totalPrice += p.Price
		}
		return totalPrice
	}

	for i := range list {
		patentIDs, err := pps.GetPatentIDsByPackageID(list[i].PackageId)
		if err != nil {
			e.Logger.Error(err)
			continue
		}
		list[i].PatentsNum = len(patentIDs)

		var _c int64
		patents, err := ps.GetPatentsByIds(patentIDs, &_c)
		if err != nil {
			e.Logger.Error(err)
			continue
		}
		list[i].TotalPrice = calTotalPrice(patents)
	}

	e.PageOK(list, int(count), req.PageIndex, req.PageSize, "查询成功")
}

// GetPackageListByIds
// @Summary 获取当前用户专利包列表(通过id列表)
// @Description 获取JSON
// @Tags 专利包
// @Router /api/v1/user-agent/package/ids [get]
// @Param ids query string true "ids"
// @Security Bearer
func (e Package) GetPackageListByIds(c *gin.Context) {
	s := service.Package{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	idsStr := c.Query("ids")
	var ids []int
	if err = json.Unmarshal([]byte(idsStr), &ids); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list, err := s.ListByIds(ids)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(list, "查询成功")
}

// Get
// @Summary 获取专利包
// @Description 获取JSON
// @Tags 专利包
// @Param packageId path int true "package_id"
// @Router /api/v1/user-agent/package/{package_id} [get]
// @Security Bearer
func (e Package) Get(c *gin.Context) {
	s := service.Package{}
	pps := service.PatentPackage{}
	ps := service.Patent{}
	req := dto.PackageById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		MakeService(&pps.Service).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Package

	err = s.Get(&req, &object)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	patentIDs, err := pps.GetPatentIDsByPackageID(object.PackageId)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	object.PatentsNum = len(patentIDs)

	var _c int64
	patents, err := ps.GetPatentsByIds(patentIDs, &_c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	calTotalPrice := func(patents []models.Patent) int {
		totalPrice := 0
		for _, p := range patents {
			totalPrice += p.Price
		}
		return totalPrice
	}

	object.TotalPrice = calTotalPrice(patents)

	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建专利包
// @Description 获取JSON
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PackageInsertReq true "专利包数据"
// @Router /api/v1/user-agent/package [post]
// @Security Bearer
func (e Package) Insert(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.UserID = user.GetUserId(c)
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update
// @Summary 修改专利包数据
// @Description 获取JSON
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PackageUpdateReq true "body"
// @Router /api/v1/user-agent/package/{package_id} [put]
// @Security Bearer
func (e Package) Update(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	if pid, err := strconv.Atoi(c.Param("id")); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	} else {
		req.PackageId = pid
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "更新成功")
}

// Delete
// @Summary 删除专利包
// @Description 删除专利包
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id} [delete]
// @Security Bearer
func (e Package) Delete(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageById{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置编辑人
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "删除成功")
}

//----------------------------------------patent-package---------------------------------------

// todo: please modify the swagger comment

// GetPackagePatents
// @Summary 获取指定专利包中的专利列表
// @Description 获取指定专利包中的专利列表
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id}/patent [get]
// @Security Bearer
func (e Package) GetPackagePatents(c *gin.Context) {

	s := service.PatentPackage{}
	s1 := service.Patent{}
	req := dto.PackagePageGetReq{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.PackageId, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	var list []models.PatentPackage
	var count int64

	err = s.GetPatentIdByPackageId(&req, &list, &count)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var count2 int64

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s1.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}

	res, err := s1.GetPatentsByIds(ids, &count2)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(res, int(count2), req.GetPageIndex(), req.GetPageSize(), "查询成功")

}

// IsPatentInPackage
// @Summary 查询专利是否已在专利包中
// @Description 查询专利是否已在专利包中
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id}/patent/{patent_id}/isExist [get]
// @Security Bearer
func (e Package) IsPatentInPackage(c *gin.Context) {
	var err error

	pps := service.PatentPackage{}
	req := dto.PatentPackageReq{}

	req.PNM = c.Param("PNM")

	req.PackageId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.CreateBy = user.GetUserId(c)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&pps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	existed, err := pps.IsPatentInPackage(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(&dto.IsPatentInPackageResp{Existed: existed}, "查询成功")
}

// InsertPackagePatent
// @Summary 将专利加入专利包
// @Description  将专利加入专利包
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "专利表数据"
// @Router /api/v1/user-agent/package/{package_id}/patent [post]
// @Security Bearer
func (e Package) InsertPackagePatent(c *gin.Context) {
	var err error
	pps := service.PatentPackage{}
	req := dto.PatentPackageReq{}

	ps := service.Patent{}
	patentReq := dto.PatentReq{}
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&patentReq).
		MakeService(&ps.Service).
		Errors
	patentReq.CreateBy = user.GetUserId(c)
	p, err := ps.InsertIfAbsent(&patentReq)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.PatentId = p.PatentId
	req.PNM = p.PNM
	req.PackageId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.CreateBy = user.GetUserId(c)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&pps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = pps.InsertPatentPackage(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "创建成功")
}

// DeletePackagePatent
// @Summary 删除专利包专利
// @Description  删除专利包专利
// @Tags 专利包
// @Param PatentId query string false "专利ID"
// @Param PackageId query string false "专利包ID"
// @Router /api/v1/user-agent/package/{package_id}/patent/{PNM} [delete]
// @Security Bearer
func (e Package) DeletePackagePatent(c *gin.Context) {
	s := service.PatentPackage{}
	req := dto.PackagePageGetReq{}
	req.SetUpdateBy(user.GetUserId(c))
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	packageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PackageId = packageId

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PNM = PNM

	err = s.RemovePackagePatent(&req)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.PackageBack, "删除成功")
}

// UpdatePackagePatentDesc
// @Summary 更新专利包专利描述
// @Description  更新专利包专利描述
// @Tags 专利包
// @Param data body dto.PatentDescReq true "专利描述"
// @Router /api/v1/user-agent/package/{package_id}/patent/{PNM}/desc [put]
// @Security Bearer
func (e Package) UpdatePackagePatentDesc(c *gin.Context) {
	s := service.PatentPackage{}
	req := dto.PatentDescReq{}
	req.SetUpdateBy(user.GetUserId(c))
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	packageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PackageID = packageId

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PNM = PNM

	err = s.UpdatePackagePatentDesc(&req)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req, "更新成功")
}

// InitDemoPackage
// @Summary 初始化demo package
// @Description  初始化demo package
// @Tags 专利包
// @Router /api/v1/user-agent/package/init/demo [post]
// @Security Bearer
func (e Package) InitDemoPackage(c *gin.Context) {
	s := service.Package{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)

	err = s.InitDemoPackage(userID)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "初始化成功")
}
