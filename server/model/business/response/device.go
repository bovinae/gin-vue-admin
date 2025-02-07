package response

type DeviceConfig struct {
	ID         uint   `json:"id"`
	Camera     string `json:"camera"`
	EncryptBox string `json:"encryptBox"`
	DecryptBox string `json:"decryptBox"`
	Proto      string `json:"proto"`
	Port       int    `json:"port"`
	ModifyTime string `json:"modifyTime"`
	CreateTime string `json:"createTime"`
}
