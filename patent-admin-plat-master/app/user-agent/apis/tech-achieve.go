package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"strconv"
)

type TechAchieve struct {
	api.Api
}

// GetTechAchievePages
// @Summary 获取科技成果列表
// @Description 获取科技成果列表
// @Tags 科技成果
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/tech-achieves [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Security Bearer
func (e TechAchieve) GetTechAchievePages(c *gin.Context) {
	s := service.TechAchieve{}
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

	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dto.TechAchievePagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Query = c.Query("query")
	req.UserID = userID

	list := make([]models.TechAchieve, 0)
	var count int64
	if err = s.GetPages(&req, &list, &count); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(list, int(count), req.PageSize, req.PageIndex, "查询成功")
}

// Insert
// @Summary 插入科技成果
// @Description 插入科技成果
// @Tags 科技成果
// @Accept  application/json
// @Product application/json
// @Param data body dto.TechAchieveReq{} true "专利包数据"
// @Router /api/v1/user-agent/tech-achieves [post]
// @Security Bearer
func (e TechAchieve) Insert(c *gin.Context) {
	s := service.TechAchieve{}
	req := dto.TechAchieveReq{}
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

	userID := user.GetUserId(c)
	req.UserID = userID

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
// @Summary 更新科技成果
// @Description 更新科技成果
// @Tags 科技成果
// @Accept  application/json
// @Product application/json
// @Param data body dto.TechAchieveReq{} true "专利包数据"
// @Router /api/v1/user-agent/tech-achieves/{id} [put]
// @Security Bearer
func (e TechAchieve) Update(c *gin.Context) {
	s := service.TechAchieve{}
	req := dto.TechAchieveReq{}
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

	userID := user.GetUserId(c)
	req.UserID = userID

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
// @Summary 删除科技成果
// @Description 删除科技成果
// @Tags 科技成果
// @Accept  application/json
// @Product application/json
// @Router /api/v1/user-agent/tech-achieves/{id} [delete]
// @Security Bearer
func (e TechAchieve) Remove(c *gin.Context) {
	s := service.TechAchieve{}
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
