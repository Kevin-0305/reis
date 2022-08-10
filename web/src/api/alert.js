import service from '@/utils/request'

//@Summary 获取激活告警方式列表
//@Produce  application/json
//@Param data body {}
//@Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"

export const getAlertTypes = (params) => {
  return service({
    url: '/alert/getAlertTypes',
    method: 'get',
    params: params
  })
}

//@Summary 告警方式测试
//@Produce  application/json
//@Param data body {alertType:"string"}
//@Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
export const alertTest = (params) => {
    return service({
        url: '/alert/alertTest',
        method: 'get',
        params: params
    })
}

//@Summary 获取prometheus告警记录
//@Produce  application/json
//@Param data body {}
//@Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
export const getPrometheusAlertHistory = (data) => {
    return service({
        url: '/alert/getPrometheusAlertHistory',
        method: 'post',
        data
    })
}
