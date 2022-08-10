package prom

import "github.com/gin-gonic/gin"

// metrics
// @Tags prom
// @Summary 获取指定服务的指定实例的指定监控项的指定时间段的数据
// @Description 获取指定服务的指定实例的指定监控项的指定时间段的数据
// @Accept application/json
// @Produce application/json
// @Param data body prom.PromHost true "获取指定服务的指定实例的指定监控项的指定时间段的数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prom/GetHostRealInfo [get]
func (promApi *PromHostApi) Metrics(c *gin.Context) {

}
