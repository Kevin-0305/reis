import service from '@/utils/request'

// @Tags JenkinsBuildHistory
// @Summary 创建JenkinsBuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsBuildHistory true "创建JenkinsBuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jbh/createJenkinsBuildHistory [post]
export const createJenkinsBuildHistory = (data) => {
  return service({
    url: '/jbh/createJenkinsBuildHistory',
    method: 'post',
    data
  })
}

// @Tags JenkinsBuildHistory
// @Summary 删除JenkinsBuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsBuildHistory true "删除JenkinsBuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jbh/deleteJenkinsBuildHistory [delete]
export const deleteJenkinsBuildHistory = (data) => {
  return service({
    url: '/jbh/deleteJenkinsBuildHistory',
    method: 'delete',
    data
  })
}

// @Tags JenkinsBuildHistory
// @Summary 删除JenkinsBuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除JenkinsBuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jbh/deleteJenkinsBuildHistory [delete]
export const deleteJenkinsBuildHistoryByIds = (data) => {
  return service({
    url: '/jbh/deleteJenkinsBuildHistoryByIds',
    method: 'delete',
    data
  })
}

// @Tags JenkinsBuildHistory
// @Summary 更新JenkinsBuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JenkinsBuildHistory true "更新JenkinsBuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jbh/updateJenkinsBuildHistory [put]
export const updateJenkinsBuildHistory = (data) => {
  return service({
    url: '/jbh/updateJenkinsBuildHistory',
    method: 'put',
    data
  })
}

// @Tags JenkinsBuildHistory
// @Summary 用id查询JenkinsBuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.JenkinsBuildHistory true "用id查询JenkinsBuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jbh/findJenkinsBuildHistory [get]
export const findJenkinsBuildHistory = (params) => {
  return service({
    url: '/jbh/findJenkinsBuildHistory',
    method: 'get',
    params
  })
}

// @Tags JenkinsBuildHistory
// @Summary 分页获取JenkinsBuildHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取JenkinsBuildHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jbh/getJenkinsBuildHistoryList [get]
export const getJenkinsBuildHistoryList = (params) => {
  return service({
    url: '/jbh/getJenkinsBuildHistoryList',
    method: 'get',
    params
  })
}


//@Tags JenkinsBuildHistory
//@Summary 获取构建的日志
//@Security ApiKeyAuth
//@accept application/json
//@Produce application/json
//@Param data query model.JenkinsBuildHistory true "获取构建的日志"
//@Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
//@Router /jbh/getJenkinsBuildLogText [get]
export const getJenkinsBuildLogText = (params) => {
  return service({
    url: '/jbh/getJenkinsBuildLogText',
    method: 'get',
    params
  })
}
