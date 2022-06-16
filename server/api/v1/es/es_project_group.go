package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/es"
	esReq "github.com/flipped-aurora/gin-vue-admin/server/model/es/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProjectGroupApi struct {
}

var pgService = service.ServiceGroupApp.EsServiceGroup.ProjectGroupService

// CreateProjectGroup 创建ProjectGroup
// @Tags ProjectGroup
// @Summary 创建ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.ProjectGroup true "创建ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pg/createProjectGroup [post]
func (pgApi *ProjectGroupApi) CreateProjectGroup(c *gin.Context) {
	var pg es.ProjectGroup
	_ = c.ShouldBindJSON(&pg)
	if err := pgService.CreateProjectGroup(pg); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProjectGroup 删除ProjectGroup
// @Tags ProjectGroup
// @Summary 删除ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.ProjectGroup true "删除ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pg/deleteProjectGroup [delete]
func (pgApi *ProjectGroupApi) DeleteProjectGroup(c *gin.Context) {
	var pg es.ProjectGroup
	_ = c.ShouldBindJSON(&pg)
	if err := pgService.DeleteProjectGroup(pg); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProjectGroupByIds 批量删除ProjectGroup
// @Tags ProjectGroup
// @Summary 批量删除ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pg/deleteProjectGroupByIds [delete]
func (pgApi *ProjectGroupApi) DeleteProjectGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := pgService.DeleteProjectGroupByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProjectGroup 更新ProjectGroup
// @Tags ProjectGroup
// @Summary 更新ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body es.ProjectGroup true "更新ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pg/updateProjectGroup [put]
func (pgApi *ProjectGroupApi) UpdateProjectGroup(c *gin.Context) {
	var pg es.ProjectGroup
	_ = c.ShouldBindJSON(&pg)
	if err := pgService.UpdateProjectGroup(pg); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProjectGroup 用id查询ProjectGroup
// @Tags ProjectGroup
// @Summary 用id查询ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query es.ProjectGroup true "用id查询ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pg/findProjectGroup [get]
func (pgApi *ProjectGroupApi) FindProjectGroup(c *gin.Context) {
	var pg es.ProjectGroup
	_ = c.ShouldBindQuery(&pg)
	if repg, err := pgService.GetProjectGroup(pg.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repg": repg}, c)
	}
}

// GetProjectGroupList 分页获取ProjectGroup列表
// @Tags ProjectGroup
// @Summary 分页获取ProjectGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query esReq.ProjectGroupSearch true "分页获取ProjectGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pg/getProjectGroupList [get]
func (pgApi *ProjectGroupApi) GetProjectGroupList(c *gin.Context) {
	var pageInfo esReq.ProjectGroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := pgService.GetProjectGroupInfoList(pageInfo); err != nil {
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

// GetProjectGroupList 分页获取ProjectGroup列表树形结构
// @Tags ProjectGroup
// @Summary 分页获取ProjectGroup列表树形结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query esReq.ProjectGroupSearch true "分页获取ProjectGroup列表树形结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pg/getProjectGroupTreeList [get]

func (pgApi *ProjectGroupApi) GetProjectGroupTreeList(c *gin.Context) {
	var pageInfo esReq.ProjectGroupSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := pgService.GetProjectGroupTreeList(pageInfo); err != nil {
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
