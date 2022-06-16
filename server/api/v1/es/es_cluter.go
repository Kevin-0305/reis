package es

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/es"
	esReq "github.com/flipped-aurora/gin-vue-admin/server/model/es/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EsCluterApi struct {
}

var escService = service.ServiceGroupApp.EsServiceGroup.EsCluterService

// CreateEsCluter 创建EsCluter
// @Tags EsCluter
// @Summary 创建EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.EsCluter true "创建EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /esc/createEsCluter [post]
func (escApi *EsCluterApi) CreateEsCluter(c *gin.Context) {
	var esc es.EsCluter
	_ = c.ShouldBindJSON(&esc)
	if err := escService.CreateEsCluter(esc); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEsCluter 删除EsCluter
// @Tags EsCluter
// @Summary 删除EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.EsCluter true "删除EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /esc/deleteEsCluter [delete]
func (escApi *EsCluterApi) DeleteEsCluter(c *gin.Context) {
	var esc es.EsCluter
	_ = c.ShouldBindJSON(&esc)
	if err := escService.DeleteEsCluter(esc); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEsCluterByIds 批量删除EsCluter
// @Tags EsCluter
// @Summary 批量删除EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /esc/deleteEsCluterByIds [delete]
func (escApi *EsCluterApi) DeleteEsCluterByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := escService.DeleteEsCluterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEsCluter 更新EsCluter
// @Tags EsCluter
// @Summary 更新EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.EsCluter true "更新EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /esc/updateEsCluter [put]
func (escApi *EsCluterApi) UpdateEsCluter(c *gin.Context) {
	var esc es.EsCluter
	_ = c.ShouldBindJSON(&esc)
	if err := escService.UpdateEsCluter(esc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEsCluter 用id查询EsCluter
// @Tags EsCluter
// @Summary 用id查询EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query es.EsCluter true "用id查询EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /esc/findEsCluter [get]
func (escApi *EsCluterApi) FindEsCluter(c *gin.Context) {
	var esc es.EsCluter
	_ = c.ShouldBindQuery(&esc)
	if reesc, err := escService.GetEsCluter(esc.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reesc": reesc}, c)
	}
}

// GetEsCluterList 分页获取EsCluter列表
// @Tags EsCluter
// @Summary 分页获取EsCluter列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query esReq.EsCluterSearch true "分页获取EsCluter列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /esc/getEsCluterList [get]
func (escApi *EsCluterApi) GetEsCluterList(c *gin.Context) {
	var pageInfo esReq.EsCluterSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := escService.GetEsCluterInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CheckEsCluterStatus 检测EsCluter状态
// @Tags EsCluter
// @Summary 检测EsCluter状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.EsCluter true "检测EsCluter状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /esc/checkEsCluter [post]
func (escApi *EsCluterApi) CheckEsCluterStatus(c *gin.Context) {
	var esc es.EsCluter
	err := c.ShouldBindJSON(&esc)
	if err != nil {
		fmt.Println(err)
	}
	if status, err := escService.CheckEsCluterStatus(esc); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"status": status}, c)
	}

}

// SaveEsCluterGroup 修改EsCluter组
// @Tags EsCluter
// @Summary 修改EsCluter组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.EsCluterGroup true "修改EsCluter组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /esc/saveEsCluterGroup [post]
func (escApi *EsCluterApi) SetEsCluterGroup(c *gin.Context) {
	// cluterGroup := make(map[string]interface{})
	cluterGroup := struct {
		CluterId uint   `json:"cluterId"`
		GroupIds []uint `json:"groupIds"`
	}{}
	_ = c.ShouldBindJSON(&cluterGroup)
	if err := escService.RefreshCluterGroup(cluterGroup.CluterId, cluterGroup.GroupIds); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
