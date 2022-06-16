package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/es"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    esReq "github.com/flipped-aurora/gin-vue-admin/server/model/es/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type EsCluterGroupApi struct {
}

var ecgService = service.ServiceGroupApp.EsServiceGroup.EsCluterGroupService


// CreateEsCluterGroup 创建EsCluterGroup
// @Tags EsCluterGroup
// @Summary 创建EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.EsCluterGroup true "创建EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ecg/createEsCluterGroup [post]
func (ecgApi *EsCluterGroupApi) CreateEsCluterGroup(c *gin.Context) {
	var ecg es.EsCluterGroup
	_ = c.ShouldBindJSON(&ecg)
	if err := ecgService.CreateEsCluterGroup(ecg); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEsCluterGroup 删除EsCluterGroup
// @Tags EsCluterGroup
// @Summary 删除EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.EsCluterGroup true "删除EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ecg/deleteEsCluterGroup [delete]
func (ecgApi *EsCluterGroupApi) DeleteEsCluterGroup(c *gin.Context) {
	var ecg es.EsCluterGroup
	_ = c.ShouldBindJSON(&ecg)
	if err := ecgService.DeleteEsCluterGroup(ecg); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEsCluterGroupByIds 批量删除EsCluterGroup
// @Tags EsCluterGroup
// @Summary 批量删除EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ecg/deleteEsCluterGroupByIds [delete]
func (ecgApi *EsCluterGroupApi) DeleteEsCluterGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := ecgService.DeleteEsCluterGroupByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEsCluterGroup 更新EsCluterGroup
// @Tags EsCluterGroup
// @Summary 更新EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.EsCluterGroup true "更新EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ecg/updateEsCluterGroup [put]
func (ecgApi *EsCluterGroupApi) UpdateEsCluterGroup(c *gin.Context) {
	var ecg es.EsCluterGroup
	_ = c.ShouldBindJSON(&ecg)
	if err := ecgService.UpdateEsCluterGroup(ecg); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEsCluterGroup 用id查询EsCluterGroup
// @Tags EsCluterGroup
// @Summary 用id查询EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query es.EsCluterGroup true "用id查询EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ecg/findEsCluterGroup [get]
func (ecgApi *EsCluterGroupApi) FindEsCluterGroup(c *gin.Context) {
	var ecg es.EsCluterGroup
	_ = c.ShouldBindQuery(&ecg)
	if reecg, err := ecgService.GetEsCluterGroup(ecg.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reecg": reecg}, c)
	}
}

// GetEsCluterGroupList 分页获取EsCluterGroup列表
// @Tags EsCluterGroup
// @Summary 分页获取EsCluterGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query esReq.EsCluterGroupSearch true "分页获取EsCluterGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ecg/getEsCluterGroupList [get]
func (ecgApi *EsCluterGroupApi) GetEsCluterGroupList(c *gin.Context) {
	var pageInfo esReq.EsCluterGroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := ecgService.GetEsCluterGroupInfoList(pageInfo); err != nil {
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
