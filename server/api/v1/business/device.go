package business

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeviceApi struct{}

// CreateDevice
// @Tags      Device
// @Summary   创建设备配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      business.Device            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /device/deviceConfig [post]
func (e *DeviceApi) CreateDeviceConfig(c *gin.Context) {
	var dc business.DeviceConfig
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dc.ModifyTime = time.Now().Unix()
	dc.CreateTime = time.Now().Unix()
	err = deviceService.CreateDeviceConfig(dc)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateDeviceConfig
// @Tags      Device
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /device/deviceConfig [put]
func (e *DeviceApi) UpdateDeviceConfig(c *gin.Context) {
	var dc business.DeviceConfig
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = deviceService.UpdateDeviceConfig(&dc)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetDeviceConfigList
// @Tags      Device
// @Summary   分页获取设备配置列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /device/deviceConfigList [get]
func (e *DeviceApi) GetDeviceConfigList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deviceList, total, err := deviceService.GetDeviceConfigList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     deviceList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
