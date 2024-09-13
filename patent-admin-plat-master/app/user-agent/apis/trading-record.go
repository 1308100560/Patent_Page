package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	userService "go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"strconv"
)

type TradingRecord struct {
	api.Api
}

// GetTradingRecordPages
// @Summary 获取交易记录列表
// @Description 获取交易记录列表
// @Tags 交易记录
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/trading-records [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Security Bearer
func (e TradingRecord) GetTradingRecordPages(c *gin.Context) {
	s := userService.TradingRecord{}
	ups := userService.UserPatent{}
	ps := userService.Patent{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		MakeService(&ups.Service).
		MakeService(&ps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)
	claimList := make([]models.UserPatent, 0)
	if err = ups.GetClaimLists(userID, &claimList); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// get collaborators
	claimIds := make([]int, len(claimList))
	for _, cp := range claimList {
		claimIds = append(claimIds, cp.PatentId)
	}

	var _c int64
	claimPatentsMap := make(map[int]*models.Patent)
	claimPatents, err := ps.GetPatentsByIds(claimIds, &_c)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	for _, p := range claimPatents {
		claimPatentsMap[p.PatentId] = &p
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dto.TradingRecordPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Query = c.Query("query")
	req.PatentIDS = claimIds

	list := make([]models.TradingRecord, 0)
	var count int64
	if err = s.GetPages(&req, &list, &count); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	for _, record := range list {
		record.Patent = claimPatentsMap[record.PatentID]
	}

	e.PageOK(list, int(count), req.PageSize, req.PageIndex, "查询成功")
}
