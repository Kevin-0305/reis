import service from '@/utils/request'

// @Tags EsCluter
// @Summary 创建EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EsCluter true "创建EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /esc/createEsCluter [post]
export const createEsCluter = (data) => {
  return service({
    url: '/esc/createEsCluter',
    method: 'post',
    data
  })
}

// @Tags EsCluter
// @Summary 删除EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EsCluter true "删除EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /esc/deleteEsCluter [delete]
export const deleteEsCluter = (data) => {
  return service({
    url: '/esc/deleteEsCluter',
    method: 'delete',
    data
  })
}

// @Tags EsCluter
// @Summary 删除EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /esc/deleteEsCluter [delete]
export const deleteEsCluterByIds = (data) => {
  return service({
    url: '/esc/deleteEsCluterByIds',
    method: 'delete',
    data
  })
}

// @Tags EsCluter
// @Summary 更新EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EsCluter true "更新EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /esc/updateEsCluter [put]
export const updateEsCluter = (data) => {
  return service({
    url: '/esc/updateEsCluter',
    method: 'put',
    data
  })
}

// @Tags EsCluter
// @Summary 用id查询EsCluter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EsCluter true "用id查询EsCluter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /esc/findEsCluter [get]
export const findEsCluter = (params) => {
  return service({
    url: '/esc/findEsCluter',
    method: 'get',
    params
  })
}

// @Tags EsCluter
// @Summary 分页获取EsCluter列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EsCluter列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /esc/getEsCluterList [get]
export const getEsCluterList = (params) => {
  return service({
    url: '/esc/getEsCluterList',
    method: 'get',
    params
  })
}


// @Tags EsCluter
// @Summary 检测EsCluter状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EsCluter true "检测EsCluter状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /esc/createEsCluter [post]
export const checkEsCluter = (data) => {
  return service({
    url: '/esc/checkEsCluter',
    method: 'post',
    data
  })
}



// @Tags EsCluter
// @Summary 设置集群的项目分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {"cluterId":0,"groupId":[]} true "设置集群的项目分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /esc/setEsCluterGroup [post]
export const setEsCluterGroup = (data) => {
  return service({
    url: '/esc/setEsCluterGroup',
    method: 'put',
    data
  })
}