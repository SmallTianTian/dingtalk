package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceDeviceUpdatenickRequest() *OapiSmartdeviceDeviceUpdatenickRequest {
	return &OapiSmartdeviceDeviceUpdatenickRequest{}
}

type OapiSmartdeviceDeviceUpdatenickRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiSmartdeviceDeviceUpdatenickResponse
	DeviceNickModifyVo string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiSmartdeviceDeviceUpdatenickRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) SetDeviceNickModifyVo(deviceNickModifyVo2 string) {
	this.DeviceNickModifyVo = deviceNickModifyVo2
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) GetDeviceNickModifyVo() string {
	return this.DeviceNickModifyVo
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.device.updatenick"
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_nick_modify_vo"] = this.DeviceNickModifyVo
	return txtParams
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceDeviceUpdatenickRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeviceNickModifyVo struct {
	DeviceId   string `json:"device_id,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
	Nick       string `json:"nick,omitempty"`
	Pk         string `json:"pk,omitempty"`
}
type OapiSmartdeviceDeviceUpdatenickResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
