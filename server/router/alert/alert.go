package alert

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AlertRouter struct {
}

// InitAlertRouter initialize alert router
func (s *AlertRouter) InitAlertRouter(Router *gin.RouterGroup) {
	//alertRouter := Router.Group("alert").Use(middleware.OperationRecord())
	alertRouterWithoutRecord := Router.Group("alert")
	var alertApi = v1.ApiGroupApp.AlertApiGroup.AlertApi
	{
	}
	{
		alertRouterWithoutRecord.GET("alertTest", alertApi.AlertTest)                                  //获取活跃主机列表
		alertRouterWithoutRecord.GET("getAlertTypes", alertApi.GetAlertTypes)                          //获取方式告警类型
		alertRouterWithoutRecord.POST("getPrometheusAlertHistory", alertApi.GetPrometheusAlertHistory) //获取告警历史记录

	}

}
