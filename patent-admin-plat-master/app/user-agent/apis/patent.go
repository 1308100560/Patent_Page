package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/app/admin-agent/service/dtos"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"net/http"
	"strconv"
)

type Patent struct {
	api.Api
}

//----------------------------------------patent----------------------------------------

// GetPatentById
// @Summary 检索专利
// @Description  通过PatentId检索专利
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/{patent_id} [get]
// @Security Bearer
func (e Patent) GetPatentById(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentById{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object2 models.Patent
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	req.PatentId, err = strconv.Atoi(c.Param("patent_id"))
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "not found params from router")
		return
	}
	err = s.Get(&req, &object2)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object2, "查询成功")
}

// GetPatentLists
// @Summary 列表专利信息数据
// @Description 获取本地专利
// @Tags 专利表
// @Router /api/v1/user-agent/patent [get]
// @Security Bearer
func (e Patent) GetPatentLists(c *gin.Context) { //gin框架里的上下文
	s := service.Patent{}  //service中查询或者返回的结果赋值给s变量
	req := dto.PatentReq{} //被绑定的数据
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

	list := make([]models.Patent, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

// UpdatePatent
// @Summary 修改专利
// @Description 必须要有主键PatentId值
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "body"
// @Router /api/v1/user-agent/patent [put]
// @Security Bearer
func (e Patent) UpdatePatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentReq{}
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

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.UpdateLists(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req, "更新成功")
}

// DeletePatent
// @Summary 删除专利
// @Description  输入专利id删除专利表
// @Tags 专利表
// @Param PatentId query string false "专利ID"
// @Router /api/v1/user-agent/patent/{patent_id} [delete]
// @Security Bearer
func (e Patent) DeletePatent(c *gin.Context) {
	s := service.Patent{}
	req := dto.PatentById{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.PatentId, err = strconv.Atoi(c.Param("patent_id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.UpdateBy = user.GetUserId(c)

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req, "删除成功")
}

//----------------------------------------user-patent-----------------------------------------------------------------

// GetUserPatentsPages
// @Summary 获取用户的专利列表
// @Description 获取用户的专利列表
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/user [get]
// @Security Bearer
// todo: remove redundant
func (e Patent) GetUserPatentsPages(c *gin.Context) {

	s := service.UserPatent{}
	s1 := service.Patent{}
	req := dto.UserPatentObject{}

	req.UserId = user.GetUserId(c)

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

	list := make([]models.UserPatent, 0)

	var count int64

	err = s.GetUserPatentIds(&req, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var count2 int64
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}

	res, err := s1.GetPatentsByIds(ids, &count2)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(res, "查询成功")
}

// ClaimPatent
// @Summary 认领专利
// @Description 认领专利
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "Type和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/claim [post]
// @Security Bearer
func (e Patent) ClaimPatent(c *gin.Context) {

	pid, PNM, properties, err := e.internalInsertIfAbsent(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.UserPatent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentClaim(user.GetUserId(c), pid, user.GetUserId(c), user.GetUserId(c), PNM, properties)

	if err = s.InsertUserPatent(req); err != nil {
		e.Logger.Error(err)
		if errors.Is(err, service.ErrConflictBindPatent) {
			e.Error(409, err, err.Error())
		} else {
			e.Error(500, err, err.Error())
		}
		return
	}

	e.OK(req, "认领成功")
}

// FocusPatent
// @Summary 关注专利
// @Description 关注专利
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "Type和PatentId为必要输入"
// @Router /api/v1/user-agent/patent/focus [post]
// @Security Bearer
func (e Patent) FocusPatent(c *gin.Context) {

	pid, PNM, desc, err := e.internalInsertIfAbsent(c)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	s := service.UserPatent{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentFocus(user.GetUserId(c), pid, user.GetUserId(c), user.GetUserId(c), PNM, desc)

	if err = s.InsertUserPatent(req); err != nil {
		e.Logger.Error(err)
		if errors.Is(err, service.ErrConflictBindPatent) {
			e.Error(409, err, err.Error())
		} else {
			e.Error(500, err, err.Error())
		}
		return
	}

	e.OK(req, "关注成功")
}

// InsertIfAbsent
// @Summary 添加专利
// @Description 添加专利到本地
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "专利表数据"
// @Router /api/v1/user-agent/patent [post]
// @Security Bearer
func (e Patent) InsertIfAbsent(c *gin.Context) {
	pid, pnm, _, err := e.internalInsertIfAbsent(c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = e.MakeContext(c).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(&dto.PatentBriefInfo{PatentId: pid, PNM: pnm}, "success")
}

func (e Patent) internalInsertIfAbsent(c *gin.Context) (int, string, map[string]interface{}, error) {
	ps := service.Patent{}
	req := dto.PatentReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		return 0, "", nil, err
	}
	req.CreateBy = user.GetUserId(c)
	p, err := ps.InsertIfAbsent(&req)
	if err != nil {
		return 0, "", nil, err
	}
	return p.PatentId, p.PNM, req.UserProperties, nil
}

// GetFocusPages
// @Summary 获取关注列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/focus [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Security Bearer
func (e Patent) GetFocusPages(c *gin.Context) {
	ups := service.UserPatent{}
	ps := service.Patent{}
	pkgs := service.Package{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&ups.Service).
		MakeService(&pkgs.Service).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.UserPatent, 0)
	userID := user.GetUserId(c)
	err = ups.GetFocusLists(userID, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dto.PatentPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Query = c.Query("query")

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}
	sort.Ints(ids)
	var count int64
	res, err := ps.GetPatentPagesByIds(ids, req, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	packageIDs, err := pkgs.GetPackageIDsByPatentIDs(ids, userID)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	for i := range res {
		res[i].UserProperties = make(map[string]interface{})
		json.Unmarshal([]byte(list[i].Properties), &res[i].UserProperties)
		if pkgIDs, ok := packageIDs[res[i].PatentId]; ok {
			res[i].PackageIDs = pkgIDs
		}
	}

	e.PageOK(res, int(count), req.PageSize, req.PageIndex, "查询成功")
}

// GetClaimPages
// @Summary 获取认领列表
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/claim [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Param needEvalResult query string true "needEvalResult"
// @Security Bearer
func (e Patent) GetClaimPages(c *gin.Context) {
	s := service.UserPatent{}
	ps := service.Patent{}
	pkgs := service.Package{}
	rs := service.Report{}
	userID := user.GetUserId(c)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		MakeService(&ps.Service).
		MakeService(&pkgs.Service).
		MakeService(&rs.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.UserPatent, 0)
	relaMap := make(map[int]models.UserPatent)
	err = s.GetClaimLists(userID, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dto.PatentPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Query = c.Query("query")

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
		relaMap[list[i].PatentId] = list[i]
	}

	var count int64
	res, err := ps.GetPatentPagesByIds(ids, req, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	packageIDs, err := pkgs.GetPackageIDsByPatentIDs(ids, userID)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	reportTicketIDs, err := rs.GetReportTicketIDsByPatentIDs(ids, userID)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	for i := range res {
		res[i].UserProperties = make(map[string]interface{})
		json.Unmarshal([]byte(relaMap[res[i].PatentId].Properties), &res[i].UserProperties)
		if pkgIDs, ok := packageIDs[res[i].PatentId]; ok {
			res[i].PackageIDs = pkgIDs
		}
		if rpIDs, ok := reportTicketIDs[res[i].PatentId]; ok {
			res[i].ReportTicketIDs = rpIDs

			// find eval result for patents... (sb jiafang)
			if len(rpIDs) != 0 {
				reports, err := rs.GetReportListByTicketIDs(rpIDs, userID)
				if err != nil {
					e.Logger.Error(err)
					continue
				}
				res[i].RelaTypedReports = make(map[string][]int)
				for _, report := range reports {
					if _, ok = res[i].RelaTypedReports[report.Type]; !ok {
						res[i].RelaTypedReports[report.Type] = []int{report.ReportId}
					} else {
						res[i].RelaTypedReports[report.Type] = append(res[i].RelaTypedReports[report.Type],
							report.ReportId)
					}
					if report.Type == dtos.ReportTypeEval && report.ReportProperties != "" {
						data := make(map[string]interface{})
						if err = json.Unmarshal([]byte(report.ReportProperties), &data); err != nil {
							continue
						}
						if _, ok := data[dtos.EvalPriceKey]; ok {
							evalPrice, _ := strconv.Atoi(data[dtos.EvalPriceKey].(string))
							if evalPrice != 0 {
								res[i].EvalResult = &models.EvalResult{
									EvalPrice:  int(evalPrice),
									EvalReport: &report,
								}
								break
							}
						}
					}
				}
			}
		}
	}

	e.PageOK(res, int(count), req.PageIndex, req.PageSize, "查询成功")
}

// ICGAnalyseInClaim
// @Summary 专利领域分析
// @Description
// @Tags 专利表
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/patent/claim/analyse/icg [get]
// @Security Bearer
func (e Patent) ICGAnalyseInClaim(c *gin.Context) {
	s := service.UserPatent{}
	ps := service.Patent{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.UserPatent, 0)
	userID := user.GetUserId(c)
	err = s.GetClaimLists(userID, &list)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ids := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ids[i] = list[i].PatentId
	}

	var count int64
	patents, err := ps.GetPatentsByIds(ids, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ICGMaps := make(map[string]int)
	for _, p := range patents {
		detail := &dto.PatentDetail{}
		if err = json.Unmarshal([]byte(p.PatentProperties), detail); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		if detail.Icg != "" {
			icgList := strings.Split(detail.Icg, ";")
			for _, icg := range icgList {
				if ic, ok := ICGMaps[icg]; ok {
					ICGMaps[icg] = ic + 1
				} else {
					ICGMaps[icg] = 1
				}
			}
		}
	}

	icgs := make(dto.PatentICGs, 0, len(ICGMaps))
	for icg, ic := range ICGMaps {
		icgs = append(icgs, dto.PatentICG{
			ICG:   icg,
			Count: ic,
		})
	}
	sort.Sort(icgs)
	if len(icgs) > 3 {
		icgs = icgs[:3]
	}

	e.OK(icgs, "查询成功")
}

// DeleteFocus
// @Summary 取消关注
// @Description  取消关注
// @Tags 专利表
// @Param PNM query string false "专利PNM"
// @Router /api/v1/user-agent/patent/focus/{PNM}  [delete]
// @Security Bearer
func (e Patent) DeleteFocus(c *gin.Context) {
	var err error
	s := service.UserPatent{}
	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentFocus(user.GetUserId(c), -1, user.GetUserId(c), user.GetUserId(c), PNM, nil)

	err = e.MakeContext(c).
		MakeOrm().
		Bind(req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.RemoveFocus(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req, "取消关注成功")
}

// DeleteClaim
// @Summary 取消认领
// @Description  取消认领
// @Tags 专利表
// @Param PNM query string false "专利PNM"
// @Router /api/v1/user-agent/patent/claim/{PNM} [delete]
// @Security Bearer
func (e Patent) DeleteClaim(c *gin.Context) {
	var err error
	s := service.UserPatent{}

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewUserPatentClaim(user.GetUserId(c), -1, user.GetUserId(c), user.GetUserId(c), PNM, nil)

	err = e.MakeContext(c).
		MakeOrm().
		Bind(req). //修改&
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.RemoveClaim(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "取消认领成功")
}

// UpdateClaimProperties
// @Summary 更新认领专利属性
// @Description  更新认领专利属性
// @Tags 专利表
// @Param data body dto.PatentDescReq true "专利描述"
// @Router /api/v1/user-agent/patent/claim/{PNM} [put]
// @Security Bearer
func (e Patent) UpdateClaimProperties(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.NewEmptyClaim()
	req.UserId = user.GetUserId(c)
	req.SetUpdateBy(user.GetUserId(c))
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PNM = PNM

	err = s.UpdateUserPatentDesc(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "更新成功")
}

// UpdateFocusProperties
// @Summary 更新关注专利属性
// @Description  更新关注专利属性
// @Tags 专利表
// @Param data body dto.PatentDescReq true "专利描述"
// @Router /api/v1/user-agent/patent/focus/{PNM} [put]
// @Security Bearer
func (e Patent) UpdateFocusProperties(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.NewEmptyFocus()
	req.UserId = user.GetUserId(c)
	req.SetUpdateBy(user.GetUserId(c))
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req).
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	PNM := c.Param("PNM")
	if len(PNM) == 0 {
		err = fmt.Errorf("PNM should be provided in path")
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PNM = PNM

	err = s.UpdateUserPatentDesc(req)

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req, "更新成功")
}

// GetPatentGraph
// @Summary 获取专利的图谱
// @Description  获取专利的图谱
// @Tags 专利表
// @Router /api/v1/user-agent/patent/graph [get]
// @Param type query string true "type"
// @Param scope query string true "scope"
// @Param packageId query int true "packageId"
// @Security Bearer
func (e Patent) GetPatentGraph(c *gin.Context) {
	ps := service.Patent{}
	sup := service.UserPatent{}
	pps := service.PatentPackage{}
	var err error
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		MakeService(&pps.Service).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	graphType := c.Query("type")
	graphScope := c.Query("scope")
	packageId, _ := strconv.Atoi(c.Query("packageId"))

	if err = checkGraphParams(graphType, graphScope, packageId); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var ids []int
	userID := user.GetUserId(c)
	switch graphScope {
	case dto.AllClaimGraphScope:
		upList := make([]models.UserPatent, 0)
		err = sup.GetClaimLists(userID, &upList)
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		ids = make([]int, len(upList))
		for i := 0; i < len(upList); i++ {
			ids[i] = upList[i].PatentId
		}
	case dto.AllFocusGraphScope:
		upList := make([]models.UserPatent, 0)
		err = sup.GetFocusLists(userID, &upList)
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		ids = make([]int, len(upList))
		for i := 0; i < len(upList); i++ {
			ids[i] = upList[i].PatentId
		}
	case dto.PackageScope:
		ids, err = pps.GetPatentIDsByPackageID(packageId)
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
	}

	var count int64
	var patents []models.Patent
	patents, err = ps.GetPatentsByIds(ids, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var nodes []models.SimplifiedNode
	var relations []int
	switch graphType {
	case dto.RelaGraphType:
		nodes, relations, err = ps.FindInventorsAndRelationsFromPatents(patents) //relations is an Upper Triangle
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
	case dto.TechGraphType:
		nodes, relations, err = ps.FindKeywordsAndRelationsFromPatents(patents) //relations is an Upper Triangle
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
	}

	graph, err := ps.GetGraphByPatents(nodes, relations)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(graph, "查询成功")
}

func checkGraphParams(graphType string, graphScope string, packageId int) error {
	if graphType != dto.RelaGraphType && graphType != dto.TechGraphType {
		return fmt.Errorf("invalid graph parameters")
	}

	if graphScope != dto.AllFocusGraphScope && graphScope != dto.AllClaimGraphScope && graphScope != dto.PackageScope {
		return fmt.Errorf("invalid graph parameters")
	}

	if graphScope == dto.PackageScope && packageId == 0 {
		return fmt.Errorf("invalid graph parameters")
	}

	return nil
}
