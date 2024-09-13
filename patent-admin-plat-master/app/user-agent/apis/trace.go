package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
	"strconv"
)

type Trace struct {
	api.Api
}

var traceAPI Trace

func init() {
	traceAPI = Trace{}
}

// SelectTraceLog
// @Summary 检索用户行为Tracing
// @Description 检索用户行为Tracing
// @Tags 用户行为Tracing
// @Param data body dto.TracePageReq true "用户数据"
// @Success 200 {object} []models.TraceLog
// @Router /api/v1/user-agent/tracing/logs [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param action query string true "action"
// @Param userId query string true "userId"
// @Security Bearer
func (e Trace) SelectTraceLog(c *gin.Context) {
	s := service.Tracer{}

	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.TracePageReq{}
	req.Action = c.Query("action")
	req.PageIndex, _ = strconv.Atoi(c.Query("pageIndex"))
	req.PageSize, _ = strconv.Atoi(c.Query("pageSize"))

	req.UserID, _ = strconv.Atoi(c.Query("userId"))

	list := make([]models.TraceLog, 0)
	var count int64
	err = s.SelectTraceLog(&req, &list, &count)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(list, int(count), req.PageIndex, req.PageSize, "查询成功")
}

type Tracer struct {
	c *gin.Context
}

func (t *Tracer) SearchTracing(query string) {
	traceReq := &dto.TraceReq{
		UserID:  user.GetUserId(t.c),
		Action:  "Search",
		Desc:    fmt.Sprintf("查询操作，表达式：%s", query),
		Request: fmt.Sprintf("请求URI：%s", t.c.Request.RequestURI),
	}

	t.trace(traceReq)
}

func (t *Tracer) GraphTracing(graph string) {
	traceReq := &dto.TraceReq{
		UserID:  user.GetUserId(t.c),
		Action:  "Graph",
		Desc:    fmt.Sprintf("数据分析，图表：%s", graph),
		Request: fmt.Sprintf("请求URI：%s", t.c.Request.RequestURI),
	}

	t.trace(traceReq)
}

func (t *Tracer) ReportTracing(report string) {
	traceReq := &dto.TraceReq{
		UserID:  user.GetUserId(t.c),
		Action:  "Report",
		Desc:    fmt.Sprintf("生成报告，报告：%s", report),
		Request: fmt.Sprintf("请求URI：%s", t.c.Request.RequestURI),
	}

	t.trace(traceReq)
}

func (t *Tracer) trace(req *dto.TraceReq) {
	s := service.Tracer{}
	err := traceAPI.MakeContext(t.c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		traceAPI.Logger.Error(err)
		traceAPI.Error(500, err, err.Error())
		return
	}

	err = s.Trace(req)
	if err != nil {
		traceAPI.Logger.Error(err)
		traceAPI.Error(500, err, err.Error())
		return
	}
}

func InternalTrace(c *gin.Context) *Tracer {
	return &Tracer{
		c: c,
	}
}
