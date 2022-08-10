package alert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alert/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//查询告警历史记录
// @Tags alert
// @Summary 查询告警历史记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /alert/GetAlertHistory [post]
func (alertApi *AlertApi) GetPrometheusAlertHistory(c *gin.Context) {
	var alertHistoryRequest request.GetPrometheusAlertHistoryRequest
	_ = c.ShouldBindJSON(&alertHistoryRequest)
	res, count, err := AlertService.GetPrometheusAlertHistory(alertHistoryRequest)
	if err != nil {
		global.GVA_LOG.Error("err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{List: res, Total: count, Page: alertHistoryRequest.PageInfo.Page, PageSize: alertHistoryRequest.PageInfo.PageSize}, "获取成功", c)
	}
}
