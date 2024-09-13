package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"go-admin/app/user-agent/utils"
	"sort"
	"strconv"
	"strings"
)

type Dashboard struct {
	api.Api
}

// GetDashboard
// @Summary Dashboard主页
// @Description  Dashboard主页
// @Tags Dashboard主页
// @Success 200 {object} dto.Dashboard
// @Router /apis/v1/user-agent/dashboard [get]
// @Security Bearer
func (e Dashboard) GetDashboard(c *gin.Context) {
	ups := service.UserPatent{}
	ps := service.Patent{}
	pks := service.Package{}
	rps := service.Report{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&ups.Service).
		MakeService(&ps.Service).
		MakeService(&pks.Service).
		MakeService(&rps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)

	focusList := make([]models.UserPatent, 0)
	if err = ups.GetFocusLists(userID, &focusList); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	focusCount := len(focusList)

	// get competitors
	focusIds := make([]int, len(focusList))
	for _, fp := range focusList {
		focusIds = append(focusIds, fp.Id)
	}
	var _c int64
	focusPatents, err := ps.GetPatentsByIds(focusIds, &_c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	competitorNodes, _, err := ps.FindInventorsAndRelationsFromPatents(focusPatents)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	claimList := make([]models.UserPatent, 0)
	if err = ups.GetClaimLists(userID, &claimList); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	claimCount := len(claimList)

	// get collaborators
	claimIds := make([]int, len(claimList))
	for _, cp := range claimList {
		claimIds = append(claimIds, cp.PatentId)
	}
	var count int64
	claimPatents, err := ps.GetPatentsByIds(claimIds, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	collaboratorNodes, _, err := ps.FindInventorsAndRelationsFromPatents(claimPatents)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	claimINNMap := make(map[string]int)
	claimPAMap := make(map[string]int)
	for _, p := range claimPatents {
		detail := &dto.PatentDetail{}
		if err = json.Unmarshal([]byte(p.PatentProperties), detail); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		if detail.Inn != "" {
			names := strings.Split(detail.Inn, ";")
			for _, name := range names {
				if utils.IsChinese(name) {
					if ic, ok := claimINNMap[name]; ok {
						claimINNMap[name] = ic + 1
					} else {
						claimINNMap[name] = 1
					}
				}
			}
		}
		if detail.Pa != "" {
			if ic, ok := claimPAMap[detail.Pa]; ok {
				claimPAMap[detail.Pa] = ic + 1
			} else {
				claimPAMap[detail.Pa] = 1
			}
		}
	}

	claimINNs := make(dto.PatentProps, 0)
	for k, v := range claimINNMap {
		claimINNs = append(claimINNs, dto.PatentProp{
			Name:  k,
			Count: v,
		})
	}
	sort.Sort(claimINNs)
	if len(claimINNs) > 10 {
		claimINNs = claimINNs[:10]
	}

	claimPAs := make(dto.PatentProps, 0)
	for k, v := range claimPAMap {
		claimPAs = append(claimPAs, dto.PatentProp{
			Name:  k,
			Count: v,
		})
	}
	sort.Sort(claimPAs)
	if len(claimPAs) > 10 {
		claimPAs = claimPAs[:10]
	}

	ids := make([]int, 0, len(claimList))
	for _, claim := range claimList {
		ids = append(ids, claim.PatentId)
	}
	var patentCount int64
	patents, err := ps.GetPatentsByIds(ids, &patentCount)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	patentStatus := make(map[string]int)
	publicationDates := dto.NewPublicationDates()
	var totalPrice int
	for _, p := range patents {
		pd := dto.PatentDetail{}
		if err = json.Unmarshal([]byte(p.PatentProperties), &pd); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		if len(pd.Cls) == 0 {
			continue
		}
		if _, ok := patentStatus[pd.Cls]; ok {
			patentStatus[pd.Cls]++
		} else {
			patentStatus[pd.Cls] = 1
		}

		if len(pd.Ad) != 0 {
			publicationDates.AddYear(pd.Ad)
		}

		totalPrice += pd.Idx * dto.PatentPriceBase
	}

	pkgList := make([]models.Package, 0)
	listPkgReq := dto.PackageListReq{UserId: user.GetUserId(c)}
	err = pks.ListByUserId(&listPkgReq, &pkgList)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	pkgCount := 0
	for _, pkg := range pkgList {
		if pkg.Type == dto.PackageClaimType {
			pkgCount++
		}
	}

	finishedReports, err := rps.GetReportListByUidAndTicketStatus(userID, dtos.TicketStatusFinished)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	reportStatus := make(map[string]int)
	for _, fp := range finishedReports {
		if num, ok := reportStatus[fp.Type]; ok {
			reportStatus[fp.Type] = num + 1
		} else {
			reportStatus[fp.Type] = 1
		}
	}

	res := &dto.Dashboard{
		PatentClaimCount:     claimCount,
		PatentFocusCount:     focusCount,
		PatentStatus:         patentStatus,
		PublicationDates:     publicationDates.List(),
		PackageCount:         pkgCount,
		PatentRecommendation: nil,
		ReportStatus:         reportStatus,
		PatentTotalPrice:     totalPrice,
		Collaborators:        covertNodesToResearchers(collaboratorNodes),
		Competitors:          covertNodesToResearchers(competitorNodes),
		ClaimApartments:      claimPAs,
		ClaimInventors:       claimINNs,
	}

	e.OK(res, "查询成功")
}

// GetPackageDashboard
// @Summary Package Dashboard主页
// @Description  Package Dashboard主页
// @Tags Dashboard主页
// @Success 200 {object} dto.Dashboard
// @Router /apis/v1/user-agent/dashboard/package/{pid} [get]
// @Security Bearer
func (e Dashboard) GetPackageDashboard(c *gin.Context) {
	pps := service.PatentPackage{}
	ps := service.Patent{}
	rs := service.Report{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&pps.Service).
		MakeService(&ps.Service).
		MakeService(&rs.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)
	pid, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	patentIDs, err := pps.GetPatentIDsByPackageID(pid)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var _c int64
	packagePatents, err := ps.GetPatentsByIds(patentIDs, &_c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	collaboratorNodes, _, err := ps.FindInventorsAndRelationsFromPatents(packagePatents)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	INNMap := make(map[string]int)
	PAMap := make(map[string]int)
	patentStatus := make(map[string]int)
	publicationDates := dto.NewPublicationDates()
	var totalPrice int
	for _, p := range packagePatents {
		detail := &dto.PatentDetail{}
		if err = json.Unmarshal([]byte(p.PatentProperties), detail); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		if detail.Inn != "" {
			names := strings.Split(detail.Inn, ";")
			for _, name := range names {
				if utils.IsChinese(name) {
					if ic, ok := INNMap[name]; ok {
						INNMap[name] = ic + 1
					} else {
						INNMap[name] = 1
					}
				}
			}
		}
		if detail.Pa != "" {
			if ic, ok := PAMap[detail.Pa]; ok {
				PAMap[detail.Pa] = ic + 1
			} else {
				PAMap[detail.Pa] = 1
			}
		}

		if len(detail.Cls) == 0 {
			continue
		}
		if _, ok := patentStatus[detail.Cls]; ok {
			patentStatus[detail.Cls]++
		} else {
			patentStatus[detail.Cls] = 1
		}

		if len(detail.Ad) != 0 {
			publicationDates.AddYear(detail.Ad)
		}

		totalPrice += detail.Idx * dto.PatentPriceBase
	}

	packageINNs := make(dto.PatentProps, 0)
	for k, v := range INNMap {
		packageINNs = append(packageINNs, dto.PatentProp{
			Name:  k,
			Count: v,
		})
	}
	sort.Sort(packageINNs)
	if len(packageINNs) > 10 {
		packageINNs = packageINNs[:10]
	}

	packagePAs := make(dto.PatentProps, 0)
	for k, v := range PAMap {
		packagePAs = append(packagePAs, dto.PatentProp{
			Name:  k,
			Count: v,
		})
	}
	sort.Sort(packagePAs)
	if len(packagePAs) > 10 {
		packagePAs = packagePAs[:10]
	}

	reports, err := rs.GetReportListByPatentIDs(patentIDs, userID)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	reportStatus := make(map[string]int)
	for _, fp := range reports {
		if num, ok := reportStatus[fp.Type]; ok {
			reportStatus[fp.Type] = num + 1
		} else {
			reportStatus[fp.Type] = 1
		}
	}

	res := &dto.PackageDashboard{
		PatentCount:       len(patentIDs),
		PatentStatus:      patentStatus,
		PublicationDates:  publicationDates.List(),
		ReportStatus:      reportStatus,
		PatentTotalPrice:  totalPrice,
		Collaborators:     covertNodesToResearchers(collaboratorNodes),
		PackageInventors:  packageINNs,
		PackageApartments: packagePAs,
	}

	e.OK(res, "查询成功")
}

func covertNodesToResearchers(nodes []models.SimplifiedNode) []*dto.Researcher {
	res := make([]*dto.Researcher, 0, len(nodes))
	for _, n := range nodes {
		res = append(res, &dto.Researcher{
			Name:  n.Name,
			Times: n.TheNumberOfPatents,
		})
	}
	return res
}
