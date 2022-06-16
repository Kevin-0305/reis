import service from '@/utils/request'

// @Tags ProjectGroup
// @Summary 创建ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectGroup true "创建ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pg/createProjectGroup [post]
export const createProjectGroup = (data) => {
  return service({
    url: '/pg/createProjectGroup',
    method: 'post',
    data
  })
}

// @Tags ProjectGroup
// @Summary 删除ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectGroup true "删除ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pg/deleteProjectGroup [delete]
export const deleteProjectGroup = (data) => {
  return service({
    url: '/pg/deleteProjectGroup',
    method: 'delete',
    data
  })
}

// @Tags ProjectGroup
// @Summary 删除ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pg/deleteProjectGroup [delete]
export const deleteProjectGroupByIds = (data) => {
  return service({
    url: '/pg/deleteProjectGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags ProjectGroup
// @Summary 更新ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectGroup true "更新ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pg/updateProjectGroup [put]
export const updateProjectGroup = (data) => {
  return service({
    url: '/pg/updateProjectGroup',
    method: 'put',
    data
  })
}

// @Tags ProjectGroup
// @Summary 用id查询ProjectGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProjectGroup true "用id查询ProjectGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pg/findProjectGroup [get]
export const findProjectGroup = (params) => {
  return service({
    url: '/pg/findProjectGroup',
    method: 'get',
    params
  })
}

// @Tags ProjectGroup
// @Summary 分页获取ProjectGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ProjectGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pg/getProjectGroupList [get]
export const getProjectGroupList = (params) => {
  return service({
    url: '/pg/getProjectGroupList',
    method: 'get',
    params
  })
}


// @Tags ProjectGroup
// @Summary 获取ProjectGroup树形结构列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pg/getProjectGroupTreeList [get]
export const getProjectGroupTreeList = (params) => {
  return service({
    url: '/pg/getProjectGroupTreeList',
    method: 'get'
  })
}



