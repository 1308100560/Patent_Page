package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"

	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDeptRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerDeptNoAuthRouter)
}

// 需认证的路由代码
func registerDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	api := apis.Dept{}
	r := v1.Group("/dept").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.PUT("/:id", api.UpdateDept)
		r.POST("", api.CreateDept)
		r.DELETE("/:id", api.RemoveDept)
	}
}

func registerDeptNoAuthRouter(v1 *gin.RouterGroup) {
	api := apis.Dept{}
	v1auth := v1.Group("/dept")
	{
		v1auth.GET("", api.GetDeptPages)
	}
}
