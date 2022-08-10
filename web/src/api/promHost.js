import service from '@/utils/request'
// @Summary 获取主机实时信息
// @Produce  application/json
// @Param data body {hostname:"string"}
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
export const getHostRealInfo = (params) => {
  return service({
    url: '/prom/GetHostRealInfo',
    method: 'get',
    params: params
  })
}

// @Summary 获取一周的内活跃主机列表
// @Produce  application/json
// @Param data body {hostname:"string"}
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"

export const getActiveHostList = (params) =>{
  return service({
    url: '/prom/GetActiveHostList',
    method: 'get',
    params: params
  })
}


// @Summary 获取prometheus服务器列表
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
export const getPromServiceList = (params) =>{
  return service({
    url: '/prom/GetPromServiceList',
    method: 'get',
    params: params
  })
}