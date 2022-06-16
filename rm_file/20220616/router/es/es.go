package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EsCluterGroupRouter struct {
}

// InitEsCluterGroupRouter 初始化 EsCluterGroup 路由信息
func (s *EsCluterGroupRouter) InitEsCluterGroupRouter(Router *gin.RouterGroup) {
	ecgRouter := Router.Group("ecg").Use(middleware.OperationRecord())
	ecgRouterWithoutRecord := Router.Group("ecg")
	var ecgApi = v1.ApiGroupApp.EsApiGroup.EsCluterGroupApi
	{
		ecgRouter.POST("createEsCluterGroup", ecgApi.CreateEsCluterGroup)   // 新建EsCluterGroup
		ecgRouter.DELETE("deleteEsCluterGroup", ecgApi.DeleteEsCluterGroup) // 删除EsCluterGroup
		ecgRouter.DELETE("deleteEsCluterGroupByIds", ecgApi.DeleteEsCluterGroupByIds) // 批量删除EsCluterGroup
		ecgRouter.PUT("updateEsCluterGroup", ecgApi.UpdateEsCluterGroup)    // 更新EsCluterGroup
	}
	{
		ecgRouterWithoutRecord.GET("findEsCluterGroup", ecgApi.FindEsCluterGroup)        // 根据ID获取EsCluterGroup
		ecgRouterWithoutRecord.GET("getEsCluterGroupList", ecgApi.GetEsCluterGroupList)  // 获取EsCluterGroup列表
	}
}
