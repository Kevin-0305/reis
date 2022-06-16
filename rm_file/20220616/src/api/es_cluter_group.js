import service from '@/utils/request'

// @Tags EsCluterGroup
// @Summary 创建EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EsCluterGroup true "创建EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ecg/createEsCluterGroup [post]
export const createEsCluterGroup = (data) => {
  return service({
    url: '/ecg/createEsCluterGroup',
    method: 'post',
    data
  })
}

// @Tags EsCluterGroup
// @Summary 删除EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EsCluterGroup true "删除EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ecg/deleteEsCluterGroup [delete]
export const deleteEsCluterGroup = (data) => {
  return service({
    url: '/ecg/deleteEsCluterGroup',
    method: 'delete',
    data
  })
}

// @Tags EsCluterGroup
// @Summary 删除EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ecg/deleteEsCluterGroup [delete]
export const deleteEsCluterGroupByIds = (data) => {
  return service({
    url: '/ecg/deleteEsCluterGroupByIds',
    method: 'delete',
    data
  })
}

// @Tags EsCluterGroup
// @Summary 更新EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EsCluterGroup true "更新EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ecg/updateEsCluterGroup [put]
export const updateEsCluterGroup = (data) => {
  return service({
    url: '/ecg/updateEsCluterGroup',
    method: 'put',
    data
  })
}

// @Tags EsCluterGroup
// @Summary 用id查询EsCluterGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EsCluterGroup true "用id查询EsCluterGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ecg/findEsCluterGroup [get]
export const findEsCluterGroup = (params) => {
  return service({
    url: '/ecg/findEsCluterGroup',
    method: 'get',
    params
  })
}

// @Tags EsCluterGroup
// @Summary 分页获取EsCluterGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EsCluterGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ecg/getEsCluterGroupList [get]
export const getEsCluterGroupList = (params) => {
  return service({
    url: '/ecg/getEsCluterGroupList',
    method: 'get',
    params
  })
}
