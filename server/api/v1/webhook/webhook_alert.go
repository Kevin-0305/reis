package webhook

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	webhookReq "github.com/flipped-aurora/gin-vue-admin/server/model/webhook/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var prometheusAlertService = service.ServiceGroupApp.WebHookServiceGroup.PrometheusAlertService

type WebHookAlertApi struct {
}

func (s *WebHookAlertApi) PrometheusAlert(c *gin.Context) {
	PrometheusAlertRequest := webhookReq.PrometheusAlertRequest{}
	PrometheusAlertRequest.AlertTime = time.Now()
	if err := c.ShouldBindJSON(&PrometheusAlertRequest); err != nil {
		global.GVA_LOG.Error("err", zap.Error(err))
		response.FailWithMessage("参数错误", c)
	}
	err := prometheusAlertService.PrometheusAlert(PrometheusAlertRequest)
	if err != nil {
		global.GVA_LOG.Error("err", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}
