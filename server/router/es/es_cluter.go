package es

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EsCluterRouter struct {
}

// InitEsCluterRouter 初始化 EsCluter 路由信息
func (s *EsCluterRouter) InitEsCluterRouter(Router *gin.RouterGroup) {
	escRouter := Router.Group("esc").Use(middleware.OperationRecord())
	escRouterWithoutRecord := Router.Group("esc")
	var escApi = v1.ApiGroupApp.EsApiGroup.EsCluterApi
	{
		escRouter.POST("createEsCluter", escApi.CreateEsCluter)             // 新建EsCluter
		escRouter.DELETE("deleteEsCluter", escApi.DeleteEsCluter)           // 删除EsCluter
		escRouter.DELETE("deleteEsCluterByIds", escApi.DeleteEsCluterByIds) // 批量删除EsCluter
		escRouter.PUT("updateEsCluter", escApi.UpdateEsCluter)              // 更新EsCluter
		escRouter.POST("checkEsCluter", escApi.CheckEsCluterStatus)         // 检测集群状态
		escRouter.PUT("setEsCluterGroup", escApi.SetEsCluterGroup)          // 更新EsCluterGroup

	}
	{
		escRouterWithoutRecord.GET("findEsCluter", escApi.FindEsCluter)       // 根据ID获取EsCluter
		escRouterWithoutRecord.GET("getEsCluterList", escApi.GetEsCluterList) // 获取EsCluter列表
	}
}
