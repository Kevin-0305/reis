package prom

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PromHostRouter struct {
}

// InitPromHostRouter 初始化 PromHost 路由信息
func (s *PromHostRouter) InitPromHostRouter(Router *gin.RouterGroup) {
	//promRouter := Router.Group("prom").Use(middleware.OperationRecord())
	promRouterWithoutRecord := Router.Group("prom")
	var promHostApi = v1.ApiGroupApp.PromApiGroup.PromHostApi
	{
	}
	{
		promRouterWithoutRecord.GET("GetActiveHostList", promHostApi.GetActiveHostList)   //获取活跃主机列表
		promRouterWithoutRecord.GET("GetHostRealInfo", promHostApi.GetHostRealInfo)       // 获取主机实时信息
		promRouterWithoutRecord.GET("GetPromServiceList", promHostApi.GetPromServiceList) //获取所有的prometheus服务列表
	}

}
