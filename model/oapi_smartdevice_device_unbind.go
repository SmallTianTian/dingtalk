package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceDeviceUnbindRequest() *OapiSmartdeviceDeviceUnbindRequest {
	return &OapiSmartdeviceDeviceUnbindRequest{}
}

type OapiSmartdeviceDeviceUnbindRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceDeviceUnbindResponse
	DeviceUnbindVo  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceDeviceUnbindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceDeviceUnbindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceDeviceUnbindRequest) SetDeviceUnbindVo(deviceUnbindVo2 string) {
	this.DeviceUnbindVo = deviceUnbindVo2
}
func (this *OapiSmartdeviceDeviceUnbindRequest) GetDeviceUnbindVo() string {
	return this.DeviceUnbindVo
}
func (this *OapiSmartdeviceDeviceUnbindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.device.unbind"
}
func (this *OapiSmartdeviceDeviceUnbindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceDeviceUnbindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceDeviceUnbindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceDeviceUnbindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceDeviceUnbindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_unbind_vo"] = this.DeviceUnbindVo
	return txtParams
}
func (this *OapiSmartdeviceDeviceUnbindRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceDeviceUnbindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceDeviceUnbindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeviceUnbindVo struct {
	DeviceId   string `json:"device_id,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
	Pk         string `json:"pk,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
type OapiSmartdeviceDeviceUnbindResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
