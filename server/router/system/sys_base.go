package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	webHookRouter := Router.Group("webHook")
	webHookRobotApi := v1.ApiGroupApp.WebHookApiGroup.WebHookRobotApi
	{
		webHookRouter.POST("FeiShuApp", webHookRobotApi.FeiShuApp)
	}

	return baseRouter
}
