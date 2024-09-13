package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/my_config"
	"strconv"
)

type Ticket struct {
	api.Api
}

// GetTicketPages
// @Summary 获取工单列表
// @Description 获取工单列表
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/tickets [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param type query string true "type"
// @Param query query string true "query"
// @Param status query string true "status"
// @Param reportType query string false "reportType"
// @Security Bearer
func (e Ticket) GetTicketPages(c *gin.Context) {
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

	t := c.Query("type")
	status := c.Query("status")
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dtos.TicketPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Type = t
	req.Status = status
	req.UserID = user.GetUserId(c)
	req.Query = c.Query("query")

	list := make([]model.Ticket, 0)
	var count int64

	switch t {
	case dtos.TicketTypeReport:
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
		err = s.GetReportTicketPages(&req, rt, &list, &count)
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
	default:
		if err = s.GetTicketPages(&req, &list, &count); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
	}

	e.PageOK(list, int(count), req.PageSize, req.PageIndex, "查询成功")
}

// GetTicketContactInfo
// @Summary 获取工单列表
// @Description 获取工单列表
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/tickets/contactInfo [get]
// @Security Bearer
func (e Ticket) GetTicketContactInfo(c *gin.Context) {
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

	e.OK(my_config.CurrentPatentConfig.NoveltyReportConfig.ContactTel, "查询成功")
}
