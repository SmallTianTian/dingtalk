package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceBindCreateRequest() *OapiSmartdeviceBindCreateRequest {
	return &OapiSmartdeviceBindCreateRequest{}
}

type OapiSmartdeviceBindCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceBindCreateResponse
	DeviceBindReqVo string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceBindCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceBindCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceBindCreateRequest) SetDeviceBindReqVo(deviceBindReqVo2 string) {
	this.DeviceBindReqVo = deviceBindReqVo2
}
func (this *OapiSmartdeviceBindCreateRequest) GetDeviceBindReqVo() string {
	return this.DeviceBindReqVo
}
func (this *OapiSmartdeviceBindCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.bind.create"
}
func (this *OapiSmartdeviceBindCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceBindCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceBindCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceBindCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceBindCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_bind_req_vo"] = this.DeviceBindReqVo
	return txtParams
}
func (this *OapiSmartdeviceBindCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceBindCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceBindCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeviceBindReqVo struct {
	BizDeviceIdentity string `json:"biz_device_identity,omitempty"`
	DeviceMac         string `json:"device_mac,omitempty"`
	DeviceServiceId   int64  `json:"device_service_id,omitempty"`
	DeviceSn          string `json:"device_sn,omitempty"`
	Nick              string `json:"nick,omitempty"`
}
type OapiSmartdeviceBindCreateResponse struct {
	taobao.TaobaoResponse
	Result  DeviceBindRespVo `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type DeviceBindRespVo struct {
	DeviceId string `json:"device_id,omitempty"`
}
