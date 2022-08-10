package webhook

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WebHookRouter struct {
}

// InitWebHookRouter initialize webhook router
func (s *WebHookRouter) InitWebHookRouter(Router *gin.RouterGroup) {
	//alertRouter := Router.Group("alert").Use(middleware.OperationRecord())
	webHookRouterWithoutRecord := Router.Group("webhook")
	var webHookAlertApi = v1.ApiGroupApp.WebHookApiGroup.WebHookAlertApi
	{
	}
	{
		webHookRouterWithoutRecord.POST("prometheusAlert", webHookAlertApi.PrometheusAlert)

	}

}
