package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
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
// @Router /api/v1/admin-agent/trading-records [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Param ids query string true "ids"
// @Security Bearer
func (e TradingRecord) GetTradingRecordPages(c *gin.Context) {
	s := service.TradingRecord{}
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
	if len(idsStr) != 0 {
		if err = json.Unmarshal([]byte(idsStr), &ids); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
	}

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dto.TradingRecordPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Query = c.Query("query")
	req.PatentIDS = ids

	list := make([]models.TradingRecord, 0)
	var count int64
	if err = s.GetPages(&req, &list, &count); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(list, int(count), req.PageSize, req.PageIndex, "查询成功")
}

// Insert
// @Summary 插入交易记录
// @Description 插入交易记录
// @Tags 交易记录
// @Accept  application/json
// @Product application/json
// @Param data body dto.TradingRecordReq{} true "专利包数据"
// @Router /api/v1/admin-agent/trading-records [post]
// @Security Bearer
func (e TradingRecord) Insert(c *gin.Context) {
	s := service.TradingRecord{}
	req := dto.TradingRecordReq{}
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
	data, err := s.Create(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(data, "创建成功")
}

// Update
// @Summary 更新交易记录
// @Description 更新交易记录
// @Tags 交易记录
// @Accept  application/json
// @Product application/json
// @Param data body dto.TradingRecordReq{} true "专利包数据"
// @Router /api/v1/admin-agent/trading-records/{id} [put]
// @Security Bearer
func (e TradingRecord) Update(c *gin.Context) {
	s := service.TradingRecord{}
	req := dto.TradingRecordReq{}
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.ID = id

	// 设置创建人
	data, err := s.Update(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(data, "更新成功")
}

// Remove
// @Summary 删除交易记录
// @Description 删除交易记录
// @Tags 交易记录
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/trading-records/{id} [delete]
// @Security Bearer
func (e TradingRecord) Remove(c *gin.Context) {
	s := service.TradingRecord{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置创建人
	err = s.Delete(id)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "删除成功")
}
