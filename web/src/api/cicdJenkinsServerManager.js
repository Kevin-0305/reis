import service from '@/utils/request'

// @Tags JenkinsServerManager
// @Summary 创建JenkinsServerManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsServerManager true "创建JenkinsServerManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jsm/createJenkinsServerManager [post]
export const createJenkinsServerManager = (data) => {
  return service({
    url: '/jsm/createJenkinsServerManager',
    method: 'post',
    data
  })
}

// @Tags JenkinsServerManager
// @Summary 删除JenkinsServerManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsServerManager true "删除JenkinsServerManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jsm/deleteJenkinsServerManager [delete]
export const deleteJenkinsServerManager = (data) => {
  return service({
    url: '/jsm/deleteJenkinsServerManager',
    method: 'delete',
    data
  })
}

// @Tags JenkinsServerManager
// @Summary 删除JenkinsServerManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除JenkinsServerManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jsm/deleteJenkinsServerManager [delete]
export const deleteJenkinsServerManagerByIds = (data) => {
  return service({
    url: '/jsm/deleteJenkinsServerManagerByIds',
    method: 'delete',
    data
  })
}

// @Tags JenkinsServerManager
// @Summary 更新JenkinsServerManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsServerManager true "更新JenkinsServerManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jsm/updateJenkinsServerManager [put]
export const updateJenkinsServerManager = (data) => {
  return service({
    url: '/jsm/updateJenkinsServerManager',
    method: 'put',
    data
  })
}

// @Tags JenkinsServerManager
// @Summary 用id查询JenkinsServerManager
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.JenkinsServerManager true "用id查询JenkinsServerManager"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jsm/findJenkinsServerManager [get]
export const findJenkinsServerManager = (params) => {
  return service({
    url: '/jsm/findJenkinsServerManager',
    method: 'get',
    params
  })
}

// @Tags JenkinsServerManager
// @Summary 分页获取JenkinsServerManager列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取JenkinsServerManager列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jsm/getJenkinsServerManagerList [get]
export const getJenkinsServerManagerList = (params) => {
  return service({
    url: '/jsm/getJenkinsServerManagerList',
    method: 'get',
    params
  })
}
