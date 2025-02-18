package business

type DeviceConfig struct {
	ID               uint   `gorm:"primarykey" json:"id"`
	Camera           string `json:"camera" form:"camera" gorm:"comment:摄像头ip端口"`
	EncryptBox       string `json:"encryptBox" form:"encryptBox" gorm:"comment:加密盒子ip"`
	EnableEncryptBox bool   `json:"enableEncryptBox" form:"enableEncryptBox" gorm:"comment:启用加密盒子"`
	DecryptBox       string `json:"decryptBox" form:"decryptBox" gorm:"comment:解密盒子ip"`
	EnableDecryptBox bool   `json:"enableDecryptBox" form:"enableDecryptBox" gorm:"comment:启用解密盒子"`
	TcpPort          string `json:"tcpPort" form:"tcpPort" gorm:"comment:加解密盒子处理的tcp端口"`
	UdpPort          string `json:"udpPort" form:"udpPort" gorm:"comment:加解密盒子处理的udp端口"`
	ModifyTime       int64  `json:"modifyTime" form:"modifyTime" gorm:"comment:修改时间"`
	CreateTime       int64  `json:"createTime" form:"createTime" gorm:"comment:创建时间"`
}

func (d *DeviceConfig) TableName() string {
	return "config"
}
