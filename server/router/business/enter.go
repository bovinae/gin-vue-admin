package business

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	DeviceRouter
}

var (
	deviceApi = api.ApiGroupApp.BusinessApiGroup.DeviceApi
)
