// import request from "@/utils/request";
import service from '@/utils/request'

// export function getDeviceConfigList() {
//   return request({
//     url: "/device/deviceConfigList",
//     method: "get"
//   });
// }

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
export const createDeviceConfig = (data) => {
  return service({
    url: '/device/deviceConfig',
    method: 'post',
    data
  })
}

// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [put]
export const updateDeviceConfig = (data) => {
  return service({
    url: '/device/deviceConfig',
    method: 'put',
    data
  })
}

// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [delete]
export const deleteDeviceConfig = (data) => {
  return service({
    url: '/device/deviceConfig',
    method: 'delete',
    data
  })
}

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
export const getDeviceConfig = (params) => {
  return service({
    url: '/device/deviceConfig',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
export const getDeviceConfigList = (params) => {
  return service({
    url: '/device/deviceConfigList',
    method: 'get',
    params
  })
}

// @Tags SysApi
// @Summary 播放
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [put]
export const playVideo = (data) => {
  return service({
    url: '/device/play',
    method: 'put',
    data
  })
}
