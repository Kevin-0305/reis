package alert

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var AlertService = service.ServiceGroupApp.AlertServiceGroup.AlertService

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type AlertApi struct {
}

// AlertTest 测试告警
// @Tags alert
// @Summary 测试告警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /alert/AlertTest [get]
func (alertApi *AlertApi) AlertTest(c *gin.Context) {
	alertType := c.Query("alertType")
	value, err := AlertService.AlertTest(alertType)
	if err != nil {
		global.GVA_LOG.Error("err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(value, "获取成功", c)
	}
}

// GetAlertTypes 获取告警类型
// @Tags alert
// @Summary 获取告警类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /alert/GetAlertTypes [get]
func (alertApi *AlertApi) GetAlertTypes(c *gin.Context) {
	list, err := AlertService.GetAlertTypes()
	if err != nil {
		global.GVA_LOG.Error("err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{List: list}, "获取成功", c)
	}
}
