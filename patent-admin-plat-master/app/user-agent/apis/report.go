package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	"strconv"
)

type Report struct {
	api.Api
}

// GetReportPages
// @Summary 获取报告列表
// @Description 获取报告列表
// @Tags 用户-报告
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/reports [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param reportType query string true "reportType"
// @Param query query string true "query"
// @Security Bearer
func (e Report) GetReportPages(c *gin.Context) {
	s := service.Ticket{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dtos.TicketListReq{}
	req.Type = dtos.TicketTypeReport
	req.Status = dtos.TicketStatusFinished

	list := make([]model.Ticket, 0)
	var count int64
	if err = s.GetTicketList(&req, &list, &count); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	rs := service.Report{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&rs.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	rt := c.Query("reportType")
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	rtr := dtos.ReportPagesReq{}
	rtr.PageIndex = pageIndex
	rtr.PageSize = pageSize
	rtr.Type = rt
	rtr.UserID = user.GetUserId(c)
	rtr.Query = c.Query("query")

	res, err := rs.GetReportPagesByTickets(&rtr, list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(res, int(count), rtr.PageSize, rtr.PageIndex, "查询成功")
}

// GetReportListByIDs
// @Summary 使用ids获取报告列表
// @Description 使用ids获取报告列表
// @Tags 用户-报告
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/reports/ids [get]
// @Param ids query string true "ids"
// @Security Bearer
func (e Report) GetReportListByIDs(c *gin.Context) {
	rs := service.Report{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&rs.Service).
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

	reports, err := rs.GetReportByIDs(ids)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(reports, "查询成功")
}
