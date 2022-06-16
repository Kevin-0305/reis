package es

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectGroupRouter struct {
}

// InitProjectGroupRouter 初始化 ProjectGroup 路由信息
func (s *ProjectGroupRouter) InitProjectGroupRouter(Router *gin.RouterGroup) {
	pgRouter := Router.Group("pg").Use(middleware.OperationRecord())
	pgRouterWithoutRecord := Router.Group("pg")
	var pgApi = v1.ApiGroupApp.EsApiGroup.ProjectGroupApi
	{
		pgRouter.POST("createProjectGroup", pgApi.CreateProjectGroup)             // 新建ProjectGroup
		pgRouter.DELETE("deleteProjectGroup", pgApi.DeleteProjectGroup)           // 删除ProjectGroup
		pgRouter.DELETE("deleteProjectGroupByIds", pgApi.DeleteProjectGroupByIds) // 批量删除ProjectGroup
		pgRouter.PUT("updateProjectGroup", pgApi.UpdateProjectGroup)              // 更新ProjectGroup
	}
	{
		pgRouterWithoutRecord.GET("findProjectGroup", pgApi.FindProjectGroup)               // 根据ID获取ProjectGroup
		pgRouterWithoutRecord.GET("getProjectGroupList", pgApi.GetProjectGroupList)         // 获取ProjectGroup列表
		pgRouterWithoutRecord.GET("getProjectGroupTreeList", pgApi.GetProjectGroupTreeList) // 获取ProjectGroup树形结构列表
	}
}
