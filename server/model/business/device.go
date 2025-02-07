package business

type DeviceConfig struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Camera     string `json:"camera" form:"camera" gorm:"comment:摄像头ip端口"`
	EncryptBox string `json:"encryptBox" form:"encryptBox" gorm:"comment:加密盒子ip"`
	DecryptBox string `json:"decryptBox" form:"decryptBox" gorm:"comment:解密盒子ip"`
	Proto      string `json:"proto" form:"proto" gorm:"comment:加解密盒子处理的协议"`
	Port       int    `json:"port" form:"port" gorm:"comment:加解密盒子处理的端口"`
	ModifyTime int64  `json:"modifyTime" form:"modifyTime" gorm:"comment:修改时间"`
	CreateTime int64  `json:"createTime" form:"createTime" gorm:"comment:创建时间"`
}

func (d *DeviceConfig) TableName() string {
	return "device_config"
}
