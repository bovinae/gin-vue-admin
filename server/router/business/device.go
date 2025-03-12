package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

func (e *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup) {
	deviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	deviceRouterWithoutRecord := Router.Group("device")
	{
		deviceRouter.POST("deviceConfig", deviceApi.CreateDeviceConfig)
		deviceRouter.PUT("deviceConfig", deviceApi.UpdateDeviceConfig)
		deviceRouter.DELETE("deviceConfig", deviceApi.DeleteDeviceConfig)
	}
	{
		deviceRouterWithoutRecord.GET("deviceConfig", deviceApi.GetDeviceConfig)
		deviceRouterWithoutRecord.GET("deviceConfigList", deviceApi.GetDeviceConfigList)
	}
}
