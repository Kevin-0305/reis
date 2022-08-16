import service from '@/utils/request'

// @Tags JenkinsJob
// @Summary 创建JenkinsJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsJob true "创建JenkinsJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jsmj/createJenkinsJob [post]
export const createJenkinsJob = (data) => {
  return service({
    url: '/jsmj/createJenkinsJob',
    method: 'post',
    data
  })
}

// @Tags JenkinsJob
// @Summary 删除JenkinsJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsJob true "删除JenkinsJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jsmj/deleteJenkinsJob [delete]
export const deleteJenkinsJob = (data) => {
  return service({
    url: '/jsmj/deleteJenkinsJob',
    method: 'delete',
    data
  })
}

// @Tags JenkinsJob
// @Summary 删除JenkinsJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除JenkinsJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jsmj/deleteJenkinsJob [delete]
export const deleteJenkinsJobByIds = (data) => {
  return service({
    url: '/jsmj/deleteJenkinsJobByIds',
    method: 'delete',
    data
  })
}

// @Tags JenkinsJob
// @Summary 更新JenkinsJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsJob true "更新JenkinsJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jsmj/updateJenkinsJob [put]
export const updateJenkinsJob = (data) => {
  return service({
    url: '/jsmj/updateJenkinsJob',
    method: 'put',
    data
  })
}

// @Tags JenkinsJob
// @Summary 用id查询JenkinsJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.JenkinsJob true "用id查询JenkinsJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jsmj/findJenkinsJob [get]
export const findJenkinsJob = (params) => {
  return service({
    url: '/jsmj/findJenkinsJob',
    method: 'get',
    params
  })
}

// @Tags JenkinsJob
// @Summary 分页获取JenkinsJob列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取JenkinsJob列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jsmj/getJenkinsJobList [get]
export const getJenkinsJobList = (params) => {
  return service({
    url: '/jsmj/getJenkinsJobList',
    method: 'get',
    params
  })
}
