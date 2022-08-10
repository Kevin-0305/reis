package prom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var promService = service.ServiceGroupApp.PromServiceGroup.PromHostService

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type PromHostApi struct {
}

// var escService = service.ServiceGroupApp.EsServiceGroup.EsCluterService

// FindEsCluter 用id查询EsCluter
// @Tags EsCluter
// @Summary 用id查询EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /esc/findEsCluter [get]
func (promApi *PromHostApi) GetHostRealInfo(c *gin.Context) {
	instance := c.Query("instance")
	fmt.Println(instance)
	//value, err := promService.GetHostRealInfo("test", "192.168.2.20:9187")
	// if err != nil {
	// 	global.GVA_LOG.Error("err", zap.Error(err))
	// 	response.FailWithMessage(err.Error(), c)
	// } else {
	// 	response.OkWithDetailed(value, "获取成功", c)
	// }
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade", err.Error())
		response.FailWithMessage(err.Error(), c)
	}
	go readMessage(ws)
	go sendMessage(ws, instance, "test")

}

func readMessage(ws *websocket.Conn) {
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			ws.Close()
			global.GVA_LOG.Error("err", zap.Error(err))
			break
		}
		fmt.Println("recv:", string(message))
	}
}

func sendMessage(ws *websocket.Conn, instance string, service string) {
	for {
		time.Sleep(time.Second * 5)
		value, err := promService.GetHostRealInfo(service, instance)
		if err != nil {
			global.GVA_LOG.Error("err", zap.Error(err))
		} else {
			j, err := json.Marshal(value)
			if err != nil {
				global.GVA_LOG.Error("err", zap.Error(err))
			}
			err = ws.WriteMessage(websocket.TextMessage, []byte(j))
			if err != nil {
				global.GVA_LOG.Error("err", zap.Error(err))
				break
			}
			fmt.Println("send message")
		}

	}
}

//GetActiveHostList 获取一周内活跃过的主机的列表
//@Tags Prom
//@Summary 获取一周内活跃主机的列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query prom.PromHost true "获取一周内活跃主机的列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prom/GetActiveHostLis [get]
func (promApi *PromHostApi) GetActiveHostList(c *gin.Context) {
	serviceName := c.Query("serviceName")
	list, err := promService.GetActiveHostList(serviceName)
	if err != nil {
		global.GVA_LOG.Error("err", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

//GetPromServiceList 获取所有的prometheus服务
//@Tags Prom
//@Summary 获取所有的prometheus服务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query es.EsCluter true ""
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prom/GetPromServiceList [get]

func (promApi *PromHostApi) GetPromServiceList(c *gin.Context) {
	list := []string{}
	for k := range global.GVA_PromAPIs {
		list = append(list, k)
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}
