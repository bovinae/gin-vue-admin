package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DeviceService struct{}

var DeviceServiceApp = new(DeviceService)

//@author: [bovinae](https://github.com/bovinae)
//@function: CreateDeviceConfig
//@description: 创建设备配置
//@param: e model.DeviceConfig
//@return: err error

func (s *DeviceService) CreateDeviceConfig(e business.DeviceConfig) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [bovinae](https://github.com/bovinae)
//@function: UpdateDeviceConfig
//@description: 更新设备配置
//@param: e *model.DeviceConfig
//@return: err error

func (exa *DeviceService) UpdateDeviceConfig(e *business.DeviceConfig) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [bovinae](https://github.com/bovinae)
//@function: DeleteDeviceConfig
//@description: 删除设备配置
//@param: e *model.DeviceConfig
//@return: err error

func (exa *DeviceService) DeleteDeviceConfig(e *business.DeviceConfig) (err error) {
	err = global.GVA_DB.Delete(e).Error
	return err
}

//@author: [bovinae](https://github.com/bovinae)
//@function: GetDeviceConfig
//@description: 获取设备配置
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *DeviceService) GetDeviceConfig(id request.EntityId) (dc *business.DeviceConfig, err error) {
	dc = &business.DeviceConfig{}
	err = global.GVA_DB.Where("id=?", id.ID).Find(dc).Error
	return dc, err
}

//@author: [bovinae](https://github.com/bovinae)
//@function: GetDeviceConfigList
//@description: 分页获取设备配置列表
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *DeviceService) GetDeviceConfigList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.DeviceConfig{})
	var deviceConfigList []business.DeviceConfig
	err = db.Count(&total).Error
	if err != nil {
		return deviceConfigList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&deviceConfigList).Error
	}
	return deviceConfigList, total, err
}
