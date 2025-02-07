import request from "@/utils/request";

export function getDeviceConfigList() {
  return request({
    url: "/device/deviceConfigList",
    method: "get"
  });
}
