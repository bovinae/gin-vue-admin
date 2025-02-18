package response

type DeviceConfig struct {
	ID               uint   `json:"id"`
	Camera           string `json:"camera"`
	EncryptBox       string `json:"encryptBox"`
	EnableEncryptBox bool   `json:"enableEncryptBox"`
	DecryptBox       string `json:"decryptBox"`
	EnableDecryptBox bool   `json:"enableDecryptBox"`
	TcpPort          string `json:"tcpPort"`
	UdpPort          string `json:"udpPort"`
	ModifyTime       string `json:"modifyTime"`
	CreateTime       string `json:"createTime"`
}
