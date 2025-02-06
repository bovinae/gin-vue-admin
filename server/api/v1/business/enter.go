package business

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DeviceApi
}

var (
	deviceService = service.ServiceGroupApp.BusinessServiceGroup.DeviceService
)
