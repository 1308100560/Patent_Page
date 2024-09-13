package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"strconv"
)

type Dept struct {
	api.Api
}

// GetDeptPages
// @Summary 列表部门分页
// @Description 列表部门信息
// @Tags 部门
// @Router /apis/v1/dept [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Security Bearer
func (e Dept) GetDeptPages(c *gin.Context) {
	s := service.SysDept{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.DeptPagesReq{}
	req.PageIndex, _ = strconv.Atoi(c.Query("pageIndex"))
	req.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	req.Query = c.Query("query")

	list := make([]models.SysDept, 0)
	var count int64
	err = s.GetDeptPages(req, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.PageIndex, req.PageSize, "查询成功")
}

// CreateDept
// @Summary 管理员创建团队
// @Description  管理员创建团队
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body dto.DeptReq true "部门详情"
// @Router /api/v1/dept [post]
// @Security Bearer
func (e Dept) CreateDept(c *gin.Context) {

	s := service.SysDept{}
	reqIn := dto.DeptReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&reqIn, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Insert(&reqIn)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(reqIn.DeptName, "创建成功")
}

// UpdateDept
// @Summary 管理员更新团队
// @Description  管理员更新团队
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body dto.DeptReq true "部门详情"
// @Router /api/v1/dept/{id} [put]
// @Security Bearer
func (e Dept) UpdateDept(c *gin.Context) {
	s := service.SysDept{}
	reqIn := dto.DeptReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&reqIn, binding.JSON).
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
	reqIn.DeptId = id

	err = s.Update(&reqIn)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(reqIn.DeptName, "创建成功")
}

// RemoveDept
// @Summary 管理员删除团队
// @Description  管理员删除团队
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Router /api/v1/dept/{id} [delete]
// @Security Bearer
func (e Dept) RemoveDept(c *gin.Context) {
	s := service.SysDept{}
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

	err = s.Remove(id)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "删除成功")
}
