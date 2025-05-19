package business

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessRes "github.com/flipped-aurora/gin-vue-admin/server/model/business/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
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

// DeleteDeviceConfig
// @Tags      Device
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /device/deviceConfig [delete]
func (e *DeviceApi) DeleteDeviceConfig(c *gin.Context) {
	var dc business.DeviceConfig
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = deviceService.DeleteDeviceConfig(&dc)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetDeviceConfig
// @Tags      Device
// @Summary   分页获取设备配置列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /device/deviceConfigList [get]
func (e *DeviceApi) GetDeviceConfig(c *gin.Context) {
	var ei request.EntityId
	err := c.ShouldBindQuery(&ei)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deviceConfig, err := deviceService.GetDeviceConfig(ei)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(deviceConfig, "获取成功", c)
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
	// err = utils.Verify(pageInfo, utils.PageInfoVerify)
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 100
	}
	deviceList, total, err := deviceService.GetDeviceConfigList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	var deviceResList []businessRes.DeviceConfig
	for _, device := range deviceList.([]business.DeviceConfig) {
		deviceResList = append(deviceResList, businessRes.DeviceConfig{
			ID:               device.ID,
			Camera:           device.Camera,
			EncryptBox:       device.EncryptBox,
			EnableEncryptBox: device.EnableEncryptBox,
			DecryptBox:       device.DecryptBox,
			EnableDecryptBox: device.EnableDecryptBox,
			TcpPort:          device.TcpPort,
			UdpPort:          device.UdpPort,
			ModifyTime:       time.Unix(device.ModifyTime, 0).Format("2006-01-02 15:04:05"),
			CreateTime:       time.Unix(device.CreateTime, 0).Format("2006-01-02 15:04:05"),
		})
	}
	response.OkWithDetailed(response.PageResult{
		List:     deviceResList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// Play
// @Tags      Device
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      business.DeviceConfig
// @Success   200   {object}  response.Response{msg=string}  "播放"
// @Router    /device/play [put]
func (e *DeviceApi) Play(c *gin.Context) {
	type PlayParam struct {
		EncryptBox string `json:"encryptBox" form:"encryptBox" gorm:"comment:加密盒子ip"`
		UdpPort    string `json:"udpPort" form:"udpPort" gorm:"comment:加解密盒子处理的udp端口"`
	}
	var pp PlayParam
	err := c.ShouldBindJSON(&pp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dc := &business.DeviceConfig{
		EncryptBox: pp.EncryptBox,
		UdpPort:    pp.UdpPort,
	}
	err = deviceService.Play(dc)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
